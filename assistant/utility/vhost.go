package utility

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "assistant/internal/packed"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
)

var docker = Docker{}

type Vhost struct {
	Domain       string // 主域名
	OtherDomains string // 多个域名用空格分隔
	Root         string // 项目根目录

	VhostType string // 5种类型类型，php,html,python,node,reverse

	Ssl        bool   // 是否开启https
	FoursHttps bool   // 强制302跳转到https
	FreeSsl    bool   // 是否使用免费证书
	SslCert    string // 证书文件路径
	SslKey     string // 私钥文件路径

	PathInfo bool   // 是否开启pathinfo
	Rewrite  string // 重写规则

	ErrorLog        bool   // 错误日志
	AccessLog       bool   // 访问日志
	AccessLogFormat string // 访问日志格式,默认main，还有一个选择是cloudflare

	Gzip             bool // 是否开启gzip
	Brotli           bool // 是否开启brotli
	Security         bool // 是否开启安全头
	DisableHtmlCache bool // 是否禁用html缓存

}

// 获取302跳转到https的配置
func (v *Vhost) get302Server() string {
	return `server {
    listen      80;
    listen      [::]:80;
    server_name ` + v.Domain + `;

    location / {
        return https://$host$request_uri;
    }
}`
}

func (v *Vhost) getServerBlock(port string) string {

	// 预先处理ssl证书部分
	sslConf := ""
	var e error
	if v.Ssl && port == "443" {
		if v.FreeSsl {
			sslConf, e = v.createFreeSsl(v.Domain, v.Root)
			if e != nil {
				fmt.Println("创建免费证书失败，程序退出，错误信息如下：")
				fmt.Println(e.Error())
				os.Exit(1)
			} else {
				fmt.Println("免费证书创建成功")
			}
		} else {
			sslConf = v.createCustomSsl()
		}
	}

	if port == "443" {
		port = "443 ssl http2"
	}

	// 主要配置开始
	str := `server {
    listen      ` + port + `;
    listen      [::]:` + port + `;
    server_name ` + v.Domain + `;`

	// index page
	if v.VhostType == "php" {
		str += `
	index index.php index.html index.htm;`
	} else if v.VhostType == "html" {
		str += `
	index index.html index.htm;`
	}

	// root
	str += `
	root /www/` + v.Domain + `;`

	// logs
	if v.AccessLog {
		str += `
	access_log /var/log/nginx/` + v.Domain + `/access.log ` + v.AccessLogFormat + `;`
	} else {
		str += `
	access_log off;`
	}

	// error log
	if v.ErrorLog {
		str += `
	error_log /var/log/nginx/` + v.Domain + `/error.log warn;`
	} else {
		str += `
	error_log off;`
	}
	str += `
	#error_page  404              /404.html;`

	// rewrite
	if v.Rewrite != "none" {
		str += `
	include rewrite/` + v.Rewrite + `.conf;`
	}

	// gzip
	if v.Gzip {
		str += `
	include other/gzip.conf;`
	}

	// brotil
	if v.Brotli {
		str += `
	include other/brotli.conf;`
	}

	// security header
	if v.Security {
		str += `
	include other/security.conf;`
	}

	// disable html cache
	if v.DisableHtmlCache {
		str += `
	include other/disable_html_cache.conf;`
	}

	// general path
	str += `
	include other/general.conf;`

	// vhost type
	if v.VhostType == "php" {
		if v.PathInfo {
			str += `
	location ~ [^/]\.php(/|$)
			        {
			            try_files $uri =404;
			            fastcgi_pass   php80:9000;
			            fastcgi_index index.php;
						include   fastcgi_params;
			        }`
		} else {
			str += `
	location ~ [^/]\.php(/|$) {
				        fastcgi_pass   php80:9000;
				        include        fastcgi-php.conf;
				        include        fastcgi_params;
				    }`
		}

	} else if v.VhostType == "python" {
		str += `
	location / {
			        include other/python_uwsgi.conf;
					uwsgi_pass                    127.0.0.1:8000;  #此处应该改为你的python程序ip和端口，或者容器名:端口
					uwsgi_param Host              $host;
					uwsgi_param X-Real-IP         $remote_addr;
					uwsgi_param X-Forwarded-For   $proxy_add_x_forwarded_for;
					uwsgi_param X-Forwarded-Proto $http_x_forwarded_proto;
			    }`
	} else if v.VhostType == "reverse proxy" {
		str += `
	# reverse proxy
			    location / {
			        proxy_pass            http://127.0.0.1:3000; ##此处应该改为你反向代理ip和端口，或者容器名:端口
			        proxy_set_header Host $host;
			        include               other/proxy.conf;
			    }`

	} else if v.VhostType == "node" {
		str += `
	#node
	location / {
					proxy_pass http://127.0.0.1:3000; ##此处应该改为你反向代理ip和端口，或者容器名:端口
				}`
	} else if v.VhostType == "html" {
		str += `location / {
			        root /www/` + v.Domain + `;
			    }`
	}

	// ssl
	if len(sslConf) > 0 {
		str += sslConf
	}

	str += `}`

	// 格式化返回
	p := parser.NewStringParser(str)
	c := p.Parse()
	return gonginx.DumpConfig(c, gonginx.IndentedStyle)

	// 原始返回
	// return str
}

// 添加网站
func (v *Vhost) AddVhost() {
	// @todo:告知用户操作流程

	fmt.Printf("%+v", v)

	// 1.先检测ssl和域名是否存在
	if !gfile.Exists("env/nginx/vhost/" + v.Domain + ".conf") {
		or := false
		prompt := &survey.Confirm{
			Message: "存在同名配置文件，是否要覆盖，配置文件路径:" + "env/nginx/vhost/" + v.Domain + ".conf",
			Default: true,
		}
		survey.AskOne(prompt, &or)
		if or {
			gfile.Remove("env/nginx/vhost/" + v.Domain + ".conf")
		} else {
			fmt.Println("您已取消操作，终止添加网站")
			os.Exit(0)
		}
	}
	if !gfile.Exists("env/nginx/ssl/" + v.Domain) {
		or := false
		prompt := &survey.Confirm{
			Message: "发现储存SSL配置目录，如继续将删除该存放ssl证书的目录。目录路径:" + "env/nginx/ssl/" + v.Domain,
			Default: true,
		}
		survey.AskOne(prompt, &or)
		if or {
			gfile.Remove("env/nginx/ssl/" + v.Domain)
		} else {
			fmt.Println("您已取消操作，终止添加网站")
			os.Exit(0)
		}
	}

	// 2.创建日志目录、创建网站目录,80端口的nginx配置文件
	if !gfile.Exists("logs/nginx/" + v.Domain) {
		gfile.Mkdir("logs/nginx/" + v.Domain)
	}
	if v.Root == "" {
		v.Root = "/www/" + v.Domain
	}
	if !gfile.Exists(v.Root) {
		gfile.Mkdir(v.Root)
	}
	str80 := v.getServerBlock("80")
	gfile.PutContents("env/nginx/vhost/"+v.Domain+".conf", str80)
	gfile.Chmod("env/nginx/vhost/"+v.Domain+".conf", 0755)

	// 3.重启nginx
	res, e := docker.DockerAction("restart", "nginx")
	fmt.Println(res, e)
	if IsContainerLive("nginx") {
		fmt.Println("nginx重启成功")
	} else {
		fmt.Println("nginx重启失败,请手动检查。")
		fmt.Println("推荐参考检查方法1：使用docker logs nginx查看日志")
		fmt.Println("推荐参考检查方法2：删除掉env/nginx/vhost/" + v.Domain + ".conf文件，重启再执行xii restart nginx 或进入项目目录执行docker-compose restart nginx")
		fmt.Println("最后，如果排查完成后，确定非自己引入的问题，敬请方便时反馈到项目issue中，感谢您的支持。")
		os.Exit(0)
	}

	// 4.创建ssl证书流程
	if v.Ssl {
		// 4.1. 如果需要，生成dhparam.pem
		// if !gfile.Exists("env/nginx/dhparam.pem") {
		// 	fmt.Println("发现dhparam.pem文件不存在，需要生成一次dhparam.pem文件，这个过程可能需要一段时间，请耐心等待。")
		// 	_, err := gproc.ShellExec(`docker exec -it nginx /bin/sh openssl dhparam -out /etc/nginx/dhparam.pem 2048`)
		// 	if err != nil {
		// 		fmt.Println("生成dhparam.pem文件失败，请手动执行docker exec -it nginx /bin/sh openssl dhparam -out /etc/nginx/dhparam.pem 2048")
		// 		os.Exit(0)
		// 	}
		// }

		// 4.2. 创建ssl证书目录
		fmt.Println("开始创建ssl证书目录")
		gfile.Mkdir("env/nginx/ssl/" + v.Domain)

		// 4.3. 判断是否要302跳转
		str80 := ""
		if v.FoursHttps {
			str80 = v.get302Server()
		}

		// 4.4. 重新生成nginx的80、443端口配置文件，替换原有的配置文件
		if !v.FoursHttps {
			str80 = v.getServerBlock("80")
		}
		str443 := v.getServerBlock("443")

		finalStr := str80 + "\n" + str443
		gfile.PutContents("env/nginx/vhost/"+v.Domain+".conf", finalStr)
		gfile.Chmod("env/nginx/vhost/"+v.Domain+".conf", 0755)

		// 4.5. 重启nginx
		docker.DockerAction("restart", "nginx")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("添加网站成功")
	fmt.Println("")

}

// AskForVhost cli交互获取Vhost信息
func AskForVhost() (vh *Vhost, e error) {
	vhost := Vhost{}
	var qs = []*survey.Question{
		{
			Name: "VhostType",
			Prompt: &survey.Select{
				Message: "当前网站/项目的技术类型是:",
				Options: []string{"php", "html", "python", "node", "reverse proxy"},
			},
		},

		{
			Name:     "Domain",
			Prompt:   &survey.Input{Message: "请输入主域名:"},
			Validate: survey.Required,
		},
		{
			Name:   "OtherDomains",
			Prompt: &survey.Input{Message: "请输入其他域名或者回车跳过:"},
		},
		{
			Name:   "Root",
			Prompt: &survey.Input{Message: "请输入网站根目录，不输入默认为/www/主域名:"},
		},
		{
			Name:   "PathInfo",
			Prompt: &survey.Confirm{Message: "是否开启pathinfo:"},
		}, {
			Name: "Rewrite",
			Prompt: &survey.Select{
				Message: "是否需要网址重写，无需选none:",
				Options: rewriteRules(),
			},
		}, {
			Name: "Gzip",
			Prompt: &survey.Confirm{
				Message: "是否开启Gzip压缩，提高网页访问传输速度，默认开启:",
				Default: true,
			},
		}, {
			Name: "Brotli",
			Prompt: &survey.Confirm{
				Message: "是否开启Brotli压缩(当前nginx镜像不一定支持，建议开启，程序会自动判断)，提高网页访问传输速度，默认开启:",
				Default: true,
			},
		}, {
			Name: "Security",
			Prompt: &survey.Confirm{
				Message: "是否开启通用安全信息头（Security Header），提高网站安全性，默认否:",
				Default: false,
			},
		}, {
			Name: "DisableHtmlCache",
			Prompt: &survey.Confirm{
				Message: "是否禁止浏览器缓存网页内容，开启后，每个网页都会追加一个header信息告知浏览器不要缓存。可自己在代码里控制。默认否:",
				Default: false,
			},
		}, {
			Name: "ErrorLog",
			Prompt: &survey.Confirm{
				Message: "是否开启错误日志，默认开启:",
				Default: true,
			},
		}, {
			Name: "AccessLog",
			Prompt: &survey.Confirm{
				Message: "是否开启访问日志，默认开启:",
				Default: true,
			},
		},
	}
	if err := survey.Ask(qs, &vhost); err != nil {
		return nil, err
	}

	if vhost.AccessLog {
		alf := &survey.Select{
			Message: "访问日志格式，如果有使用cloudflare的服务，建议选择cloudflare:",
			Options: []string{"main", "cloudflare"},
		}
		survey.AskOne(alf, &vhost.AccessLogFormat)
	}

	// ssl是否开启
	prompt := &survey.Confirm{
		Message: "使用启用SSL?",
		Default: true,
	}
	survey.AskOne(prompt, &vhost.Ssl)

	// 是否开启强制https跳转
	if vhost.Ssl {
		prompt := &survey.Confirm{
			Message: "强制302 http跳https",
			Default: true,
		}
		survey.AskOne(prompt, &vhost.FoursHttps)

		// 是否使用免费证书
		prompt = &survey.Confirm{
			Message: "使用免费证书:",
			Default: true,
		}
		survey.AskOne(prompt, &vhost.FreeSsl)

		if !vhost.FreeSsl {
			s := &survey.Input{
				Message: "请输入证书文件路径:",
			}
			survey.AskOne(s, &vhost.SslCert)

			s = &survey.Input{
				Message: "请输入私钥文件路径\n Enter your private key file path:",
			}
			survey.AskOne(s, &vhost.SslKey)
		}

	}

	return &vhost, nil
}

// 获取可以rewrite规则
func rewriteRules() []string {
	rewriteRules := []string{"none"}
	rset, _ := gfile.ScanDir("env/nginx/rewrite", "*.conf")
	for _, v := range rset {
		rewriteRules = append(rewriteRules, gfile.Name(v))
	}
	return rewriteRules
}

// 获取项目所在目录
func getProjectFolder() string {
	curPath, _ := os.Executable()
	return filepath.Dir(filepath.Dir(curPath))
}

// 根据vhost信息生成文件夹等
func (v *Vhost) createFolder() {
	p := getProjectFolder()
	f := []string{
		p + "/logs/nginx/" + v.Domain,
		p + "/www/" + v.Domain,
		p + "/env/nginx/ssl/" + v.Domain,
	}
	for _, v := range f {
		gfile.Mkdir(v)
	}
}

// createFreeSsl 生成证书
//  @Description: 创建证书
//  @receiver v
//  @param domain 域名
//  @param webroot 域名所在目录
//  @return string conf配置文件内容
//  @return error 错误信息
func (v *Vhost) createFreeSsl(domain string, webroot string) (string, error) {
	fmt.Println("正在开生成免费证书：" + domain)
	fmt.Println("可能会花费2-10分钟时间，请耐心等待，如失败会明确告知！")

	cmd := `docker exec -it nginx /bin/sh ~/.acme.sh/acme.sh --issue -d ` + domain + ` --webroot ` + webroot + ` --force`
	fmt.Println("执行生成命令：" + cmd)
	r, err := gproc.ShellExec(cmd)
	if strings.Contains(r, "Your cert is in") || strings.Contains(r, "Cert success.") {
		r, err = gproc.ShellExec(`docker exec -it nginx /bin/sh ~/.acme.sh/acme.sh --install-cert -d ` + domain + ` --key-file /etc/nginx/ssl/` + domain + `/key.pem --fullchain-file  /etc/nginx/ssl/` + domain + `/cert.pem`)
		if err == nil {
			// if strings.Contains(r, "Installing key to") {
			fmt.Println("证书安装成功")
			return `
						ssl_certificate /etc/nginx/ssl/` + domain + `/cert.pem;
						ssl_certificate_key /etc/nginx/ssl/` + domain + `/key.pem;
						ssl_session_timeout 5m;
						ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3;
						ssl_prefer_server_ciphers on;
						ssl_ciphers "TLS13-AES-256-GCM-SHA384:TLS13-CHACHA20-POLY1305-SHA256:TLS13-AES-128-GCM-SHA256:TLS13-AES-128-CCM-8-SHA256:TLS13-AES-128-CCM-SHA256:EECDH+CHACHA20:EECDH+CHACHA20-draft:EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5";
						ssl_session_cache builtin:1000 shared:SSL:10m;
						ssl_dhparam dhparam.pem;
				`, nil

			// }
		} else {
			return "", err
		}

	} else {
		fmt.Println("Error occured:", err)
		fmt.Println("result:", r)
		return "", errors.New("证书生成失败")
	}

	return "", nil
}

// getCustomSslConf 生成自定义证书配置
func (v *Vhost) createCustomSsl() string {
	gfile.Move(v.SslKey, "env/nginx/ssl/"+v.Domain+"/key.pem")
	gfile.Move(v.SslCert, "env/nginx/ssl/"+v.Domain+"/cert.pem")
	return `
						ssl_certificate /etc/nginx/ssl/` + v.Domain + `/cert.pem;
						ssl_certificate_key /etc/nginx/ssl/` + v.Domain + `/key.pem;
						ssl_session_timeout 5m;
						ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3;
						ssl_prefer_server_ciphers on;
						ssl_ciphers "TLS13-AES-256-GCM-SHA384:TLS13-CHACHA20-POLY1305-SHA256:TLS13-AES-128-GCM-SHA256:TLS13-AES-128-CCM-8-SHA256:TLS13-AES-128-CCM-SHA256:EECDH+CHACHA20:EECDH+CHACHA20-draft:EECDH+AES128:RSA+AES128:EECDH+AES256:RSA+AES256:EECDH+3DES:RSA+3DES:!MD5";
						ssl_session_cache builtin:1000 shared:SSL:10m;
						ssl_dhparam dhparam.pem;
				`
}

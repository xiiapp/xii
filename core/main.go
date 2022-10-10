package main

import (
	"fmt"
	"os"
	"path/filepath"

	_ "core/internal/packed"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/tufanbarisyildirim/gonginx"

	_ "github.com/tufanbarisyildirim/gonginx"
	"github.com/tufanbarisyildirim/gonginx/parser"
)

func main() {

	// go cmd.Main.Run(gctx.New())
	v, e := AskForVhost()
	if e != nil {
		fmt.Println(e.Error())
	}
	// v.AddVhost()
	s := v.getServerBlock("80")

	p := parser.NewStringParser(s)
	c := p.Parse()
	fmt.Println(gonginx.DumpConfig(c, gonginx.IndentedStyle))
}

type Vhost struct {
	Domain       string // 主域名
	OtherDomains string // 多个域名用空格分隔

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

func (v *Vhost) getFolderToCreate() []string {
	f := []string{
		"logs/nginx/" + v.Domain,
		"www/" + v.Domain,
	}
	return f
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

	str += `}`
	return str
}
func (v *Vhost) AddVhost() {

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
			Name:      "Domain",
			Prompt:    &survey.Input{Message: "请输入主域名:"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name:   "OtherDomains",
			Prompt: &survey.Input{Message: "请输入其他域名或者回车跳过:"},
		}, {
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
	rset, _ := gfile.ScanDir("../env/nginx/rewrite", "*.conf")
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

package main

import (
	"fmt"

	_ "core/internal/packed"
	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// go cmd.Main.Run(gctx.New())
	v, e := AskForVhost()
	if e != nil {
		fmt.Println(e.Error())
	}
	AddVhost(v)

}

type Vhost struct {
	Domain       string // 主域名
	OtherDomains string // 多个域名用空格分隔

	Ssl        bool   // 是否开启https
	FoursHttps bool   // 强制302跳转到https
	FreeSsl    bool   // 是否使用免费证书
	SslCert    string // 证书文件路径
	SslKey     string // 私钥文件路径

	PathInfo bool   // 是否开启pathinfo
	Rewrite  string // 重写规则
}

func (v *Vhost) Get80Str() string {
	if v.FoursHttps {
		return `server {
    listen      80;
    listen      [::]:80;
    server_name ` + v.Domain + `;

    location / {
        return 301 https://` + v.Domain + `$request_uri;
    }
}`
	}

}
func (v *Vhost) AddVhost() {

}

func AskForVhost() (Vhost, error) {

	vhost := Vhost{}
	var qs = []*survey.Question{
		{
			Name:      "Domain",
			Prompt:    &survey.Input{Message: "请输入主域名\nEnter the domain name:"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name:   "OtherDomains",
			Prompt: &survey.Input{Message: "请输入其他域名或者回车跳过\nEnter other domain name if ,else enter to skip:"},
		}, {
			Name:   "PathInfo",
			Prompt: &survey.Confirm{Message: "是否开启pathinfo\nEnable pathinfo:"},
		}, {
			Name: "Rewrite",
			Prompt: &survey.Select{
				Message: "是否需要网址重写，无需选none\n Url Rewrite Rule?:",
				Options: []string{"none", "yii2", "laravel"},
			},
		},
	}
	if err := survey.Ask(qs, &vhost); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error occurred!")
		return vhost, err
	}

	// ssl是否开启
	prompt := &survey.Confirm{
		Message: "使用启用SSL?\n Do you want to enable https ?",
		Default: true,
	}
	survey.AskOne(prompt, &vhost.Ssl)

	// 是否开启强制https跳转
	if vhost.Ssl {
		prompt := &survey.Confirm{
			Message: "强制http跳https\n Set 302 to https when visit http ?",
			Default: true,
		}
		survey.AskOne(prompt, &vhost.FoursHttps)

		// 是否使用免费证书
		prompt = &survey.Confirm{
			Message: "使用免费证书\n Use free ssl ,choose No to input your custom cert file?",
			Default: true,
		}
		survey.AskOne(prompt, &vhost.FreeSsl)

		if !vhost.FreeSsl {
			s := &survey.Input{
				Message: "请输入证书文件路径\n Enter your public cert file path:",
			}
			survey.AskOne(s, &vhost.SslCert)

			s = &survey.Input{
				Message: "请输入私钥文件路径\n Enter your private key file path:",
			}
			survey.AskOne(s, &vhost.SslKey)
		}

	}

	return vhost, nil
}

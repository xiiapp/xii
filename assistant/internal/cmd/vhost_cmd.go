package cmd

import (
	"context"
	"fmt"

	"assistant/utility"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	VhostCmd = &gcmd.Command{
		Name:  "vhost",
		Usage: "新增网站：xii vhost -add 删除网站：xii vhost -del  查看所有网站：xii vhost -list",
		Brief: "新增、删除、罗列所有网站",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("")
			fmt.Println("为名为nginx的容器增删查网站。\n注意，如需修改具体网站nginx配置，请自己编辑nginx配置文件。一般配置文件地址为：env/nginx/vhost/域名.conf")
			fmt.Println("用法：")
			fmt.Println("   新增网站：xii vhost add")
			fmt.Println("   删除网站：xii vhost del")
			fmt.Println("   查看所有网站：xii vhost list")
			fmt.Println("")
			return nil
		},
	}
	VhostAddCmd = &gcmd.Command{
		Name:  "add",
		Usage: "新增网站: xii vhost add",
		Brief: "新增网站",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			v, e := utility.AskForVhost()
			if e != nil {
				fmt.Println(e.Error())
			}
			v.AddVhost()
			return nil
		},
	}
	VhostDelCmd = &gcmd.Command{
		Name:  "del",
		Usage: "删除网站: xii vhost del",
		Brief: "删除网站",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("todo")
			return nil
		},
	}
	VhostListCmd = &gcmd.Command{
		Name:  "list",
		Usage: "查看所有网站: xii vhost list",
		Brief: "查看所有网站",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("todo")
			return nil
		},
	}
)

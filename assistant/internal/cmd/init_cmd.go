package cmd

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	InitCmd = &gcmd.Command{
		Name:  "init",
		Usage: "xii init",
		Brief: "可多次选择所需的镜像组件。该命令不删任何生产数据。",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("")
			fmt.Println("")
			return nil
		},
	}

	InitClearCmd = &gcmd.Command{
		Name:  "clear",
		Usage: "xii clear",
		Brief: "清理镜像除配置外的所有数据",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			// 一期限定
			basePath := "/home/xii"

			op := false
			alf := &survey.Confirm{
				Message: "确认清理镜像除配置外的所有数据？\n  他将清理掉data/目录、www/目录、logs/目录下的所有数据。\n  原理是删除后重建，操作不可逆，请谨慎选择！",
				Default: false,
			}
			survey.AskOne(alf, &op)
			if op {
				gfile.Remove(gfile.Join(basePath, "data/"))
				gfile.Remove(gfile.Join(basePath, "logs/"))
				gfile.Remove(gfile.Join(basePath, "www/"))
				gfile.Mkdir(gfile.Join(basePath, "data/"))
				gfile.Mkdir(gfile.Join(basePath, "logs/"))
				gfile.Mkdir(gfile.Join(basePath, "www/"))

				fmt.Println("清理完成! 请重新启动镜像")
			}

			return nil
		},
	}
)

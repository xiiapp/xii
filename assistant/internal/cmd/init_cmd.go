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
			op := false
			alf := &survey.Confirm{
				Message: "确认清理镜像除配置外的所有数据？他将清理掉data/目录、www/目录、logs/目录下的所有数据。操作不可逆！",
				Default: true,
			}
			survey.AskOne(alf, &op)
			if op {
				gfile.Remove(gfile.Join("/Users/mou/goProjects/xii", "data/"))
				gfile.Remove(gfile.Join("/Users/mou/goProjects/xii", "logs/"))
				gfile.Remove(gfile.Join("/Users/mou/goProjects/xii", "www/"))

				fmt.Println("清理完成! 请重新启动镜像")
			}

			return nil
		},
	}
)

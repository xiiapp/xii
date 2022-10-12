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
		Brief: "[TODO] 可多次选择所需的镜像组件。该命令不删任何生产数据。",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println(`
当前一期版本暂时不支持通过该命令来选择组件，该功能为二期内容。
请手动修改目录下的docker-compose.yml 文件(取消或增加注释)来修改组件。
设置账号密码数据则手动修改目录下的.env 文件。
`)
			fmt.Println("")
			return nil
		},
	}

	InitClearCmd = &gcmd.Command{
		Name:  "clear",
		Usage: "xii clear",
		Brief: "清理镜像除配置外的所有数据\n-\n-\n---------以下为封装docker-compose相关命令---------\n封装原因是：docker-compose通常需要进入目录或带-f参数，操作过于麻烦。\n使用上注意所有容器名字需要带-，比如php容器就写做-php。\n-",
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

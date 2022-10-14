package cmd

import (
	"context"
	"fmt"
	"os"

	"assistant/utility"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	InitCmd = &gcmd.Command{
		Name:  "init",
		Usage: "xii init",
		Brief: "可多次选择所需的镜像容器组件。该命令不删任何生产数据。",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			// 检索仓库模板
			p := "/Users/mou/goProjects/xii/repo"
			c := []string{}
			rset, _ := gfile.ScanDir(p, "*")
			for _, v := range rset {
				v2 := gfile.Basename(v)
				if v2 != "base" && gfile.IsDir(v) {
					c = append(c, v2)
				}
			}

			// 选择区块
			defaultSelect := utility.GetExistDockerFile()
			cSlice := []string{}
			alf := &survey.MultiSelect{
				Message:  " 请选择你所需的镜像容器组件,他将重新生成配置文件。\n   如果已经在用的镜像，将保存现有配置。！\n",
				PageSize: 20,
				Options:  c,
				Default:  defaultSelect,
			}
			survey.AskOne(alf, &cSlice)

			fmt.Println("你选择的容器组件是：" + gstr.Join(cSlice, ","))

			utility.GenerateDockerComposeFile(cSlice)
			utility.GenerateEnvFile(cSlice)
			utility.CopyDirIfExist(cSlice)

			fmt.Println("\n\n")
			fmt.Println("------------------")
			fmt.Println("初始化/修改完成")
			fmt.Println("注 意：谨慎起见，移除容器组件不会删除已经存在的数据，你需要手动删除数据和容器/镜像")
			fmt.Println("建议1：如果是第一次初始化，请执行：xii up 或 xii up -d 来启动容器实例")
			fmt.Println("建议2：如果是新增容器组件的，可以执行 xii up -容器组件 来单独启动容器组件，或者 xii up 直接启动所有容器组件,已存活的容器组件将被忽略")
			fmt.Println("建议3：如果是移除容器组件的，可以执行 xii stop -容器组件 来停止容器组件，然后执行 xii rm -容器组件 来移除容器/镜像，然后手动删除配置文件，数据文件等（/data,/env/容器组件名）")
			fmt.Println("------------------")
			fmt.Println("\n\n")

			os.Exit(0)
			return nil
		},
	}

	InitClearCmd = &gcmd.Command{
		Name:  "clear",
		Usage: "xii clear",
		Brief: "清理镜像除配置外的所有数据\n-\n-\n---------以下为封装docker-compose/docker相关命令---------\n封装原因是：docker-compose通常需要进入目录或带-f参数，操作过于麻烦。\n使用上注意所有容器名字需要带-，比如php容器就写做-php。\n-",
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

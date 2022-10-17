package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"assistant/utility"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gproc"
)

var (
	UpCmd = &gcmd.Command{
		Name:  "up",
		Usage: "xii up -参数",
		Brief: "封装docker-compose up 命令,\n\n首次使用请执行： xii up -h 查看下使用说明。",
		Examples: `# 创建并且启动所有容器:  xii up
# 创建并且(后台运行方式)启动所有容器:  xii up -d
# 创建并且启动nginx、php、mysql的多个容器:  xii up  -php -nginx -mysql
# 创建并且(后台运行方式)启动nginx、php、mysql的多个容器:  xii up -d -php -nginx -mysql
`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("up", parser)
			return nil
		},
	}

	StartCmd = &gcmd.Command{
		Name:  "start",
		Usage: "xii start -容器名",
		Brief: "封装docker-compose start 命令,启动服务。范例：xii start -php \n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("start", parser)
			return nil
		},
	}

	StopCmd = &gcmd.Command{
		Name:  "stop",
		Usage: "xii stop -容器名",
		Brief: "封装docker-compose stop 命令,停止指定容器服务。范例：xii stop -php\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("stop", parser)
			return nil
		},
	}
	RestartCmd = &gcmd.Command{
		Name:  "restart",
		Usage: "xii restart -容器名",
		Brief: "封装docker-compose stop 命令,重启指定容器服务。范例：xii restart -php\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("restart", parser)
			return nil
		},
	}

	BuildCmd = &gcmd.Command{
		Name:  "build",
		Usage: "xii build -容器名",
		Brief: "封装docker-compose build 命令,构建或者重新构建服务。范例：xii build -php\n-\n-\n---------以下为快速进入某个容器内部的快捷方法-------\n-",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("build", parser)
			return nil
		},
	}

	RmCmd = &gcmd.Command{
		Name:  "rm",
		Usage: "xii rm -容器名",
		Brief: "封装docker-compose rm 命令,删除并且停止php容器。但不会删除挂载的数据卷\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("rm", parser)
			return nil
		},
	}
	RmAllCmd = &gcmd.Command{
		Name:  "rmall",
		Usage: "xii rmall",
		Brief: "停止并删除所有容器、所有镜像,但不会删除挂载的数据卷.\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			// 二次确认
			confirm := false
			survey.AskOne(&survey.Confirm{
				Message: "确定要删除所有容器、所有镜像吗？",
				Default: false,
			}, &confirm)
			if !confirm {
				fmt.Println("取消删除")
				os.Exit(0)
			}

			// 停止所有的容器
			r, e := gproc.ShellExec(`docker stop $(docker ps -aq)`)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(1)
			}
			fmt.Println(r)

			// 删除所有的容器
			r, e = gproc.ShellExec(`docker rm $(docker ps -aq)`)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(1)
			}
			fmt.Println(r)

			// 删除所有的镜像
			r, e = gproc.ShellExec(`docker rmi $(docker images -q)`)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(1)
			}
			fmt.Println(r)
			os.Exit(0)
			return nil
		},
	}

	DownCmd = &gcmd.Command{
		Name:  "down",
		Usage: "xii down",
		Brief: "封装docker-compose down 命令,停止并删除容器，网络，图像。注意：不会删除挂载的数据卷。\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("down", parser)
			return nil
		},
	}

	RebuildCmd = &gcmd.Command{
		Name:  "rebuild",
		Usage: "xii rebuild -容器名",
		Brief: "重现编译某指定容器，用于在修改docker-compose.yml后。\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			// 获取参数
			om := parser.GetOptAll()
			containers := []string{}
			imagesIds := []string{}
			if len(om) >= 0 {
				for k, _ := range om {
					containers = append(containers, k)
					// 获取镜像id
					iid, _ := gproc.ShellExec(`docker ps -a| grep ` + k + ` | awk '{print $1}'`)
					if iid != "" {
						imagesIds = append(imagesIds, iid)
					}
				}
			}
			if len(containers) == 0 {
				fmt.Println("请输入一个或以上的容器名")
				os.Exit(1)
			}
			containersStr := strings.Join(containers, " ")
			imagesIdsStr := strings.Join(imagesIds, " ")

			// 停止所有的容器
			fmt.Println("开始停止容器:" + containersStr)
			r, e := gproc.ShellExec(`docker stop ` + containersStr)
			if e != nil {
				fmt.Println("停止容器(container)失败，失败原因：" + e.Error())
				os.Exit(1)
			}
			fmt.Println(r)

			// 删除所有的容器
			fmt.Println("开始删除容器:" + containersStr)
			r, e = gproc.ShellExec(`docker rm ` + containersStr)
			if e != nil {
				fmt.Println("删除容器(container)失败，失败原因：" + e.Error())
				os.Exit(1)
			}
			fmt.Println(r)

			// 删除所有的镜像
			fmt.Println("开始停止容器对应镜像:" + containersStr)
			r, e = gproc.ShellExec(`docker rmi ` + imagesIdsStr)
			if e != nil {
				fmt.Println("删除镜像(images)失败，执行命令：" + `docker rmi ` + imagesIdsStr)
				os.Exit(1)
			}

			// 重新编译
			fmt.Println("开始")
			doker := utility.Docker{}
			argSlice := []string{"-d", "--no-deps", "--build"}
			argSlice = append(argSlice, containers...)
			doker.DockerActionShell("up", argSlice)
			fmt.Println("重新编译完成")
			os.Exit(0)
			return nil
		},
	}

	PsCmd = &gcmd.Command{
		Name:  "ps",
		Usage: "xii ps",
		Brief: "封装docker ps 命令,经常顺手操作，所以封装一下。\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerRun("ps", parser)
			return nil
		},
	}
)

// 执行docker-compose命令参数
func dockerComposeRun(typ string, parser *gcmd.Parser) {
	// 获取参数
	om := parser.GetOptAll()
	d := ""
	containers := []string{}
	if len(om) >= 0 {
		for k, _ := range om {
			if "d" == strings.ToLower(k) {
				d = "-d"
			} else {
				containers = append(containers, k)
			}

		}
	}

	dstr := []string{}
	if d != "" {
		dstr = append(dstr, "-d")
	}

	if len(containers) > 0 {
		dstr = append(dstr, containers...)
	}

	docker := utility.Docker{}
	docker.Auto()
	docker.DockerActionShell(typ, dstr)
}

// 执行docker命令参数
func dockerRun(typ string, parser *gcmd.Parser) {
	// 获取参数
	om := parser.GetOptAll()
	containers := []string{}
	if len(om) >= 0 {
		for k, _ := range om {
			containers = append(containers, "-"+k)
		}
	}
	docker := utility.Docker{}
	docker.Auto()
	docker.DockerActionShell(typ, containers)
}

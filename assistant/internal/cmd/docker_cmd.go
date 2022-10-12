package cmd

import (
	"context"
	"strings"

	"assistant/utility"
	"github.com/gogf/gf/v2/os/gcmd"
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
		Brief: "封装docker-compose rm 命令,删除并且停止php容器\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("rm", parser)
			return nil
		},
	}

	DownCmd = &gcmd.Command{
		Name:  "down",
		Usage: "xii down",
		Brief: "封装docker-compose down 命令,停止并删除容器，网络，图像和挂载卷\n",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			dockerComposeRun("down", parser)
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
	docker.DockerActionShell("up", dstr)
}

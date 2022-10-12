package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"assistant/utility"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	NginxCmd = &gcmd.Command{
		Name:  "nginx",
		Usage: "xii nginx",
		Brief: "快速进入nginx容器",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			utility.Enter("nginx")
			return nil
		},
	}

	ContainerCmd = &gcmd.Command{
		Name:  "go",
		Usage: "xii go",
		Brief: "快速进入容器, 直接xii go 将罗列所有存活容器选择，可以xii go -关键字 加以过滤，缩小选择范围.\n部分容器的入口不是/bin/sh，则会进入失败",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			om := parser.GetOptAll()
			keyword := ""
			if len(om) >= 0 {
				for k, _ := range om {
					keyword = k
				}
			}

			d := utility.Docker{}
			d.Auto()
			s, e := d.ContainerName()
			if e != nil {
				fmt.Println(e.Error())
			}
			s2 := []string{}
			if len(s) > 0 {
				for _, v := range s {
					if utility.IsContainerLive(v) {
						if keyword != "" {
							if strings.Contains(v, keyword) {
								s2 = append(s2, v)
							}
						} else {
							s2 = append(s2, v)
						}
					}
				}
			}

			if len(s2) == 0 {
				fmt.Println("没找到存活的容器，无法进入")
				os.Exit(1)
				return nil
			}

			if len(s2) == 1 {
				utility.Enter(s2[0])
			} else {
				c := ""
				alf := &survey.Select{
					Message: "请选择容器进入:",
					Options: s2,
				}
				survey.AskOne(alf, &c)
				utility.Enter(c)
			}

			return nil
		},
	}

	PhpCmd = &gcmd.Command{
		Name:  "php",
		Usage: "xii php",
		Brief: "快速进入*php*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("php")
			return nil
		},
	}

	MysqlCmd = &gcmd.Command{
		Name:  "mysql",
		Usage: "xii mysql",
		Brief: "快速进入mysql*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("mysql")
			return nil
		},
	}

	RedisCmd = &gcmd.Command{
		Name:  "redis",
		Usage: "xii redis",
		Brief: "快速进入*redis*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("redis")
			return nil
		},
	}

	SupervisorCmd = &gcmd.Command{
		Name:  "supervisor",
		Usage: "xii supervisor",
		Brief: "快速进入*supervisor*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("supervisor")
			return nil
		},
	}

	MongodbCmd = &gcmd.Command{
		Name:  "mongodb",
		Usage: "xii mongodb",
		Brief: "快速进入*mongodb*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("mongodb")
			return nil
		},
	}

	NodeCmd = &gcmd.Command{
		Name:  "node",
		Usage: "xii node",
		Brief: "快速进入*node*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("node")
			return nil
		},
	}

	GoCmd = &gcmd.Command{
		Name:  "golang",
		Usage: "xii golang",
		Brief: "快速进入*golang*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("go")
			return nil
		},
	}

	EsCmd = &gcmd.Command{
		Name:  "elasticsearch",
		Usage: "xii es",
		Brief: "快速进入*elasticsearch*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("elasticsearch")
			return nil
		},
	}

	KibanaCmd = &gcmd.Command{
		Name:  "kibana",
		Usage: "xii kibana",
		Brief: "快速进入*elasticsearch*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("kibana")
			return nil
		},
	}

	LogstashCmd = &gcmd.Command{
		Name:  "logstash",
		Usage: "xii logstash",
		Brief: "快速进入*logstash*容器，多项目时需选择",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			QuickEnter("logstash")
			return nil
		},
	}
)

// QuickEnter 快入进入一个指定的容器
func QuickEnter(keyword string) {
	d := utility.Docker{}
	s, e := d.ContainerName()
	if e != nil {
		fmt.Println(e.Error())
	}
	s2 := []string{}
	if len(s) > 0 {
		for _, v := range s {
			if utility.IsContainerLive(v) {
				if strings.Contains(v, keyword) {
					s2 = append(s2, v)
				}
			}
		}
	}

	if len(s2) == 0 {
		fmt.Println("没找到存活的" + keyword + "容器，无法进入")
		os.Exit(1)
	}

	if len(s2) == 1 {
		utility.Enter(s2[0])
	} else {
		c := ""
		alf := &survey.Select{
			Message: "请选择容器进入:",
			Options: s2,
		}
		survey.AskOne(alf, &c)
		utility.Enter(c)
	}
}

package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/os/gcmd"
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
			fmt.Println("")
			fmt.Println("")
			return nil
		},
	}
)

package main

import (
	"assistant/internal/cmd"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = &gcmd.Command{
		Name:        "main",
		Brief:       "XII助手",
		Description: "XII助手:管理网站新增/删除/罗列，增删改Xii自带的Docker镜像组件",
	}
)

func main() {

	// vhost 管理相关
	vhostCmd := cmd.VhostCmd
	vhostCmd.AddCommand(cmd.VhostAddCmd, cmd.VhostDelCmd, cmd.VhostListCmd)

	err := Main.AddCommand(vhostCmd, cmd.InitCmd, cmd.InitClearCmd)
	if err != nil {
		panic(err)
	}
	Main.Run(gctx.New())
}

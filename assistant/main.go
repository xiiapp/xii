package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"assistant/internal/cmd"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = &gcmd.Command{
		Name:        "main",
		Brief:       "XII助手",
		Description: "XII助手:管理网站新增/删除/罗列，以及封装docker的快捷操作方式",
	}
)

func main() {
	SetupCloseHandler()

	// vhost 管理相关
	vhostCmd := cmd.VhostCmd
	vhostCmd.AddCommand(cmd.VhostAddCmd, cmd.VhostDelCmd, cmd.VhostListCmd)
	Main.AddCommand(vhostCmd, cmd.InitClearCmd)

	Main.AddCommand(cmd.UpCmd, cmd.DownCmd, cmd.StartCmd, cmd.StopCmd, cmd.RestartCmd, cmd.RmCmd, cmd.BuildCmd)

	// docker 便捷操作
	Main.AddCommand(cmd.ContainerCmd)
	Main.AddCommand(cmd.NginxCmd, cmd.PhpCmd, cmd.MysqlCmd, cmd.MongodbCmd, cmd.RedisCmd, cmd.SupervisorCmd, cmd.EsCmd, cmd.KibanaCmd, cmd.LogstashCmd, cmd.NodeCmd, cmd.GoCmd)

	Main.Run(gctx.New())
}

// SetupCloseHandler
// ref:twle.cn/t/381
func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

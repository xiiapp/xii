// Package main
// @Description: docker 容器管理相关
//

package utility

import (
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
)

type Docker struct {
	ComposeFile string // docker-compose.yml文件路径，没有就默认是当前
}

// DockerAction 包装docker-compose命令
//  @Description: 容器实例操作
//  @example:
//          docker-compose up                         # 创建并且启动所有容器
//          docker-compose up -d                      # 创建并且后台运行方式启动所有容器
//          docker-compose up nginx php mysql         # 创建并且启动nginx、php、mysql的多个容器
//          docker-compose up -d nginx php  mysql     # 创建并且已后台运行的方式启动nginx、php、mysql容器
//          docker-compose start php                  # 启动服务
//          docker-compose stop php                   # 停止服务
//          docker-compose restart php                # 重启服务
//          docker-compose build php                  # 构建或者重新构建服务
//          docker-compose rm php                     # 删除并且停止php容器
//          docker-compose down                       # 停止并删除容器，网络，图像和挂载卷
//  @param action  restart|stop|start|kill
//  @param container
//  @return string
//  @return error
func (d *Docker) DockerAction(action string, arg string) (string, error) {
	if d.ComposeFile == "" {
		return gproc.ShellExec(`docker-compose ` + action + ` ` + arg)
	} else {
		return gproc.ShellExec(`docker-compose -f ` + d.ComposeFile + ` ` + action + ` ` + arg)
	}
}

// ContainerName
//  @Description: 获取所有容器名称,不管容器是否运行
//  @return []string
//  @return error
func (d *Docker) ContainerName() ([]string, error) {
	r, e := gproc.ShellExec(`docker inspect --format='{{.Name}}' $( docker ps -aq --no-trunc) | cut -c2-`)
	if e != nil {
		return nil, e
	}
	result := gstr.SplitAndTrim(r, "\n")
	return result, nil
}

// IsContainerLive 检查容器是否正在运行
func IsContainerLive(container string) bool {
	r, e := gproc.ShellExec(`docker ps -q -f "name=^` + container + `$"`)
	if e != nil {
		return false
	}
	if r == "" {
		return false
	}
	return true
}
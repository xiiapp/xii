package utility

import (
	"log"
	"os/user"
	"runtime"
)

// GetInstallDir 获取系统安装目录
func GetInstallDir() string {

	// for dev
	// return "/Users/mou/goProjects/xii"

	switch runtime.GOOS {
	case "darwin":
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		return u.HomeDir + "/xii"
	case "linux":
		return "/home/xii"
	}
	return "暂未支持"
}

func GetDockerComposeFile() string {
	return GetInstallDir() + "/docker-compose.yml"
}

func GetEnvFile() string {
	return GetInstallDir() + "/.env"
}

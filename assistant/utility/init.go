// Package utility
// @Description: 生成配置文件/目录
//
package utility

import (
	"fmt"
	"os"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
)

// GenerateDockerComposeFile 生成docker-compose.yml
func GenerateDockerComposeFile(components []string) {
	p := GetInstallDir()
	// for dev
	existDcMap := parseDockerComposeFileOrEnvFile(true)

	// pre-check
	if len(components) == 0 {
		fmt.Println("请指定需要生成的组件")
		os.Exit(0)
	}

	fmt.Println("正在生成docker-compose.yml...")
	cStr := []string{}
	gfile.ReadLines(p+"/repo/base/compose.yml", func(line string) error {
		cStr = append(cStr, line)
		return nil
	})
	for _, v := range components {
		// 已有的不修改
		if len(existDcMap) > 0 {
			if value, ok := existDcMap[v]; ok {
				cStr = append(cStr, "\n\n\n#clearfix")
				cStr = append(cStr, value)
				continue
			}
		}

		// 新增的
		cStr = append(cStr, "\n\n#clearfix\n#app:"+v+"\n")
		gfile.ReadLines(p+"/repo/"+v+"/compose.yml", func(line string) error {
			cStr = append(cStr, line)
			return nil
		})
	}

	gfile.PutContents(p+"/docker-compose.yml", strings.ReplaceAll(strings.Join(cStr, "\n"), "\n\n", "\n"))

}

// GenerateEnvFile 生成.env文件
func GenerateEnvFile(components []string) {
	p := GetInstallDir()
	// for dev
	existDcMap := parseDockerComposeFileOrEnvFile(false)

	// pre-check
	if len(components) == 0 {
		fmt.Println("请指定需要生成的组件")
		os.Exit(0)
	}

	fmt.Println("正在生成.env文件(.env存储docker启动时所需变量，可自行修改)...")
	cStr := []string{}
	gfile.ReadLines(p+"/repo/base/env.sample", func(line string) error {
		cStr = append(cStr, line)
		return nil
	})
	for _, v := range components {
		// 已有的不修改
		if len(existDcMap) > 0 {
			if value, ok := existDcMap[v]; ok {
				cStr = append(cStr, "\n\n\n#clearfix")
				cStr = append(cStr, value)
				continue
			}
		}

		// 新增的
		if gfile.Exists(p + "/repo/" + v + "/env.sample") {
			cStr = append(cStr, "\n\n#clearfix\n#app:"+v+"\n")
			gfile.ReadLines(p+"/repo/"+v+"/env.sample", func(line string) error {
				cStr = append(cStr, line)
				return nil
			})
		}
	}

	gfile.PutContents(p+"/.env", strings.ReplaceAll(strings.Join(cStr, "\n"), "\n\n", "\n"))

}

// CopyDirIfExist 复制目录，如果目录存在
func CopyDirIfExist(components []string) {
	fmt.Println("正在生成拷贝配置文件/目录...")
	p := GetInstallDir()
	for _, v := range components {
		src := p + "/repo/" + v + "/build"
		dst := p + "/env/" + v
		if gfile.Exists(src) && gfile.IsDir(src) && !gfile.Exists(dst) {
			gfile.CopyDir(src, dst)
		}
	}
}

// 解析docker-compose.yml或者env文件。
// 因为为了保存注释，检测断点标志，所以不能用包解析。只能字符串切割
func parseDockerComposeFileOrEnvFile(isCompose bool) map[string]string {
	p := GetEnvFile()
	if isCompose {
		p = GetDockerComposeFile()
	}

	rest := make(map[string]string)
	if !gfile.IsFile(p) {
		return rest
	}

	str := gfile.GetContents(p)
	strSlice := gstr.Split(str, "#clearfix")
	for _, v := range strSlice {
		t, _ := gregex.Match(`#app:(\w+)`, []byte(v))
		if len(t) >= 2 {
			// fmt.Println(string(t[1]))
			rest[string(t[1])] = gstr.TrimLeft(v, "\n")
		}
	}
	return rest
}

// GetExistDockerFile 获取已经部署的组件
func GetExistDockerFile() []string {
	p := GetDockerComposeFile()
	rest := []string{}
	if !gfile.IsFile(p) {
		return rest
	}

	str := gfile.GetContents(p)
	strSlice := gstr.Split(str, "#clearfix")
	for _, v := range strSlice {
		t, _ := gregex.Match(`#app:(\w+)`, []byte(v))
		if len(t) >= 2 {
			rest = append(rest, string(t[1]))
		}
	}
	return rest
}

package main

import (
	"github.com/gogf/gf/v2/os/gproc"
)

func main() {
	// r, _ := ContainerName()
	// fmt.Printf("%+v", r)
	// r, _ := DockerAction("stop", "nginx")
	// fmt.Printf("%+v", r)
}

//
// func ContainerName() ([]string, error) {
// 	r, e := gproc.ShellExec(`docker inspect --format='{{.Name}}' $(sudo docker ps -aq --no-trunc) | cut -c2-`)
// 	if e != nil {
// 		return nil, e
// 	}
// 	result := gstr.SplitAndTrim(r, "\n")
// 	return result, nil
// }

func DockerAction(action string, arg string) (string, error) {
	return gproc.ShellExec(`docker-compose ` + action + ` ` + arg)
}

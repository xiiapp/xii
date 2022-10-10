package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("docker", "exec -it nginx /bin/sh ~/.acme.sh/acme.sh --issue -d docker.mallka.com --webroot /www/localhost").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

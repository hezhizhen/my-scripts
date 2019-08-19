package main

import (
	"os"
	"os/exec"
)

func main() {
	repo := os.Args[1]
	url := "https://github.com/caicloud/" + repo
	cmd := exec.Command("open", url)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

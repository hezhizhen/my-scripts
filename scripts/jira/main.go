package main

import (
	"os"
	"os/exec"
)

func main() {
	issue := os.Args[1]
	url := "http://jira.caicloud.xyz/browse/" + issue
	cmd := exec.Command("open", url)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

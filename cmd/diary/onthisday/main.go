package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hezhizhen/my-scripts/pkg/util"
)

func main() {
	specifiedDate := flag.String("date", "", "specified date. e.g.: 2016.01.02")
	editor := flag.String("editor", "mvim", "specified editor. e.g.: macvim, vimr, atom, code")
	flag.Parse()
	now := time.Now()
	var err error
	if *specifiedDate != "" {
		now, err = time.Parse("2006.01.02", *specifiedDate)
		util.Check(err)
	}
	if *editor == "macvim" {
		*editor = "mvim"
	}
	if *editor == "code" {
		*editor = "code-insiders"
	}
	dir := fmt.Sprintf("%s/Dropbox/Diary", util.GetHome())
	files, err := filepath.Glob(dir + fmt.Sprintf("/*-%02d-%02d.md", now.Month(), now.Day()))
	util.Check(err)
	cmd := exec.Command(*editor, files...)
	util.Check(cmd.Run())
}

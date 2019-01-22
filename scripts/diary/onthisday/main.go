package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hezhizhen/tiny-tools/utilz"
)

func main() {
	specifiedDate := flag.String("date", "", "specified date. e.g.: 2016.01.02")
	editor := flag.String("editor", "mvim", "specified editor. e.g.: macvim, vimr, atom, code")
	flag.Parse()
	now := time.Now()
	var err error
	if *specifiedDate != "" {
		now, err = time.Parse("2006.01.02", *specifiedDate)
		utilz.Check(err)
	}
	if *editor == "macvim" {
		*editor = "mvim"
	}
	if *editor == "code" {
		*editor = "code-insiders"
	}
	dir := "/Users/hezhizhen/Dropbox/Diary"
	files, err := filepath.Glob(dir + fmt.Sprintf("/*-%02d-%02d.md", now.Month(), now.Day()))
	utilz.Check(err)
	cmd := exec.Command(*editor, files...)
	utilz.Check(cmd.Run())
}

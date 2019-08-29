package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hezhizhen/my-scripts/pkg/util"
)

func main() {
	const editor = "mvim"
	now := time.Now()
	var err error
	dir := fmt.Sprintf("%s/Dropbox/Diary", util.GetHome())
	files, err := filepath.Glob(dir + fmt.Sprintf("/%04d-%02d-*.md", now.Year(), now.Month()))
	util.Check(err)
	cmd := exec.Command(editor, files...)
	util.Check(cmd.Run())
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/hezhizhen/my-scripts/pkg/util"
)

func main() {
	flag.Parse()
	// check the first arugment only
	operation := flag.Arg(0)
	switch operation {
	case "save":
		saveAllFiles()
	case "recover":
		recoverAllFiles()
	default:
		fmt.Println("Invalid operation:", operation)
		fmt.Println("Available operations:\n\t- save: Save tags to .matadata\n\t- recover: Recover tags from .matadata")
	}
}

func recoverAllFiles() {
	// every directory must have a file named .matadata
	bs, err := ioutil.ReadFile(".matadata")
	util.Check(err)
	raw := strings.TrimSpace(string(bs))
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		// key: filename
		// value: tags separated by ,
		kv := strings.Split(line, ";")
		recoverTags(kv[0], kv[1])
	}
}

func recoverTags(filename, tags string) {
	// use tag to add tags to the file
	cmd := exec.Command("tag", "--add", tags, filename)
	util.Check(cmd.Run())
}

func saveAllFiles() {
	// list all files in the current directory
	// TODO: exclude sub dirs
	cmd := exec.Command("ls", "-a", ".")
	bs, err := cmd.Output()
	util.Check(err)
	raw := strings.TrimSpace(string(bs))
	fs := strings.Split(raw, "\n")
	for i := range fs {
		fs[i] = strings.TrimSpace(fs[i])
	}
	// create .matadata every time
	f, err := os.Create(".matadata")
	util.Check(err)
	f.Close()
	for _, f := range fs {
		tags := retrieveTags(f)
		fmt.Println(f, tags)
		if tags != "null" {
			saveTags(f, tags)
		}
	}
}

func retrieveTags(filename string) string {
	cmd := exec.Command("mdls", "-raw", "-name", "kMDItemUserTags", filename)
	bs, err := cmd.Output()
	util.Check(err)
	raw := string(bs)
	// remove ( and ) and spaces
	raw = strings.TrimPrefix(raw, "(")
	raw = strings.TrimSuffix(raw, ")")
	raw = strings.TrimSpace(raw)
	lines := strings.Split(raw, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	tags := strings.Join(lines, "")
	return tags
}

func saveTags(filename, tags string) {
	// file opened by os.Open() is read-only
	f, err := os.OpenFile(".matadata", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	util.Check(err)
	defer f.Close()
	// syntax for .matadata:
	// one line for each file
	// filename;tag1,tag2,tag3...\n
	_, err = f.WriteString(fmt.Sprintf("%s;%s\n", filename, tags))
	util.Check(err)
}

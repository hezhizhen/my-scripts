package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"

	. "github.com/hezhizhen/tiny-tools/utilz"
)

type library string

const (
	// MWeb3 is the current library
	MWeb3 library = "/Users/hezhizhen/Dropbox/MWeb3/docs/"
	// MWeb2 is the previous library
	MWeb2 library = "/Users/hezhizhen/Dropbox/MWeb/docs/"
	// nvALT is where I store drafts
	nvALT library = "/Users/hezhizhen/Dropbox/Notational Data/"
)

type fileInfo struct {
	FileName string
	Library  library
	Done     bool
}

func (fi fileInfo) filePath() string {
	return string(fi.Library) + fi.FileName
}

func (fi fileInfo) title() string {
	f, err := os.Open(fi.filePath())
	Check(err)
	defer f.Close()

	bf := bufio.NewReader(f)
	line, err := bf.ReadString('\n')
	if err == io.EOF {
		// single line; ignore the error
		err = nil
	}
	Check(err)
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "# ")
	return line
}

func retrieveCategoriesAndFirstNotDoneFile() map[string]fileInfo {
	ret := map[string]fileInfo{}
	for category, fs := range files {
		for _, f := range fs {
			if f.Done {
				continue
			}
			ret[category] = f
			break
		}
	}
	return ret
}

// execute `go install ./...` whenever there is an update
func main() {
	var editor string
	SelectEditor(&editor)
	flag.Parse()
	args := flag.Args()
	fileMap := retrieveCategoriesAndFirstNotDoneFile()
	// help
	// TODO: use -h
	if len(args) == 0 {
		fmt.Println("Available arguments:")
		for category, f := range fileMap {
			// 18 is the supposed maximum length of categories
			fmt.Printf("\t%-18s%s\n", category, f.title())
		}
		os.Exit(0)
	}
	// too many arguments
	if len(args) > 2 {
		panic(fmt.Sprintf("Multiple arguments (%s).\n", args))
	}
	// check category (first argument)
	_, exist := fileMap[args[0]]
	if !exist {
		panic(fmt.Sprintf("Unknown category: %s", args[0]))
	}
	// check order (second argument)
	order := 1
	if len(args) == 2 {
		tmp, err := strconv.Atoi(args[1])
		Check(err)
		// ignore invaid orders
		if tmp > 1 {
			order = tmp
		}
	}
	fs := files[args[0]]
	if len(fs) < order {
		order = 1
		fmt.Println("Out of range. Open the first one.")
	}
	// open it with external editor
	cmd := exec.Command(RenameEditor(editor), fs[order-1].filePath())
	Check(cmd.Run())
}

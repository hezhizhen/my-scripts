package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hezhizhen/tiny-tools/utilz"
)

type library string

const (
	MWeb3 library = "/Users/hezhizhen/Dropbox/MWeb3/docs/"
	MWeb2 library = "/Users/hezhizhen/Dropbox/MWeb/docs/"
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
	utilz.Check(err)
	defer f.Close()

	bf := bufio.NewReader(f)
	line, err := bf.ReadString('\n')
	utilz.Check(err)
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
	flag.Parse()
	args := flag.Args()
	fileMap := retrieveCategoriesAndFirstNotDoneFile()
	// check the number of arguments
	switch len(args) {
	case 0:
		// help
		fmt.Println("Available arguments:")
		for category, f := range fileMap {
			// 18 is the supposed maximum length of categories
			fmt.Printf("\t%-18s%s\n", category, f.title())
		}
		os.Exit(0)
	case 1:
		break
	default:
		fmt.Printf("Multiple arguments (%s). Only one argument is required.\n", args)
		os.Exit(1)
	}
	// check if arguemnt is correct
	f, exist := fileMap[args[0]]
	if !exist {
		panic(fmt.Sprintf("Unknown category: %s", args[0]))
	}
	// open it with macvim
	cmd := exec.Command("mvim", f.filePath())
	utilz.Check(cmd.Run())
}

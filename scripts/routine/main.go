package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/hezhizhen/tiny-tools/utilz"
)

type fileInfo struct {
	FileName string
	Title    string
	Done     bool
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
			fmt.Printf("\t%-18s%s\n", category, f.Title) // 18 is the supposed maximum length of categories
		}
		os.Exit(0)
	case 1:
		break
	default:
		fmt.Printf("Multiple arguments (%s). Only one argument is required.\n", args)
		os.Exit(1)
	}
	// check if arguemnt is correct
	if _, exist := fileMap[args[0]]; !exist {
		panic(fmt.Sprintf("Unknown category: %s", args[0]))
	}
	// set file path
	notePath := "/Users/hezhizhen/Dropbox/MWeb3/docs/" + fileMap[args[0]].FileName
	// open it with macvim
	cmd := exec.Command("mvim", notePath)
	utilz.Check(cmd.Run())
}

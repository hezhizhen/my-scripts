package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/hezhizhen/my-scripts/pkg/util"
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
	Related  *fileInfo
}

func (fi fileInfo) filePath() string {
	return string(fi.Library) + fi.FileName
}

func (fi fileInfo) title() string {
	f, err := os.Open(fi.filePath())
	util.Check(err)
	defer f.Close()

	bf := bufio.NewReader(f)
	line, err := bf.ReadString('\n')
	if err == io.EOF {
		// single line; ignore the error
		err = nil
	}
	util.Check(err)
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "# ")
	return line
}

func retrieveCategoriesAndWIPFiles() map[string][]fileInfo {
	ret := map[string][]fileInfo{}
	for category, fs := range files {
		for _, f := range fs {
			if f.Done {
				continue
			}
			ret[category] = append(ret[category], f)
		}
	}
	return ret
}

type category struct {
	Name  string
	files []fileInfo
}

func mapToSortedArray(raw map[string][]fileInfo) []category {
	var ret []category
	for cate, fs := range raw {
		ret = append(ret, category{cate, fs})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Name < ret[j].Name
	})
	return ret
}

// execute `go install ./...` whenever there is an update
func main() {
	var editor string
	util.SelectEditor(&editor)
	flag.Parse()
	args := flag.Args()
	filesMap := retrieveCategoriesAndWIPFiles()
	// help
	// TODO: use -h
	if len(args) == 0 {
		fmt.Println("Available arguments:")
		categories := mapToSortedArray(filesMap)
		for _, category := range categories {
			if len(category.files) == 0 {
				continue
			}
			fmt.Printf("\t%-18s%d: %s\n", category.Name, 1, category.files[0].title())
			for i := 1; i < len(category.files); i++ {
				fmt.Printf("%-26s%d: %s\n", "", i+1, category.files[i].title())
			}
		}
		os.Exit(0)
	}
	// too many arguments
	if len(args) > 2 {
		panic(fmt.Sprintf("Multiple arguments (%s).\n", args))
	}
	// check category (first argument)
	_, exist := filesMap[args[0]]
	if !exist {
		panic(fmt.Sprintf("Unknown category: %s", args[0]))
	}
	// check order (second argument)
	order := 1
	if len(args) == 2 {
		tmp, err := strconv.Atoi(args[1])
		util.Check(err)
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
	cmd := exec.Command(util.RenameEditor(editor), fs[order-1].filePath())
	if fs[order-1].Related != nil {
		cmd = exec.Command(util.RenameEditor(editor), fs[order-1].filePath(), fs[order-1].Related.filePath())
	}
	util.Check(cmd.Run())
}

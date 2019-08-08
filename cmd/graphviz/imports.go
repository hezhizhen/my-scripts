package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	repo string
)

// Usage: cd to the root of a repo and run graphviz
func main() {
	// TODO: more checks about the path
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if strings.Contains(pwd, "pkg") {
		panic("change directory to the root of the repository")
	}
	parts := strings.Split(pwd, "/")
	repo = parts[len(parts)-1]
	graph := "digraph import {\n"
	var onePackage relation
	err = filepath.Walk("pkg/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// dir: keep walking
		if info.IsDir() {
			return nil
		}
		// file
		r := readFile(path)
		// no previous package
		if onePackage.packageName == "" {
			onePackage = r
			return nil
		}
		// next package: draw last package and start handling current package
		// TODO: draw with different colors
		if onePackage.packageName != r.packageName {
			line := draw(onePackage)
			graph += line
			onePackage = r
			return nil
		}
		// same package: merge
		onePackage.imports = mergeSliceWithoutDuplication(onePackage.imports, r.imports)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// handle last package
	line := draw(onePackage)
	graph += line
	// last line
	graph += "}\n"
	// write to a file
	f, err := os.Create("imports.dot")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString(graph); err != nil {
		panic(err)
	}
}

func draw(r relation) string {
	ret := fmt.Sprintf("    \"%s\" -> {", r.packageName)
	for i, im := range r.imports {
		ret += fmt.Sprintf("\"%s\"", im)
		if i != len(r.imports)-1 {
			ret += fmt.Sprintf(", ")
		}
	}
	ret += fmt.Sprintf("}\n")
	return ret
}

func mergeSliceWithoutDuplication(a, b []string) []string {
	amap := make(map[string]bool)
	for _, key := range a {
		amap[key] = true
	}
	for _, key := range b {
		if amap[key] {
			continue
		}
		amap[key] = true
	}
	var ret []string
	for key := range amap {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret
}

type relation struct {
	packageName string
	imports     []string
}

func readFile(name string) relation {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(bs), "\n")
	ret := relation{}
	parts := strings.Split(name, "/")
	pname := strings.Join(parts[1:len(parts)-1], "/")
	ret.packageName = pname
	leftPar, rightPar := false, false

	for _, line := range lines {
		if strings.HasPrefix(line, "import") {
			// one line import
			if !strings.Contains(line, "(") {
				leftPar, rightPar = true, true
				str := retrievePackage(line)
				if str != "" {
					ret.imports = append(ret.imports, str)
				}
				continue

			}
			leftPar = true
			continue
		}
		if leftPar && !rightPar {
			// end of import
			if strings.Contains(line, ")") {
				break
			}
			// imports
			str := retrievePackage(line)
			if str != "" {
				ret.imports = append(ret.imports, str)
			}
		}
	}
	return ret
}

func retrievePackage(path string) string {
	if !strings.Contains(path, repo) {
		return ""
	}
	parts := strings.Split(path, repo+"/pkg/")
	name := parts[1]
	name = strings.TrimSuffix(name, "\"")
	return name
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/hezhizhen/my-scripts/pkg/util"
)

const (
	MWeb3 = "/Users/hezhizhen/Dropbox/MWeb3/docs/"
)

type Task struct {
	Title   string
	Due     time.Time
	Project string
	Status  string // WIP, Done
}

func main() {
	fmt.Println("Todoist based on pure text")
	fmt.Println("==========================")
	tasks := findTasks()
	for _, task := range next7days(tasks) {
		fmt.Println(task)
	}
}

func next7days(tasks []Task) []Task {
	var ret []Task
	now := time.Now()
	for _, task := range tasks {
		if task.Due.Sub(now) > time.Hour*-24 && task.Due.Sub(now) < time.Hour*24*8 {
			ret = append(ret, task)
		}
	}
	return ret
}

func today(tasks []Task) []Task {
	var ret []Task
	now := time.Now()
	for _, task := range tasks {
		if task.Due.Format("2006.01.02") == now.Format("2006.01.02") {
			ret = append(ret, task)
		}
	}
	return ret
}

func getFiles() []string {
	var ret []string
	fs, err := ioutil.ReadDir(MWeb3)
	util.Check(err)
	for _, f := range fs {
		name := f.Name()
		if strings.HasSuffix(name, ".md") {
			ret = append(ret, name)
		}
	}
	return ret
}

func findTasks() []Task {
	var ret []Task
	files := getFiles()
	for _, file := range files {
		bs, err := ioutil.ReadFile(MWeb3 + file)
		util.Check(err)
		content := string(bs)
		if strings.Contains(content, "@DoNow") {
			parts := strings.Split(content, "\n")
			var project string
			for i := range parts {
				// project name
				if i == 0 {
					title := parts[0]
					// remove prefix and labels
					title = strings.TrimPrefix(title, "# ")
					title = strings.TrimSuffix(title, "@DoNow") // TODO: how to trim all labels?
					title = strings.TrimSpace(title)
					project = title
					continue
				}
				// remove empty line
				if parts[i] == "" {
					continue
				}
				var task Task
				task.Project = project
				// add every task
				t := parts[i]
				t = strings.TrimPrefix(t, "* ")
				// status
				if strings.HasPrefix(t, "[ ] ") {
					task.Status = "WIP"
					t = strings.TrimPrefix(t, "[ ] ")
				} else {
					task.Status = "Done"
					t = strings.TrimPrefix(t, "[x] ")
				}
				// split labels
				cs := strings.Split(t, "@")
				for i := range cs {
					if i == 0 {
						task.Title = cs[i]
						continue
					}
					// labels (only date now)
					date, err := time.Parse("2006.01.02", cs[i])
					util.Check(err)
					task.Due = date

				}
				ret = append(ret, task)
			}
		}
	}
	return ret
}

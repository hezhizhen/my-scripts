package utilz

import (
	"fmt"
	"os"
)

// GetHome get home's absolute path
func GetHome() string {
	path := os.Getenv("HOME")
	if path == "" {
		user := os.Getenv("USER")
		if user == "" {
			panic("what's wrong with your environment? there's no HOME or USER")
		}
		path = fmt.Sprintf("/Users/%s", user)
	}
	return path
}

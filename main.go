package main

import (
	"errors"
	"fmt"

	"github.com/hezhizhen/tiny-tools/utilz"
)

func main() {
	fmt.Println("hello")
	err := errors.New("this is a fake error")
	defer func() {
		panicV := recover()
		if panicV != nil {
			fmt.Println("recover from panic:", panicV)
		}
	}()
	utilz.Check(err)
}

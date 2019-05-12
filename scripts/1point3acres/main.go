package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		panic("missing URL")
	}
	target := convertURL(url)
	fmt.Println("Target URL: ", target)
}

// Example
// origin: https://www.1point3acres.com/bbs/forum.php?mod=viewthread&tid=515990
// target: https://www.1point3acres.com/bbs/thread-515990-1-1.html
func convertURL(origin string) string {
	// no need to convert
	if strings.HasSuffix(origin, "html") {
		return origin
	}
	parts := strings.Split(origin, "tid=")
	tid, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	target := fmt.Sprintf("https://www.1point3acres.com/bbs/thread-%d-1-1.html", tid)
	return target
}

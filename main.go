package main

import (
	"flag"
	"strings"
	"fmt"
)

// readInput reads pattern and source string
// from command line arguments and returns them.
func main() {
	src := readInput()
	var size int
	if src == "" {
		size = 0
	} else {
		size = len(strings.Split(src, " "))
	}
	fmt.Println(size)
}

// read input from command line and return
// concatenated string
func readInput() string {
	flag.Parse()
	src := strings.Join(flag.Args(), "")
	return src
}

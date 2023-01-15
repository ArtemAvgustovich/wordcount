package main

import (
	"flag"
	"strings"
	"errors"
	"fmt"
	"os"
)

// readInput reads pattern and source string
// from command line arguments and returns them.
func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	size := len(strings.Split(src, " "))
	fmt.Println(size)
}

// read input from command line and return
// concatenated string
func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
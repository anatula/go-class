package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read args from command line, the thing to look for and the replace
	// simple words no spaces
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough arguments")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	// split by lines, one at time
	scan := bufio.NewScanner(os.Stdin)
	// reads a line
	for scan.Scan() {
		// split my string into words
		// scan.Text() gives us the actual line
		// s is a slice
		s := strings.Split(scan.Text(), old)
		// t is a string, receives slice
		t := strings.Join(s, new)
		// substitute word
		fmt.Println(t)
	}
}

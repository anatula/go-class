package main

import (
	"fmt"
	"io"
	"os"
)

// go run . *.txt > c.txt

func main() {
	// cat file1 file2 redirect concatenation of input
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// know how to read input in chunks, read file and put it stdout
		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// to avoid limit open files
		file.Close()
	}
}

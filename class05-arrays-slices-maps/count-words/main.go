package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// make a scanner reading stdin
	scan := bufio.NewScanner(os.Stdin)
	// declare a map, with storage
	words := make(map[string]int)

	// initialize scanner with splitting function
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		word := scan.Text()
		if word == "stop" {
			break
		}
		words[word]++
	}
	fmt.Println(len(words), "unique words")
	fmt.Println(words)
}

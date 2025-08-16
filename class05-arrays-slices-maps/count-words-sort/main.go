package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// words in map index by string, sort them to get most common word, sort by the value not key
	// need to extract keys and value
	type kv struct {
		key string
		val int
	}
	// create slice of kv struct
	var ss []kv

	// extract
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	//  can pass functions as arg to a func
	// function literal of a particular kind, with certain parameters
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	// print them in order
	// _ ignore that variable, get the value that was in the slice
	for _, s := range ss {
		fmt.Println(s.key, "appears", s.val, "times")
	}

}

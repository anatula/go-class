package main

import (
	"encoding/json"
	"fmt"
)

// to make them exportable
// if field is lowercase is not export
// json will not encode it
// try words
type Response struct {
	Page  int      `json:"page"`
	words []string `json:"words,omitempty"`
}

func main() {
	r := &Response{Page: 1, words: []string{"up", "in", "out"}}
	// test omitempty
	//r := &Response{Page: 1}
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
}

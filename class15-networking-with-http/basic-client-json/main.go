package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// server that return some json
const url = "https://jsonplaceholder.typicode.com"

type todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	// response can have a body
	// could be a webpage with pictures
	// we need to close it() it wont close the socket and we'll endup socket and run out
	// server program or with lots of requests
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// sometimes is convenient to do RealAll
		// body is a []byte bytes coming off the network
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		// because we know its a string, we convert it and print it
		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v\n", item)

	}

}

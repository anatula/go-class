package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8080/" + os.Args[1])

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
		fmt.Println(string(body))
	}

}

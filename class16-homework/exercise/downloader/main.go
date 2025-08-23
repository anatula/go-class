package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// we are not gonna decode the json, json object for each cartoon
// create a file, start with [ json, json, json, .. ]
// json array in a file where each of the object is metadata for 1 cartoon

// 1. We will read until we get two 404 response in a row
// 2. Each request will generate a JSON object as a string
// 3. We will bracket them with [ and ] and a , between (no comma before the first object)
// 4. The result will be a file with a JSON array of metadata objects, so we won't need to decode
// 5. We will optionally take a filename from the command line for output

const base = "https://xkcd.com"
const file = "info.0.json"

// returns the metadata for one comic by number
func getOne(i int) []byte {
	url := fmt.Sprintf("%s/%d/%s", base, i, file)
	fmt.Println(url)
	resp, error := http.Get(url)
	if error != nil {
		fmt.Fprintf(os.Stderr, "can't read: %s \n", error)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "skipping %d: got %d\n", i, resp.StatusCode)
		return nil // keep program going
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid body: %s \n", i, resp.StatusCode)
		os.Exit(-1)
	}
	fmt.Println(body)
	return body
}

func main() {
	// io.WriteCloser interface has a writer and closer, describe files
	// write to it, and possibly close it
	var output io.WriteCloser = os.Stdout
	var err error
	var fails int
	var cnt int
	var data []byte

	if len(os.Args) > 1 {
		output, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	fmt.Fprint(output, "[")
	defer fmt.Fprint(output, "]")

	for i := 1; fails <= 2; i++ {
		if data = getOne(i); data == nil {
			fails++
			continue
		}

		if cnt > 0 {
			fmt.Fprint(output, ",")
		}

		_, err := io.Copy(output, bytes.NewBuffer(data))

		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s \n", err)
			os.Exit(-1)
		}

		fails = 0
		cnt++

	}

	fmt.Fprintf(os.Stdout, "read %d comics \n", cnt)
}

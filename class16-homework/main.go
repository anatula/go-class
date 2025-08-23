package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const base = "https://xkcd.com/571/info.0.json"

func main() {
	resp, error := http.Get(base)
	if error != nil {
		fmt.Fprint(os.Stderr, error)
		os.Exit(-1)
	}

	if resp.StatusCode == http.StatusOK {
		body, error := ioutil.ReadAll(resp.Body)

		if error != nil {
			fmt.Fprint(os.Stderr, error)
			os.Exit(-1)
		}
		sb := string(body)

		fmt.Printf("%#v \n\n\n\n\n %#v", body, sb)
	}

	defer resp.Body.Close()
}

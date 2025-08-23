package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const base = "https://xkcd.com/"
const file = "/info.0.json"

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	const num = 571
	resp, error := http.Get(base + "571" + file)
	if error != nil {
		fmt.Fprint(os.Stderr, error)
		os.Exit(-1)
	}

	if resp.StatusCode == http.StatusOK {

		// Decode from stream into our struct
		var c Comic
		if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
			fmt.Fprint(os.Stderr, error)
			os.Exit(-1)
		}

		fmt.Printf("%#v \n", c)
		// Convert back to JSON with original lowercase field names
		//jsonData, _ := json.MarshalIndent(c, "", "  ")
		//fmt.Println(string(jsonData))

	}

	defer resp.Body.Close()
}

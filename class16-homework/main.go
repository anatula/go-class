package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const base = "https://xkcd.com"
const file = "info.0.json"

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

// Each HTTP response has its own Body (an io.Reader)
// Decoders are tied to a specific reader - you can't reuse a decoder across different readers
// Response bodies must be closed to avoid resource leaks

func downloadComic(num int) Comic {
	url := fmt.Sprintf("%s/%d/%s", base, num, file)
	resp, error := http.Get(url)
	if error != nil {
		fmt.Fprint(os.Stderr, error)
		os.Exit(-1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprint(os.Stderr, error)
		os.Exit(-1)
	}

	var c Comic
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()
	return c
}

func main() {
	var index = make(map[int]Comic)
	for num := 570; num <= 575; num++ {
		c := downloadComic(num)
		fmt.Println(c.Num, c.Title)
		index[num] = c
	}

	fmt.Println(index[571])
}

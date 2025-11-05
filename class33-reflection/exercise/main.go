package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response struct {
	Item   string `json:"item"`
	Album  string
	Title  string
	Artist string
}

type respWrapper struct {
	response
}

var j1 = `{
	"item": "album"
	"album": {"title": "Dark Side of the Moon"}
}`

var j2 = `{
	"item": "song"
	"song": {"title": "Bella Donna", "artist": "Stevie Nicks"}
}`

func main() {
	var resp1, resp2 respWrapper
	var err error

	if err = json.Unmarshal([]byte(j1), &resp1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp1.response)

	if err = json.Unmarshal([]byte(j2), &resp2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp1.response)
}

func (r *respWrapper) UnmarshalJSON(b []byte) (err error) {
	// can represent any json object
	// strings are the properties of the json object
	// emptyinerface is actual value of whatever it is
	var raw map[string]interface{}

	// the only thing that is gonna unmarshall into response
	// only "item" since it's the only with a json tag
	// not gonna find anything else because they are all capital names and my data has lower case

	err = json.Unmarshal(b, &r.response)
	// give me the map of string to interface and go fishing inside that based on the type of the item
	err = json.Unmarshal(b, &raw)

	switch r.Item {
	case "album":
		inner, ok := raw["album"].(map[string]interface{})
		if ok {
			album, ok := inner["title"].(string); ok {
				r.Album = album
			}
		}

	case "song":
				inner, ok := raw["song"].(map[string]interface{})
		if ok {
			title, ok := inner["title"].(string); ok {
				r.Title = title
			}
			artist, ok := inner["title"].(string); ok {
				r.Artist = artist
			}
		}
	}
	return err

}

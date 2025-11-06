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
	"item": "album",
	"album": {"title": "Dark Side of the Moon"}
}`

var j2 = `{
	"item": "song",
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

	fmt.Printf("%#v\n", resp2.response)
}

// needs this name because it's implementing the json.Unmarshaler
// func (r *respWrapper) UnmarshalJSON(b []byte) error
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
		// idem was of type "album" i went and found the data for the key in this map of string to interface
		// which gave another map of string to interface and if I find the "title" in it, the put title into r
		// which is my response wrapperbut but it embeds a response
		// so i'm really setting the album field in the response
		inner, ok := raw["album"].(map[string]interface{})
		if ok {
			if album, ok := inner["title"].(string); ok {
				r.Album = album
			}
		}

	case "song":
		inner, ok := raw["song"].(map[string]interface{})
		if ok {
			if title, ok := inner["title"].(string); ok {
				r.Title = title
			}
			if artist, ok := inner["artist"].(string); ok {
				r.Artist = artist
			}
		}
	}
	return err

}

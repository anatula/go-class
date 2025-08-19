package main

import "fmt"

func main() {
	// album variable based on anonymous struct type initialize to struct literal in the code
	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{"the white album", "the beatles", 1968, 1000000}

	//var pAlbum = &album
	//fmt.Printf("%+v", album)

	var pAlbum struct {
		title  string
		artist string
		year   int
		copies int
	}

	fmt.Println(album, pAlbum)
}

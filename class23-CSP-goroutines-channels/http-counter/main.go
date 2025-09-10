package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID int

func handler(w http.ResponseWriter, r *http.Request) {
	// Only handle the root path, ignore favicon requests
	if r.URL.Path != "/" { //favicon
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "<h1> You got %d <h1>", nextID)
	nextID++
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

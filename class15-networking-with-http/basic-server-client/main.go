package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {
	// server that return some json
	const base = "https://jsonplaceholder.typicode.com/"
	resp, err := http.Get(base + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	var item todo

	// resp.Body is io.ReadCloser
	// io.ReadCloser is an interface with methods resp.Body can pass it to NewDecoder
	// NewDecoder takes an io.Reader, is an interface, it describes where object which provides a read
	// method, so that bytes can be read from it
	// I can have a reader from a file, from a socker, or a byte buffer can be treated as a reader
	// I can pass all sorts of paramters into newDecoder that will satisfy the io.Reader interface
	// and the decoder will be able to decode by reading the bytes directly from that reader
	// resp.Body is a reader
	if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	tmpl := template.New("mine")

	tmpl.Parse(form)
	tmpl.Execute(w, item)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

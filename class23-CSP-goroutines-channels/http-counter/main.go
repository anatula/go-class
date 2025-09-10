package main

import (
	"fmt"
	"log"
	"net/http"
)

// NEW SOLUTION: turn handler method with receiver channel

// convert nextID from variable to a channel of int
// instead of increment, I read a value out of it
// i need a function with job is to start sending numbers
// var nextID int

//var nextID = make(chan int)

type nextCh chan int

// generate numbers and put it in the channel
// i can't write to channel until someone is ready to be read from it
// if not ready to be read from you block (like waiting stdin user typing)
func counter(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func (ch nextCh) handler(w http.ResponseWriter, r *http.Request) {
	// Only handle the root path, ignore favicon requests
	if r.URL.Path != "/" { //favicon
		http.NotFound(w, r)
		return
	}
	// reader can't read until somebody is ready to write
	// not a problem from the handler because the loop is not waiting for anything
	// as soon as someone "i want a the webpage with counter", try to read and counter will be able to write the next value
	// and counter will stop and handler will finish, do this over and over
	fmt.Fprintf(w, "<h1> You got %d <h1>", <-ch)
	//nextID++
}

func main() {
	// go routine, runs forever until program stops
	var nextID nextCh = make(chan int)
	go counter(nextID)
	http.HandleFunc("/", nextID.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

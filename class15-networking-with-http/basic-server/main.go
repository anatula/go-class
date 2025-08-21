package main

import (
	"fmt"
	"log"
	"net/http"
)

// This is a concurrent webserver, under the hood, the http package will handle as many connections
// as your hardware will supports. The handler is handling 1 connection: 1 client, 1 request, 1 response
// this could be running a lot of times

// take a response writer, where all the output goes
// take a request, the one that go in from a client, we can look in and see what it is may have a body
// body with json
// the job is to respond by writing stuff to the response writer that will be sent to the client
func handler(w http.ResponseWriter, r *http.Request) {
	// all this going into w
	// w is something we can send data into it
	// just write to it
	// dont set status, it will be 200 OK
	fmt.Fprintf(w, "Hello, world! from %s \n", r.URL.Path[1:])
}

// start the server
func main() {
	// bind the handler against a route
	// one handler for the top level route /
	http.HandleFunc("/", handler)
	// start a server, open a listen socket, a TCP, a socket that can accept HTTP requests
	// pass it the string :8080
	// pick the default IP address and start up on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

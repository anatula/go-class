package main

import (
	"log"
	"net/http"
	"time"
)

// data comming back from each get of the url
type result struct {
	url     string
	err     error
	latency time.Duration
}

// take url and channel we return the data on
// run each get in a go routine by itself
// go out to the web get some data and put its result in a channel and send that back
// gonna communicate its result by putting in the channel
// channel is gonna be a channel of result
// when we pass a channel to a function as a parameter we can restrict it
// restrict the use to the write end or read end
// we want our get function to return some data, so write end only
// won't be able to read any data from the channel
func get(url string, ch chan<- result) {
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Microsecond)
		ch <- result{url, nil, t}
		resp.Body.Close()

	}

}

func main() {
	// make a channel of result
	results := make(chan result)
	// 4 urls, I'm gonna start 4 go routines
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
	}
	for _, url := range list {
		// go wants a function call
		// results is channel
		// giving the write end of the channel to get function
		// that's gonna run in this go routine
		// so it has a way to communicate the data back
		go get(url, results)
	}

	// i need to read only 4 results
	// channels block
	// if there's no data to read and I go to read it, I have to wait for data
	// star a CLI program, read from stdin, expected to type something
	// if I don't type anything, it waits for me to type a line and Return or Ctrol+D to say no more data
	// range over the channel, it would read until the channel closes (I'm not closing the channel)
	// i need to make sure I don't read more times that there could be data in the channel to read or I'll get stuck
	// Also, everytime i start a go routine, its gonna give me a result: error or valid data
	// 4 go routines and 4 results and stop
	// why not close the channel? Because you can only close it once
	// have 4 go routines, who closes the channel? last person turn the lights off, no way of knowing which goroutine is the last
	// i know i started n goroutines I know i need n responses
	for range list {
		// read the channel
		// results with arrow before says "do a read on the results channel"
		// the result of that is gonna get put on r
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s \n", r.url, r.latency)
		}
	}

}

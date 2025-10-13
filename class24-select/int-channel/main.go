package main

import (
	"fmt"
	"time"
)

func main() {
	// Create 2 unbuffered channels directly in the slice
	channels := []chan int{
		make(chan int),
		make(chan int),
	}

	// Start goroutines for each channel with different speeds
	for i := 0; i < 2; i++ {
		go func(channelID int) {
			// Channel 0: 1 second, Channel 1: 2 seconds
			waitTime := time.Duration(channelID+1) * time.Second
			time.Sleep(waitTime)
			channels[channelID] <- channelID + 100 // Send unique number
		}(i)
	}

	// Use single select in a loop to receive from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg := <-channels[0]:
			fmt.Printf("Received from channel 0 (1s): %d\n", msg)
		case msg := <-channels[1]:
			fmt.Printf("Received from channel 1 (2s): %d\n", msg)
		}
	}

	fmt.Println("Program finished")
}

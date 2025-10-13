package main

import (
	"fmt"
	"time"
)

func main() {
	// Step 1: Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Step 2: Start goroutine for channel 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from channel 1"
	}()

	// Step 3: Start goroutine for channel 2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	// Step 4: Use select to receive from channels
	// This will receive from ch1 first (1 second delay)
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from channel 1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from channel 2:", msg2)
	}

	// Step 5: Receive the second message
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from channel 1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from channel 2:", msg2)
	}

	fmt.Println("Program finished")
}

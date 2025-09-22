package main

import (
	"fmt"
	"time"
)

func main() {

	// Like having a mailbox that can hold 1 item
	empty := make(chan bool, 1)
	// The sender can put data in the buffer without waiting for a receiver

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("I'm done waiting!")
		// The sender only blocks when the buffer is full
		empty <- true
	}()

	fmt.Println("Waiting....")
	//The receiver takes data from the buffer
	<-empty
	fmt.Println("Main can now exit")
}

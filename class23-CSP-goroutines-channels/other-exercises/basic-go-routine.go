package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hey from the goroutine!")
}

func delayedMessage() {
	time.Sleep(2 * time.Second)
	fmt.Println("This message was delayed!")
}

func main() {
	//go sayHello()
	// if not a goroutine it would block since channel is unbuffered and there is no receiver yet, so it hangs forever
	// for range in main never even starts because stuck
	// can be done with buffered channel
	go delayedMessage()
	fmt.Println("--START MAIN!--")
	time.Sleep(3 * time.Second)
	fmt.Println("--END MAIN!--")
}

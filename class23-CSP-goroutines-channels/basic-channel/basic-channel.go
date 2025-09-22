package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("I'm done waiting!")
		done <- true
	}()

	fmt.Println("Waiting....")
	<-done
	fmt.Println("Main can now exit")
}

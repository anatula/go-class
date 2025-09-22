package main

import "fmt"

func printMessages(ch chan string) {
	msges := []string{"msg 1", "msg 2", "msg 3"}
	for _, m := range msges {
		fmt.Println("Sending", m)
		ch <- m
	}
	close(ch) // ← THIS IS THE KEY LINE!
	fmt.Println("Channel closed")
}

func main() {
	ch := make(chan string)
	go printMessages(ch)

	for msg := range ch {
		fmt.Println("Message received", msg)
	}
	fmt.Println("All done!")
}

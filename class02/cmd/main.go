package main

import (
	hello "class02"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.Say(os.Args[1]))
	} else {
		fmt.Println("Hello, world!")
	}
}

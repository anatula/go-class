package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	in, _ := os.Open("a.txt")
	out, _ := os.Create("out.txt")
	n, _ := io.Copy(out, in)
	fmt.Println("copied", n, "bytes")
}

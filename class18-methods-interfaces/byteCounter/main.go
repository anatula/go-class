package main

import (
	"fmt"
	"io"
	"os"
)

// ByteCounter is a Writer because it has a Write method
// put it as a destination in io.Copy
// just as good as any Write
// create a "file size counting device" by opening the file
// reading its byte into ByteCounter and figure it out how many that is
type ByteCounter int

// Write method takes byte buffer and returns int
func (b *ByteCounter) Write(buffer []byte) (int, error) {
	// do copy
	l := len(buffer)
	*b += ByteCounter(l)
	return l, nil
}

func main() {
	var c ByteCounter

	in, _ := os.Open("a.txt")

	out := &c

	n, _ := io.Copy(out, in)
	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
}

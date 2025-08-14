package main

import "fmt"

func main() {
	s := "élite"
	//print type and value
	// unicode output:
	fmt.Printf("%8T %[1]v %d \n", s, len(s))
	// cast it to a sequence of runes
	// sequence of 32 bit int, i get 5 values
	// first value is 233 > 127
	fmt.Printf("%8T %[1]v \n", []rune(s))
	// force to be a sequence of bytes
	// 6 items, first two represented by 2 bytes
	// part of how UTF-8 encodes unicode
	// instead of always using uses 4 bytes, only done with values after certain size
	// chinese code point take 3 bytes to represent
	// 5 chars -> 15 bytes (each need 3)
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d \n", b, len(b))
}

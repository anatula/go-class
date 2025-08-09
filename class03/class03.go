package main

import "fmt"

func main() {
	a := 2
	b := 2.1
	// %T type of
	// %v value of
	// to align columns %8T
	fmt.Printf("a: %8T %v \n", a, a)
	// put varible once
	fmt.Printf("b: %8T %[1]v \n", b)

	a = int(b) // float to int the fractional is removed
	fmt.Printf("a: %8T %[1]v \n", a)

	b = float64(a) // float to int the fractional is removed
	fmt.Printf("b: %8T %[1]v \n", b)

	// types of a and b can't change anymore
	// fractional part goes away in convertion and lost
}

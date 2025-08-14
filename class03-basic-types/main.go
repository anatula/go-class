package main

import (
	"fmt"
	"os"
)

func main() {
	// i need them to have value 0, so default value
	// use the short declaration to init with a non-zero value
	var sum float64
	var n int

	for {
		var val float64
		// read value fromm the command line
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}
		//sum it up
		sum += val
		n++
	}

	// no numbers
	if n == 0 {
		fmt.Fprintf(os.Stderr, "no values'")
		os.Exit(-1)
	}

	fmt.Println("The average is ", sum/float64(n))

}

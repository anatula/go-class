package main

import "fmt"

type Fizgig struct{}

func (f Fizgig) Error() string {
	return "Your fizgig is bent"
}

func main() {
	// Create a Fizgig instance
	var myFizgig Fizgig

	// It already has the Error() method!
	fmt.Println(myFizgig.Error()) // Output: Your fizgig is bent

	// You can also use it as an error
	var err error = myFizgig
	fmt.Println(err.Error()) // Output: Your fizgig is bent
}

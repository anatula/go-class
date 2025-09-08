package main

import "fmt"

// 1. Define a type for our function signature
type Greeter func(string)

// 2. Attach a method to that type
func (g Greeter) GreetWithStyle(name string) {
	fmt.Println("*** Awesome Greeting ***")
	g(name) // Call the original function
	fmt.Println("*************************")
}

// 3. A normal function that matches the Greeter signature
func simpleGreet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	// 4. Convert the function to the Greeter type
	myGreeter := Greeter(simpleGreet)

	// 5. You can call it as a function...
	myGreeter("Alice") // Output: Hello, Alice!

	// ...OR you can call the method you attached to it!
	myGreeter.GreetWithStyle("Bob")
	// Output:
	// *** Awesome Greeting ***
	// Hello, Bob!
	// *************************
}

## class21 Homework

### Understanding Go HandlerFunc Pattern Explained

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```
- to satisfy this interface, a type must have a method named exactly `ServeHTTP` with the exact signature `(ResponseWriter, *Request)`

```go
type HandlerFunc func(ResponseWriter, *Request)
```
- This creates a new type called `HandlerFunc` that is a *function type*. It's not a function itself, but a type that describes functions with the signature `func(ResponseWriter, *Request)`.

```go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}
```

- Adding the `ServeHTTP` method to the `HandlerFunc` type, any function of type `HandlerFunc` now **automatically** satisfies the `Handler` interface. 
- `HandlerFunc` is an adapter that lets ordinary functions behave as `Handler` objects.
- The Go built-in webserver is looking for things that have the `Handler` interface that do the `ServeHTTP` method and we can provide various function that do that.
- We're going to show how METHOD VALUE can be used as a handler (a METHOD VALUE closes over the receiver) then it becomes a plain function (regular parameters and its parameter list, once it's closed over the receiver)
- So, if we write a method against some type that takes a `ResponseWriter` and a `Request` we can turn that into a handler. 

```go
// Your ordinary function
func myHandler(w ResponseWriter, r *Request) {
    w.Write([]byte("Hello!"))
}

// This works because:
// 1. myHandler gets converted to HandlerFunc(myHandler)
// 2. HandlerFunc has ServeHTTP method, so it implements Handler interface
// 3. The http server can call ServeHTTP on it, which calls your original function

http.HandleFunc("/", myHandler) // HandleFunc converts your function
```

- This pattern is elegant because it lets you write **simple functions** instead of having to create structs that implement interfaces. The `HandlerFunc` type acts as a bridge between the interface world and the function world.

#### Function type

A function type is a type definition that describes a function's signature (its parameters and return values). It allows you to treat functions as first-class citizens that can be assigned to variables, passed as arguments, and returned from other functions—but with the added safety and clarity of a specific type.

- Without function type:
```go
func Process(data []byte, callback func([]byte, error) int) {
    // ... what does this callback do?
}
```

- With function types:

```go
// WriteHandler is a function that processes a data chunk and
// returns the number of bytes consumed or an error.
type WriteHandler func([]byte, error) int

func Process(data []byte, callback WriteHandler) {
    // ... Ah, `callback` is a WriteHandler!
}
```

#### Another function type example

```go
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
```

#### Making Greeter Satisfy an Interface

```go
package main

import "fmt"

// 1. Define an interface that requires a Greet method
type Greeting interface {
    Greet(name string)
}

// 2. Define a function type that matches the interface method signature
type Greeter func(string)

// 3. Attach the required method to the function type
// This is the magic that makes Greeter satisfy the Greeting interface
func (g Greeter) Greet(name string) {
    g(name) // This calls the original function
}

// 4. Some ordinary functions that match the Greeter signature
func simpleGreet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func formalGreet(name string) {
    fmt.Printf("Good day, Mr./Ms. %s\n", name)
}

// 5. A function that accepts anything that satisfies the Greeting interface
func processGreeting(g Greeting, name string) {
    fmt.Print("Processing greeting: ")
    g.Greet(name) // Can call the Greet method on the interface
}

func main() {
    // 6. Convert our ordinary functions to the Greeter type
    simple := Greeter(simpleGreet)
    formal := Greeter(formalGreet)
    
    // 7. We can use them as functions directly
    fmt.Println("-- As functions --")
    simple("Alice")
    formal("Bob")
    
    // 8. OR we can use them as Greeting interfaces!
    fmt.Println("\n-- As Greeting interface --")
    processGreeting(simple, "Charlie")
    processGreeting(formal, "Diana")
    
    // 9. We can also call the method directly
    fmt.Println("\n-- Calling method directly --")
    simple.Greet("Eve")
    formal.Greet("Frank")
}
```
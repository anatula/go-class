## class17 OOP

![](./img/oop1.png)

### Abstraction

- logical notion of how we make things simpler
- We abstract the notion of file from all those underlying details

![](./img/abstraction.png)


### Encapsulation

![](./img/encapsulation.png)


- I can make that work without encapsulation
- Simplify by hiding information, user doesn't knoe the details
- you can call the functions i provide, you can't hack the internals

### Polymorphism

![](./img/polymorphism.png)

#### ad-hoc
- Same function name, different implementations based on type
```go
package main

import "fmt"

// Ad-hoc polymorphism: same function name works with different types
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	MakeSound(Dog{}) // Works with Dog
	MakeSound(Cat{}) // Works with Cat
}

```
#### Generics
-  Same code works with different types (generics)
```go
package main

import "fmt"

// Parametric polymorphism: same function works with any type
func PrintSlice[T any](slice []T) {
	for _, item := range slice {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func main() {
	intSlice := []int{1, 2, 3}
	stringSlice := []string{"hello", "world"}

	PrintSlice(intSlice)    // Works with ints
	PrintSlice(stringSlice) // Works with strings
}
```

#### subtype
- Different types satisfy the same interface contract

```go
package main

import "fmt"

// Subtype polymorphism: different types satisfy the same interface
type Animal interface {
	Move() string
}

type Bird struct{}

func (b Bird) Move() string { return "Flying" }

type Fish struct{}

func (f Fish) Move() string { return "Swimming" }

func DescribeMovement(a Animal) {
	fmt.Println("The animal is", a.Move())
}

func main() {
	DescribeMovement(Bird{}) // Bird satisfies Animal interface
	DescribeMovement(Fish{}) // Fish satisfies Animal interface
}
```
### Inheritance

![](./img/inheritance.png)
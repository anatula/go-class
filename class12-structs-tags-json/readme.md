## class12
- ***most structs are going to be given type names***
```go
// %T	a Go-syntax representation of the type of the value
// %v	the value in a default format
// when printing structs, the plus flag (%+v) adds field names
// [1] refers to positional indexing of arguments. It means "use the first argument again
fmt.Printf("%T %+[1]v \n", e)
```

Rule of thumb
- Types use *T
- Expressions use &x

This does not work:
`employee/main.go:22:12: invalid operation: cannot take address of c["Lamine"] (map index expression of struct type Employee)`

- map of strucs, limitation, you cannot take the address of a map entry
- when i do key-value lookup, look up a name, get Employee record value, can't take his address. Every time I insert/delete from the map, the map can rearrenge itself, because internally is a dynamic hashtable. It's unsafe, adress can be bogus, my pointer could become stale.
- cant' take the address of a map entry
- can't do something to a value inside that structure
- **So, almost always going to see a map of string to struct pointer**


```go
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	c := map[string]Employee{} // <------- remove pointer won't work

	c["Lamine"] = Employee{"Lamine", 2, nil, time.Now()}
	c["Lamine"].Number++ // <----------- this won't work
	c["Matt"] = Employee{
		Name:   "Matt",
		Number: 1,
		Boss:   &c["Lamine"], // <----------- this won't work 
		Hired:  time.Now(),
	}

	fmt.Printf("%T %+[1]v \n", c["Lamine"])
	fmt.Printf("%T %+[1]v \n", c["Matt"])

}
```

Anonymous struct
```go
// This creates a VARIABLE with an anonymous struct type
var album struct {
    title  string
    artist string
    year   int
    copies int
}

// album is a VARIABLE with an anonymous type
// You can ONLY create this one variable
// Cannot create more variables of this exact type
```
Named Struct
```go
// This creates a TYPE definition (no variable yet)
type Album struct {
    title  string
    artist string
    year   int
    copies int
}
// Album is a TYPE, not a variable
// Now you can create MANY variables of this type
```
Pointer to struct
```go
var pAlbum *struct {
    title  string
    artist string
    year   int
    copies int
}
// Creates pointer (initially nil)
// Only allocates memory for the pointer, not the struct itself
// Must allocate memory: pAlbum = &struct{...}{}
// Use for: Large structs or shared modification
```

### Struct tag
- A struct tag is a string after the field type and has some info on how to encode in various protocols
- key for json, use reflection, so info in that tag is avaiable at runtime for json to know how to encode/decode
- mostly used for converstion to external format json, xml, protobuffers (move data to outside)
### Struct compatibility
- the fields have the same types and names 
- in the same order
- and the same tags

- Can I compare structs? If all its fields are comparable (byte slices are not comparables)
- initialization is zero value of all fields

- named types can be converted, if they have the same field i can convert to one to the other

- structs are passed by value unless a pointer is used (so its the original)
- "dot notation" `(*a).copies`

## Empty structs
- struct{} are singleton
### Set type (instead of bool)
- struct{} has zero memory overhead
- Both have similar lookup performance O(1)
- Better for large sets due to memory savings

```go
// bool version - each value takes 1 byte
boolSet := make(map[int]bool)    // 1 byte per value

// struct{} version - each value takes 0 bytes  
structSet := make(map[int]struct{}) // 0 bytes per value
```

### What is a Channel?
- A channel is Go's way for goroutines (concurrent functions) to communicate with each other. Think of it like a pipe where one goroutine can send data and another can receive it.

#### The "Cheap Signal" Pattern
```go
// Instead of expensive boolean channels:
done := make(chan bool)    // 1 byte per signal

// Use zero-memory cheap signal:
done := make(chan struct{}) // 0 bytes per signal!
```
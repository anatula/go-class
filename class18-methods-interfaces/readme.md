## class18 Methods and interfaces

- an interface is an specification of abstract behaviour, list a methods that a concrete object must provide
- a concrete offer methods that satisfy the interface
- in go we don't have classes and don't define method in them, we define them seperately
- a method is a function that also has a **receiver** (specified before the function name)

```go

type IntSlice []int

// the receiver is is (type IntSlice)
func (is IntSlice) String() string {
    strs []string

    for _, v:= range is {
       strs = append (strs, strconv.Itoa(v))
    }

    return "[" + strings.Join(strs, ";") + "]"
}

```

- in go, we can put methods on any user declared type 
- print function, is it a string? copy the string to the output, is it a stringer? call the string() method generate the string and put it in on the output.
- Interfaces are Contracts, not Types. Think of an interface not as a thing but as a set of rules or a contract. The fmt.Stringer interface is a contract that says: "Any type that has a method String() string automatically fulfills this contract and can be used wherever a fmt.Stringer is expected."

When you declare var s fmt.Stringer without assignment:
```
var s fmt.Stringer
fmt.Printf("%T: %[1]v \n", s)
```

- Both components are nil
- There's no concrete type associated with it
- There's no concrete value stored
- This is different from a non-interface type where %T would always show the static type:


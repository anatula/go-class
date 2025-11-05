## class 33 Reflection
- Watched class 42 to understand parametric polymorphism (still need to finish it when the times comes)
- `interface{}` has no method, does not define not restrict any behaviour, it can represent any type in a program
- `sync.pool` container type
- generics type uses generics or empty interface, we needed to get data out of that, we did "downcasting" "type assertion", take generic thing and make it more specific thing
- `f, ok := w.(*os.File)` if fail by panicking, check if f is file pointer and ok false
- Reflection? a program looking at itself. C or older languages, compiler know types it uses that info to translate 
- once program is converted to machine code, all that type info is lost
- Reflection is the notion that when the program is done being compiled the compiler can leave behind type info in the executable
- literal string will appear in the binary, we'll get a bunch of type info in the binary
- we have utilities to look that type info, give some stuff about it
- slices can't be compared with == operator
- `reflect` package has `reflect.DeepEqual(got, want)` use reflection to check if they are equal
- reflection used json decoding and printing
- `...` meanings: 
    - allows a function to accept any number of arguments of the same type `func foo(values ...Type)`
    - let the compiler determine the array length automatically `[...]Type{1, 2, 3}`
    -  "unpacks" a slice into individual arguments `slice...`
    ```go 
    func printNumbers(a, b, c int) {
        fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)
    }

    numbers := []int{10, 20, 30}
    printNumbers(numbers...) // a=10, b=20, c=30
    ```
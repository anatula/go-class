## class 42 Parametric Polymorphism (Generics)
- Why? We parameterize types
- We put a Type Parameter on a type
- I can't do anything until I specify what T is
- `any` is an alias for empty interface `interface{}`
- `any` says "we are not putting a constraint on type T"
```go
type MyType [T any] struct {
    v T
    n int
}
```
- A template for creating real types, `MyType` is abstract not concrete
- Example, sync.Pool is a safe container in a concurrency sense, manages a pool of object, manages them as Empty Interface, everytime you get something out of a pool you have to take it from an empty interface and "downcast it" a type assertion to make it be a particular concrete type (hopefully) the type that you're suppose to get.
- This is a case of replacing dynamic typing with static typing, this makes the program safer because we doon't figure out at runtime, move that type checking into compile time
- go generics -> normal go (at the time)
- method example
```go
func Print[T any](s []T) {
    for _, v:= range s {
        fmt.Print(v)
    }
}

func main () {
    Print([]string{"Hello, ", "playground\n"})
}
```

- Continue in https://youtu.be/Si0rAE8yT9g?si=goG9nSrrUcds9QMe&t=614

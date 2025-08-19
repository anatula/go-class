## class12

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
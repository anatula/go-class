## class 32 Error
- Error is just a string, wrapped in an some internal object
- print it for debugging
- Error is an interface, anything that has an error method, so we can create other concrete types that can represent this

```go
type error interface {
	Error() string
}
```


```go
type Fizgig struct {}

func (f Fizgig) Error() string {
	return "Your fizgig is bent"
}
```
- Create a more useful Error
- Error variables, each is instance of struct
- custom error, taking error, getting some information and returning 
- Wrapped errors: wrap one error in another, "chain" TOP LEVEl (top-level error -> intermediate error -> original error)
- Unwrapepd method to return the underlying error, if NIL bottom of the chain
- `errors.Is` (compares with error variable) check if error has another error in its chain
-  `errors.As` (compares with error Type) go the error chain and extract the error of this type.  Think of it as downcasting, given an error, can i downcast it to an error of this type (may not be the top, in the chain)

### Errors in Go
- Normal errors: edge condition caused by input (user or network and get a response)
- External conditions: is the network available?out of memory?
- 
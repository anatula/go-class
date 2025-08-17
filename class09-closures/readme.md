## class09
- closure is about functions that lives inside functions, and refer to the enclosing function's data
- scope is static, based on the code at compile time (your code, the text)

### Stack, Heap, and Escape Analysis

In Go, memory management is handled automatically, but it helps to understand how it works.

### Stack
- Region of memory attached to each goroutine.
- Very fast to allocate and free (just moves a pointer).
- Variables normally live here if they are only used inside the function.
- Memory is freed automatically when the function returns.

```go
func add(a, b int) int {
    sum := a + b // sum lives on the stack
    return sum
}
```

### Heap
- A shared memory region for data that may live beyond the current function.
- Slower than stack (requires garbage collection).
- Variables are placed here if they "escape" the function, e.g., returned as a pointer or used by another goroutine.

```go
func makePointer() *int {
    x := 42
    return &x // x is moved to the heap
}
```

### Escape Analysis
Go compiler decides at compile time where a variable should live (stack or heap).
Go approach is to allocate as much as possible in the stack. In the heap, need to use pointers to it and it need to be garbage collector.

-  If a variable is only used inside a function → stack.
- If a variable must survive after function returns or shared across goroutines → heap.

### Closure
- is a ***runtime thing***, like a string descriptor, an invisible part you don't deal.
- I can have a function that need to take certain params, I'm gonna create a closure because I want to pass it to param to some other function. Maybe the signature that I have to pass is already determined. Takes 2 ints, but i want that function to have a value out of the context in which is called. I can close over a local variable, in addition to the params going into the function.


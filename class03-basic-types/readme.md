## class03

### Machine-native vs interpreted

`a := 2`
- `a` variable is the name or address of a memory location in the machine (RAM)
- no interpreter, no JVM
- dealing direcly with machine memory, no abstraction
- go is taking the source, compiler is generating machine code

- `int` numbers (2 is an int so a is an int)
- 64 bit machines -> 8 byte number (big)
- float are non-integers, default is float64
- don't use floating point for money

### Variable declaration

- `var a int` (var keyword)
```
    vat (
        b = 2
        f = 2.01
    ) // variable group, go infers the type
```

- `c := 2` (only inside in functions) SHORT DECLARATION OPERATOR (c has value 2 and type int)

```
x := 5   // declares x and assigns 5
x = 10   // assigns a new value to existing x
```
If used just `=` for both, the compiler wouldn't know whether you wanted to declare a new variable or just update an existing one.

So, the : in := signals:
➡️ "This is a declaration of a new variable"

#### boolean
-  `bool` true/false logical value
- ⚠️ NOT convertible to/from integers!

#### errors
- `error` special type with one function Error()
- an error may be `nil` or `non-nil`

#### pointers
- address of something
- may be nil (doesn't point at anything) or non-nil 
- do things with pointers using `unsafe` package

### Initialization

Go initializes all variables to 'zero' by default:
- numerical types 0
- bool false
- string "" empty string length 0
- pointers, slices, maps, channels, functions (function variables), interfaces -> nil
- for aggregate type, all members get their 'zero'

### constants

Go is a concurrent language. Only numbers, strings and booleans can be constants (IMMUTABLE)
```
const (
    a = 1
    s = "a string"
    t = len(s)
)
```
To run average example:
`go run . < nums.txt`
`cat nums.txt | go run .` 
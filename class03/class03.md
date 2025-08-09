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
- ```
    vat (
        b = 2
        f = 2.01
    ) // variable group, go infers the type
``` 
- `c := 2` (only inside in functions) SHORT DECLARATION OPERATOR (c has value 2 and type int)
- 
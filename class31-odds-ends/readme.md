## class 31 odds and ends

### Enumerated types

-  no keyword like enum
- `iota` magically increments, starts at 0
- constant block of `type` shoe int
- << bit shift operator
- 2^1, 2^2, 2^3, .. successive values
- put an _ underscore to ignore the first value
- use more complex iota expressions

###  Variable argument list
- when we do printing, pass != number of arguments
- `...` new operator to a paramater says: "Variable number of characters"
- only the last parameter in the parameter list
- `func sum (nums ...int) int`, can do `range nums`, nums will be treated as a slice
-  `s...` I can pass a slice s as a variable number of integer, means "unpack"
- 

### Sized integers
- almost always use type int (is a 64 bit, all stars in the sky)
- uint variation
- uint16, uint32, uint8 to handle low-level protocols like TCP
- small integers is gonna have weird behiviour

### Goto statement
- jump to specific part of the code, transferring control to a labeled statment whitin the same function
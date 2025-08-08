package hello

import "fmt"

// takes a string
// returns a string
// parameter: varname type (type comes after)
func Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

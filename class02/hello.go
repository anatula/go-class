package hello

import "strings"

// takes a string
// returns a string
// parameter: varname type (type comes after)
//func Say(name string) string {
//fmt.Print(greeting)
//	return fmt.Sprintf("Hello, %s!", name)
//}

func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}
	return "Hello, " + strings.Join(names, ", ") + "!"
}

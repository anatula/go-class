package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

// this is a method: function attached to the IntSlice type
// The (is IntSlice) part b4 the function name is called the "receiver"
// the receiver is is (type IntSlice)
// defines which type this function belongs to
func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	// var s of type fmt.Stringer
	// fmt.Stringer is an interface
	// s is an interface variable
	// i can assign to it anything that satifies the interface
	// any actual type that has a String method
	// AT RUNTIME TYPE OF THE THING THE INTERFACE IS HOLDING ON
	// interface is a reference to abstract behaviour being satisfied by actual object, in this case v
	var s fmt.Stringer = v // (actual object v)

	fmt.Printf("%T: %[1]v \n", s)

	for i, x := range v {
		fmt.Printf("%d: %d \n", i, x)
	}
	fmt.Printf("%T: %[1]v \n", s)
	fmt.Printf("%T: %[1]v \n", v)
}

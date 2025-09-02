package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IntSlice []int

// the receiver is is (type IntSlice)
func (is IntSlice) String() string {
	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}

func main() {
	var v IntSlice = []int{1, 2, 3}
	var s fmt.Stringer = v // (actual object v)

	fmt.Printf("%T: %[1]v \n", s)

	for i, x := range v {
		fmt.Printf("%d: %d \n", i, x)
	}
	fmt.Printf("%T: %[1]v \n", s)
	fmt.Printf("%T: %[1]v \n", v)
}

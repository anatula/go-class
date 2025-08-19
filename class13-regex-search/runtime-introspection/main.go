package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(A())
}

func A() any {
	return B()
}

func B() any {
	// Returns: pc (program counter), file, line, ok
	// skip = 1 means "go one level up the call stack"
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file)
	fmt.Println(line)
	// LastIndexByte: Looks for single byte '/'
	idx := strings.LastIndexByte(file, '/')

	//LastIndex:Looks for string "/" (more expensive)
	//idx := strings.LastIndex(file, "/")
	return "=>" + file[idx+1:] + ":" + strconv.Itoa(line)
}

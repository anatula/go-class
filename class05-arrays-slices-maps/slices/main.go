package main

import "fmt"

func main() {
	t := []byte("string")            // [115 116 114 105 110 103]
	fmt.Println(len(t), t)           // len 6
	fmt.Println(t[2])                // index 2 -> 114
	fmt.Println(t[:2])               // 2-0=2 -> index 0, 1 [115, 116]
	fmt.Println(t[2:])               // 2-0=2 -> from index 2 [114,..,103 ]
	fmt.Println(t[3:5], len(t[3:5])) // 5-3=2 index 3 and 4 [105, 110] - len 2
}

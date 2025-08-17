package main

import "fmt"

// func do(b []int) int {
// 	b[0] = 0
// 	fmt.Printf("%p \n", b)
// 	return b[1]
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	fmt.Printf("%p \n", a)
// 	v := do(a)
// 	fmt.Println(a, v)
// }

func do1(b map[int]int) {
	// this will also change a
	b[3] = 3
	// this is just a local variable in the function
	b = make(map[int]int)
	// totally different from a, no relation
	b[4] = 4
	fmt.Println("local b", b)
	//this map goes away
}

func do2(b *map[int]int) {
	// b is a pointer, we need to de-refered it to get the original thing that we're store into
	// need to dereference the pointer first
	(*b)[3] = 3
	// this is just a local variable in the function
	*b = make(map[int]int)
	// totally different from a, no relation
	(*b)[4] = 4
	fmt.Println("local b", *b)
	//this map goes away
}

// when i pass a map by itself, the map descriptor gets copied but i have a reference to the original hashtable
// in the next case, i'm passing the address of the map descriptor, so when overwrite it on line 34
// i'm changing the map descriptor also from main
// when function is done b is same as a, a is gone

func main() {
	a := map[int]int{4: 1, 7: 2, 8: 3}
	//do1(a)
	fmt.Println("main a", a)
	// & address of
	do2(&a)
	fmt.Println("main a", a)
}

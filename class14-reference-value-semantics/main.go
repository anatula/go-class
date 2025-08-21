package main

import "fmt"

func main() {
	// slice of 2 element arrays
	// type is array of 2 bytes (could also be int)
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	// make this into a slice of slice of bytes
	// create a slice from each from the array entries in items and put them into a
	a := [][]byte{}

	// item is a copy at particular location
	// copy of 1,2  3,4,  5,6
	// are in the same location
	// my slice is picking a pointer
	// has the last thing in the
	// should have been this [[5,6] [5,6] [5,6]]
	// all 3 slice point to same array with values 5,6
	// keeping a reference to the variable given by range using later after the loop
	// they all refer to final value
	for _, item := range items {
		i := make([]byte, len(item))
		copy(i, item[:]) // make unique
		a = append(a, i)
	}
	fmt.Println(items)
	fmt.Println(a)
}

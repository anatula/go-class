package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Organs }
type ByWeight struct{ Organs }

// specialize the Less method

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	// 	s := Organs{Organ{"brain", 1340}, Organ{"liver", 1494}, Organ{"spleen", 162}, Organ{"pancreas", 131}, Organ{"heart", 290}}
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}

	fmt.Println("Original:", s)

	// Sort by name (ascending)
	sort.Sort(ByName{s})
	fmt.Println("Sorted by name (ascending):", s)

	// Sort by name (descending) using Reverse
	sort.Sort(sort.Reverse(ByName{s}))
	fmt.Println("Sorted by name (descending):", s)

	// Sort by weight (ascending)
	sort.Sort(ByWeight{s})
	fmt.Println("Sorted by weight (ascending):", s)

	// Sort by weight (descending) using Reverse
	sort.Sort(sort.Reverse(ByWeight{s}))
	fmt.Println("Sorted by weight (descending):", s)

}

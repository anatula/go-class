package hello

import (
	"fmt"
	"testing"
)

// start with Test
// always take a pointer to a testing.T
// they don't return anything
// to fall call something on t to make it happen

// func TestSayHello(t *testing.T) {
// 	want := "Hello, test!"
// 	got := Say("test")

// 	if want != got {
// 		t.Errorf("wanted %s, got %s", want, got)
// 	}

// }

func TestSayHello(t *testing.T) {
	// create a slice of structs
	// create it on the fly, didn't give it an actual type name
	// create the first subtest, don't give it all the fields (items is empty to test default behaviour)
	// various cases to test (not actual subtests)
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Ana"},
			result: "Hello, Ana!",
		},
		{
			items:  []string{"Ana", "Matt"},
			result: "Hello, Ana, Matt!",
		},
	}

	// loop over the subtests, one at a time
	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			fmt.Printf("DEBUG: actual result '%s' of items (%v) with expected result: '%s'", s, st.items, st.result)
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}

}

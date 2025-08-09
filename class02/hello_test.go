package hello

import "testing"

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
	want := "Hello, test!"
	got := Say([]string{"test"})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}

}

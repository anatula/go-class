package main

import "fmt"

type errKind int

// create enumerated type for the kinds of error, use them as tags inside a struct
const (
	_ errKind = iota
	noHeader
	cantReadHeader
	invalidHdrType
)

type WaveError struct {
	kind  errKind
	value int
	err   error
}

func (e WaveError) Error() string {
	switch e.kind {
	case noHeader:
		return "no header (file too short?)"

	case cantReadHeader:
		return fmt.Sprintf("can't read header[%d]: %s", e.value, e.err.Error())

	case invalidHdrType:
		return "invalid header type"
	}

}

func main() {
	fmt.Println("heyyy!!")
}

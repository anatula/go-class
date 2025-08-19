package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	b := Employee{"Lamine", 2, nil, time.Now()}
	e := Employee{
		Name:   "Matt",
		Number: 1,
		Boss:   &b,
		Hired:  time.Now(),
	}

	fmt.Printf("%T %+[1]v \n", e)

}

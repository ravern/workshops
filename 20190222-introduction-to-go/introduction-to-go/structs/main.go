package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	// This is the syntax to create a struct.
	myPerson := person{
		name: "John",
		age:  20,
	}

	fmt.Println(myPerson)
}

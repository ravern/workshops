package main

import "fmt"

func main() {
	myOriginalInt := 3

	var (
		myIntPtr *int
		myInt    int
	)

	myIntPtr = &myOriginalInt
	myInt = *myIntPtr // `*` returns the underlying value of a pointer.

	fmt.Println("*myIntPtr:  ", *myIntPtr)
	fmt.Println("myInt:      ", myInt)
}

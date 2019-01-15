package main

import "fmt"

func main() {
	var (
		myInt    int
		myIntPtr *int
	)

	myInt = 3
	myIntPtr = &myInt // `&` returns the memory address

	fmt.Println("&myInt:   ", &myInt)
	fmt.Println("myIntPtr: ", myIntPtr)
}

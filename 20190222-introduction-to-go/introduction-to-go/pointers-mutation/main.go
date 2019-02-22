package main

import "fmt"

func main() {
	myOriginalInt := 3

	myInt := 3
	myIntPtr := &myOriginalInt

	modifyValues(myInt, myIntPtr)

	fmt.Println("myInt:     ", myInt)
	fmt.Println("*myIntPtr: ", *myIntPtr)
}

// modifyValues attempts to change the values of the parameters passed in.
func modifyValues(myInt int, myIntPtr *int) {
	myInt = 10
	*myIntPtr = 10
}

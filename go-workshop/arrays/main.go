package main

import (
	"fmt"
)

func main() {
	var myArray [4]int

	myArray[1] = 3
	myArray[3] = 23

	myOtherArray := [4]int{0, 3, 0, 23}

	fmt.Println("myArray:      ", myArray)
	fmt.Println("myOtherArray: ", myOtherArray)
}

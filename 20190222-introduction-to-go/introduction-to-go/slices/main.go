package main

import (
	"fmt"
)

func main() {
	var mySlice []int

	mySlice = make([]int, 4)
	mySlice[0] = 325
	mySlice[2] = 64

	myOtherSlice := []int{325, 0, 64, 0}

	fmt.Println("mySlice:         ", mySlice)
	fmt.Println("myOtherSlice:    ", myOtherSlice)

	myAppendedSlice := append(mySlice, 78)

	fmt.Println("myAppendedSlice: ", myAppendedSlice)
}

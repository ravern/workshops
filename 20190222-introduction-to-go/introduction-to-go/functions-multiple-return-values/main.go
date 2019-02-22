package main

import "fmt"

func main() {
	x, y := 45, 83
	z, w := myFunction(x, y)

	fmt.Printf("myFunction(%d, %d) = (%d, %d)\n", x, y, z, w)
}

// myFunction swaps two ints together.
func myFunction(x int, y int) (int, int) {
	return y, x
}

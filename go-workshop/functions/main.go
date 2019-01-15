package main

import "fmt"

func main() {
	x := 45
	y := 83
	z := myFunction(x, y)

	fmt.Printf("myFunction(%d, %d) = %d\n", x, y, z)
}

// myFunction adds two ints together.
func myFunction(x int, y int) int {
	return x + y
}

package main

import (
	"fmt"
)

func main() {
	myInt := 10

	if myInt < 20 {
		fmt.Println("myInt is less than 20.")
	} else if myInt > 40 {
		fmt.Println("myInt is more than 40.")
	} else {
		fmt.Println("myInt is between 19 and 41.")
	}
}

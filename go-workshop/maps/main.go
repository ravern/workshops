package main

import (
	"fmt"
)

func main() {
	var myMap map[string]string

	myMap = make(map[string]string)
	myMap["one"] = "two"
	myMap["three"] = "four"

	myOtherMap := map[string]string{"one": "two", "three": "four"}

	fmt.Println("myMap:         ", myMap)
	fmt.Println("myOtherMap:    ", myOtherMap)

	myMap["five"] = "six"

	fmt.Println("myMap (again): ", myMap)
}

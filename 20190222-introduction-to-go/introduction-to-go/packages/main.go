package main

import (
	// Replace this with the path to your package relative to GOPATH.
	// For most of you, it should be just "packages/mypackage".
	"github.com/ravernkoh/go-workshop/go-workshop/packages/mypackage"
)

func main() {
	mypackage.printHello()
	mypackage.PrintWorld()
}

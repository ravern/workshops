package interfaces

type shape interface {
	calculateArea() int
}

// Any type that has the method `calculateArea` can be passed into this function.
func calculatePrice(pricePerUnit int, s shape) int {
	return pricePerUnit * s.calculateArea()
}

func main() {
	mySquare := square{size: 10}
	myRectangle := rectangle{width: 10, height: 10}
	myCircle := circle{radius: 10}

	calculatePrice(100, mySquare)
	calculatePrice(100, myRectangle)
	calculatePrice(100, myCircle)
}

package square

type square struct {
	size int
}

func (s square) calculateArea() int {
	return s.size * s.size
}

func calculateSquarePrice(pricePerUnit int, s square) int {
	return pricePerUnit * s.calculateArea()
}

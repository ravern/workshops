package circle

import (
	"math"
)

type circle struct {
	radius int
}

func (c circle) calculateArea() int {
	return int(math.Pi * float32(c.radius) * float32(c.radius))
}

func calculateCirclePrice(pricePerUnit int, c circle) int {
	return pricePerUnit * c.calculateArea()
}

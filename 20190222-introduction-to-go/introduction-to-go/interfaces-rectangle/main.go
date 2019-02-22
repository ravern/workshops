package rectangle

type rectangle struct {
	width  int
	height int
}

func (r rectangle) calculateArea() int {
	return r.width * r.height
}

func calculateRectanglePrice(pricePerUnit int, r rectangle) int {
	return pricePerUnit * r.calculateArea()
}

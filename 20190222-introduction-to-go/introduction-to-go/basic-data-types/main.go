package main

var (
	// Numeric
	_ int     = 134
	_ int8    = 32
	_ int16   = 0x436
	_ int32   = 0x32425235
	_ int64   = 235592375823748
	_ float32 = 324.32
	_ float64 = 23523.322

	// Strings
	_ byte   = 'x'
	_ rune   = 'y'
	_ string = "some string"

	// Boolean
	_ bool = false
)

func main() {}

package main
import "math"

func Square (x float32) float32{
	return x * x
}
func Circle (radius float32) float32{
	return math.Pi * radius * radius
}

func cube( x  float32) float32{
	return x * x * x
}
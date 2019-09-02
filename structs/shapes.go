package geometry

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// Convention in Go to have the reciever variable to be
// the first letter of the type
// recieverName is analogous to "this"
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

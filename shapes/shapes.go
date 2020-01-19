package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	perimeter = 2 * (rectangle.Width + rectangle.Height)
	return
}

func (rectangle Rectangle) Area() (area float64) {
	area = rectangle.Width * rectangle.Height
	return
}

func (c Circle) Perimeter() (perimeter float64) {
	perimeter = math.Pi * c.Radius
	return
}

func (c Circle) Area() (area float64) {
	area = math.Pi * c.Radius * c.Radius
	return
}

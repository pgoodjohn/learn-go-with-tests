package shapes

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

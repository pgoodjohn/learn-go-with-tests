package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	perimeter = 2 * (width + height)
	return
}

func Area(rectangle Rectangle) (area float64) {
	area = width * height
	return
}

package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(width float64, height float64) (perimeter float64) {
	perimeter = 2 * (width + height)
	return
}

func Area(width float64, height float64) (area float64) {
	area = width * height
	return
}

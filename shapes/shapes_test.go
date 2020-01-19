package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	assertShapePerimeter := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Perimeter()

		if expected != actual {
			t.Errorf("Expected: %.2f, actual: %.2f", expected, actual) 
		}
	}

	t.Run("perimeter of rectangle is 2b + 2h", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		expected := 40.0

		assertShapePerimeter(t, rectangle, expected)
	})

	t.Run("perimeter of a circle is pi * r", func(t *testing.T) {
		circle := Circle{10}
		expected := 31.41592653589793

		assertShapePerimeter(t, circle, expected)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		expected float64
	} {
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		if tt.expected != actual {
			t.Errorf("Expected: %g, actual: %g", tt.expected, actual) 
		}
	}
}

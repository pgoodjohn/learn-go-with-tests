package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of rectangle is 2b + 2h", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		actual := Perimeter(rectangle)
		expected := 40.0

		if expected != actual {
			t.Errorf("Expected: %.2f, actual: %.2f", expected, actual)
		}
	})

	t.Run("perimeter of a circle is pi * r", func(t *testing.T) {
		circle := Circle{10}
		actual := Perimeter(circle)
		expected := 314.1592653589793

		if expected != actual {
			t.Errorf("Expected: %g, actual: %g", expected, actual)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("area of rectangle is b x h", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		actual := Area(rectangle)
		expected := 72.0

		if expected != actual {
			t.Errorf("Expected: %.2f, actual: %.2f", expected, actual)
		}
	})
}

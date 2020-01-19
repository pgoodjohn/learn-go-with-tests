package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of rectangle is 2b x 2h", func(t *testing.T) {
		actual := Perimeter(10.0, 10.0)
		expected := 40.0

		if expected != actual {
			t.Errorf("Expected: %.2f, actual: %.2f", expected, actual)
		}
	})
}

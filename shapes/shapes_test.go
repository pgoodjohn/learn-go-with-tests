package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of rectangle is 2b + 2h", func(t *testing.T) {
		actual := Perimeter(10.0, 10.0)
		expected := 40.0

		if expected != actual {
			t.Errorf("Expected: %.2f, actual: %.2f", expected, actual)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("area of rectangle is b x h", func(t *testing.T) {
		t.Skip("To be implemented")
	})
}

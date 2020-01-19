package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape Shape
		expected float64
	} {
		{Rectangle{10, 10}, 40.0},
		{Circle{10}, 31.41592653589793},
	}

	for _, tt := range perimeterTests {
		actual := tt.shape.Perimeter()
		if tt.expected != actual {
			t.Errorf("Expected: %g, actual: %g", tt.expected, actual) 
		}
	}
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

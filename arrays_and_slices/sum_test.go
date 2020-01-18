package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("it adds a fixed range of numbers", func(t *testing.T) {
		input := [5]int{1, 2, 3, 4, 5}

		actual := Sum(input)
		expected := 15

		if expected != actual {
			t.Errorf("Expected: %d, actual: %d", expected, actual)
		}
	})

	t.Run("it adds an undefined range of numbers", func(t *testing.T) {
		input := []int{1, 2, 3}
		
		actual := Sum(input)
		expected := 6

		if expected != actual {
			t.Errorf("Expected: %d, actual: %d", expected, actual)
		}
	})
}

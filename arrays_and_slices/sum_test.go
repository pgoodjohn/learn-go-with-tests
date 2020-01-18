package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("it adds an undefined range of numbers", func(t *testing.T) {
		input := []int{1, 2, 3}
		
		actual := Sum(input)
		expected := 6

		if expected != actual {
			t.Errorf("Expected: %d, actual: %d", expected, actual)
		}
	})
}

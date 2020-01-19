package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("two numbers are added", func(t *testing.T) {
		actual := Add(2, 3)
		expected := 5

		if expected != actual {
			t.Errorf("Expected: %d, actual: %d", expected, actual)
		}

	})
}

package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("letter is repeated 5 times", func(t *testing.T) {
		actual := Repeat("a", 5)
		expected := "aaaaa"

		if expected != actual {
			t.Errorf("Expected: %s, actual: %s", expected, actual)
		}
	})

	t.Run("repeatCount 0 returns empty string", func(t *testing.T) {
		actual := Repeat("a", 0)
		expected := ""

		if expected != actual {
			t.Errorf("Expected: %s, actual: %s", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

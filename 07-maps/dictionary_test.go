package dictionary

import "testing"

func TestSerach(t *testing.T) {

	t.Run("test finds this is just a test", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		actual := Search(dictionary, "test")
		expected := "this is just a test"

		if expected != actual {
			t.Errorf("Expected: %q, actual: %q", expected, actual)
		}
	})

}

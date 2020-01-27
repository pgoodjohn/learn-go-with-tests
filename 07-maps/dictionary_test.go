package dictionary

import "testing"

func assertStringsEqual(t *testing.T, expected string, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected: %q, actual: %q", expected, actual) 
	}
}

func TestSerach(t *testing.T) {

	t.Run("test finds this is just a test", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		actual := dictionary.Search("test")
		expected := "this is just a test"

		assertStringsEqual(t, expected, actual)
	})

}

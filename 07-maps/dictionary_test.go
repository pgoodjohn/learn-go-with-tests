package dictionary

import "testing"

func assertStringsEqual(t *testing.T, expected string, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected: %q, actual: %q", expected, actual) 
	}
}

func TestSerach(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("test finds this is just a test", func(t *testing.T) {

		actual, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStringsEqual(t, expected, actual)
	})

	t.Run("unknown word returns error", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		expected := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("Expected: err, got nil")
		}

		assertStringsEqual(t, expected, err.Error())
	})

}

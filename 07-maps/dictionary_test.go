package dictionary

import "testing"

func assertStringsEqual(t *testing.T, expected string, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected: %q, actual: %q", expected, actual) 
	}
}

func assertError(t *testing.T, expected, actual error) {
	t.Helper()

	if expected == nil {
		t.Errorf("Expected: error, actual: nil") 
	}

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

		assertError(t, ErrNotFound, err)
	})

	t.Run("added word is remembered", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("added", "is remembered")

		expected := "is remembered"
		actual, _ := dictionary.Search("added")

		assertStringsEqual(t, expected, actual)
	})

	t.Run("adding existing word returns error", func(t *testing.T) {
		actual := dictionary.Add("test", "this is just a test")

		assertError(t, ErrWordExists, actual)
	})

	t.Run("updating a word saves the new definition", func(t *testing.T) {
		word := "test"
		expected := "this is just an updated test"

		dictionary.Update(word, expected)

		actual, err := d.Search(word)

		assertStringsEqual(t, expected, actual)

		assertError(t, err, nil)
	})

}

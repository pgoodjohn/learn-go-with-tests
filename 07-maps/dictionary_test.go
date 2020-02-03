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

func assertNoError(t *testing.T, actual error) {
	if (actual != nil) {
		t.Errorf("Expected: nil, actual: %q", actual) 
	}
}

func TestSerach(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("test finds this is just a test", func(t *testing.T) {

		actual, err := dictionary.Search("test")
		expected := "this is just a test"

		assertStringsEqual(t, expected, actual)

		assertNoError(t, err)
	})

	t.Run("unknown word returns error", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, ErrNotFound, err)
	})

	t.Run("added word is remembered", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("added", "is remembered")

		expected := "is remembered"
		actual, err := dictionary.Search("added")

		assertStringsEqual(t, expected, actual)
		assertNoError(t, err)
	})

	t.Run("adding existing word returns error", func(t *testing.T) {
		actual := dictionary.Add("test", "this is just a test")

		assertError(t, ErrWordExists, actual)
	})

	t.Run("updating a word saves the new definition", func(t *testing.T) {
		word := "test"
		expected := "this is just an updated test"

		dictionary.Update(word, expected)

		actual, err := dictionary.Search(word)

		assertStringsEqual(t, expected, actual)
		assertNoError(t, err)
	})

	t.Run("updating a new word does not save it to the dictionary", func(t *testing.T) {
		word := "new"
		definition := "this is a new test"

		actual := dictionary.Update(word, definition)

		assertError(t, ErrWordDoesNotExist, actual)
	})

	t.Run("deleting a word removes it from the dictionary", func(t *testing.T) {
		t.Skip("To be implemented")
		word := "test"

		dictionary.Delete(word)

		_, actual := dictionary.Search(word)

		assertError(t, ErrNotFound, actual)
	})
}

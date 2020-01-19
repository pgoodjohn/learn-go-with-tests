package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected: %q, actual: %q", expected, actual)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		actual := Hello("Pietro", "")
		expected := "Hello, Pietro"

		assertCorrectMessage(t, expected, actual)

	})

	t.Run("say hello world when there is no person", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("say hello in spanish for international audience", func(t *testing.T) {
		actual := Hello("Pedro", "Spanish")
		expected := "Hola, Pedro"

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("say hello in french for international audience", func(t *testing.T) {
		actual := Hello("Pierre", "French")
		expected := "Bonjour, Pierre"

		assertCorrectMessage(t, expected, actual)
	})

}

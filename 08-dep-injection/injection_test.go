package injection

import "testing"

import "bytes"

func assertString(t *testing.T, expected, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected: %s, actual: %s", expected, actual)
	}
}

func TestGreet(t *testing.T) {

	t.Run("chris is added to the greeting", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Pietro")

		expected := "Hello, Pietro"
		actual := buffer.String()

		assertString(t, expected, actual)
	})

}

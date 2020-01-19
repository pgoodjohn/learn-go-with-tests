package banking

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposited value goes on balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		actual := wallet.Balance()
		expected := 10

		if expected != actual {
			t.Errorf("Expected: %d, actual: %d", expected, actual)
		}
	})
}

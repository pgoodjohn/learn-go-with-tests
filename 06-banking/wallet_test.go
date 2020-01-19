package banking

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposited value goes on balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		if expected != actual {
			t.Errorf("Expected: %s, actual: %s", expected, actual)
		}
	})
}
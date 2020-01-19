package banking

import "testing"

func TestWallet(t *testing.T) {

	assertWalletBalance := func(t *testing.T, w Wallet, expected Bitcoin) {
		t.Helper()
		if expected != w.Balance() {
			t.Errorf("Expected: %s, actual: %s", expected, w.Balance()) 
		}
	}

	t.Run("deposited value goes on balance", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertWalletBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdrawn value is removed from balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertWalletBalance(t, wallet, Bitcoin(10))
	})
}

package banking

import "testing"

func assertWalletBalance(t *testing.T, w Wallet, expected Bitcoin) {
	t.Helper()
	if expected != w.Balance() {
		t.Errorf("Expected: %s, actual: %s", expected, w.Balance()) 
	}
}

func assertError(t *testing.T, expected error, actual error) {
	t.Helper()
	if expected == nil {
		t.Error("Expected error but got nil") 
	}

	if expected != actual {
		t.Errorf("Expected: %q, actual: %q", expected, actual)
	}
}


func TestWallet(t *testing.T) {

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

	t.Run("withdrawing more than available balance is not allowed", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertWalletBalance(t, wallet, startingBalance)
		assertError(t, ErrInsufficientFunds, err)
	})
}

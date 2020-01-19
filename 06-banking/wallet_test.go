package banking

import "testing"

func TestWallet(t *testing.T) {

	assertWalletBalance := func(t *testing.T, w Wallet, expected Bitcoin) {
		t.Helper()
		if expected != w.Balance() {
			t.Errorf("Expected: %s, actual: %s", expected, w.Balance()) 
		}
	}

	assertError := func(t *testing.T, err error, expectedErrorMessage string) {
		t.Helper()
		if err == nil {
			t.Error("Expected error but got nil") 
		}

		if expectedErrorMessage != err.Error() {
			t.Errorf("Expected: %s, actual: %s", expectedErrorMessage, err.Error())
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

	t.Run("withdrawing more than available balance is not allowed", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertWalletBalance(t, wallet, startingBalance)
		assertError(t, err, "Cannot withdraw, insufficient funds")
	})
}

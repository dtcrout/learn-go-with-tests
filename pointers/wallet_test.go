package wallet

import "testing"

func TestWallet(t *testing.T) {

	// Create a Wallet instance
	wallet := Wallet{}

	// Deposit 10 into the wallet
	wallet.Deposit(Bitcoin(10))

	// Get the balance of the wallet
	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

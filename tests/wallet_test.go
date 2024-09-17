package tests

import (
	"testing"

	"github.com/mbilaljawwad/golang-channels-learning/internal/wallet"
)

func TestWallet(t *testing.T) {
	w := wallet.Wallet{}

	w.Deposit(10)
	balance := w.Balance()

	expected := wallet.Bitcoin(10)

	if balance != expected {
		t.Errorf("Result: %d, Expected: %d", balance, expected)
	}
}

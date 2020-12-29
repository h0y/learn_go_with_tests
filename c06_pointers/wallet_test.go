package pointers

import (
	"testing"
	"fmt"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func (t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		fmt.Println("address of balance in test:", &wallet.balance)

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("Withdraw", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Withdraw insufficient funds", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(30)

		if err != nil {
			t.Errorf("%s", err)
		}

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	
}
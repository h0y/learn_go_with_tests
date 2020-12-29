package pointers

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
	fmt.Println("address of balance in wallet:", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("the balance is less then the amount")
	}

	w.balance -= amount
	return nil
}

/*
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
*/
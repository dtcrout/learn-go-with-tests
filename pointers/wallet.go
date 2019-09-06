package wallet

import "fmt"
import "errors"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Stringer defines a "native" format for when we print out Bitcoin
type Stringer interface {
	String() string
}

// What to print out when calling Bitcoin as string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit method for wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount

}

// Withdrawl method for wallet
func (w *Wallet) Withdrawl(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil

}

// Balance method for wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

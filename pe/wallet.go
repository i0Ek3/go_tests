package pe

import (
	"errors"
	"fmt"
)

// InsufficientBalanceErr defines the error for withdraw
var InsufficientBalanceErr = errors.New("cannot withdraw casuse of insufficient balance")

// Bitcoin indicates the balance of BTC wallet
type Bitcoin int

// Wallet defines the bitcoin wallet
type Wallet struct {
	balance Bitcoin
}

type Transaction interface {
	Deposit(amount Bitcoin)
	Balance() Bitcoin
	Withdraw(amount Bitcoin) error
}

// Deposit saves the btc into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance outputs the balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// String redefines the string function with Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw withdraws the btc from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientBalanceErr
	}
	w.balance -= amount
	return nil
}

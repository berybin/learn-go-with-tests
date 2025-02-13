package main

import "fmt"

// type Wallet struct {
// 	balance int
// }

// func (w *Wallet) Deposit(amount int) {
// 	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
// 	w.balance += amount
// }

// func (w *Wallet) Balance() int {
// 	return w.balance
// }

// //////////////////////////////////////
// REFACTOR

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

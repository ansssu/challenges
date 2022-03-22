package main

import (
	"fmt"

	"github.com/andersonlribeiro/challenges/go/interfaceDemo/banks"
)

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	WithDraw(amount int) error
}

func main() {

	myAccounts := []IBankAccount{banks.NewWellsFargo(), banks.NewBitcoinAccount()}

	for _, account := range myAccounts {
		account.Deposit(200)
		if err := account.WithDraw(70); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}
		balance := account.GetBalance()
		fmt.Printf("balance = %d", balance)
	}
}

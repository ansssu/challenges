package banks

import "errors"

type BitcoinAccount struct {
	balance int
	fee     int
}

func NewBitcoinAccount() *BitcoinAccount {
	return &BitcoinAccount{balance: 0}
}

func (w *BitcoinAccount) GetBalance() int {
	return w.balance
}

func (w *BitcoinAccount) Deposit(amount int) {
	w.balance += amount
}

func (b *BitcoinAccount) WithDraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}

package main

type BicoinAccount struct {
	balance int
}

func NewBitcoinAc() *BicoinAccount {
	return &BicoinAccount{
		balance: 0,
	}
}

func (b *BicoinAccount) GetBalance() int {
	return b.balance
}

func (b *BicoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BicoinAccount) Withdraw(amount int) error {
	b.balance = b.balance - amount
	return nil
}

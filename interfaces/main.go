package main

import "fmt"

type IBankAccount interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	myAccounts := []IBankAccount{
		NewWellsFargo(),
		NewBitcoinAc(),
	}

	for _, acnt := range myAccounts {
		if acnt == NewBitcoinAc() {
			acnt.Deposit(300)
		} else {
			acnt.Deposit(500)
		}

	}

	for _, acnt := range myAccounts {
		fmt.Println(acnt.GetBalance())
	}

	// wf := NewWellsFargo()

	// wf.Deposit(200)

	// currentBalance := wf.GetBalance()

	//fmt.Println(currentBalance)

}

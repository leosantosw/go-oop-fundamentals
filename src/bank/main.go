package main

import (
	"fmt"

	"github.com/bank/accounts"
	"github.com/bank/customers"
)

func PayBoleto(account verifyAccount, boletoAmount float64) {
	account.Withdraw(boletoAmount)
}

type verifyAccount interface {
	Withdraw(amount float64) (string, float64)
}

func main() {
	fmt.Println("Hi, welcome to the bank!")
	customer := customers.Customer{Name: "Leonardo", CPF: "123.456.789.10", Occupation: "Desenvolvedor GO"}
	account1 := accounts.CheckingAccount{Customer: customer, AgencyNumber: 123, AccountNumber: 123456789}
	account2 := accounts.SavingsAccount{Customer: customer, AgencyNumber: 123, AccountNumber: 123456789}
	account1.Deposit(500)
	account2.Deposit(400)
	PayBoleto(&account1, 30)
	PayBoleto(&account2, 60)

	// fmt.Println("Account 1:", account1.GetBalance())
	fmt.Println("Account 1:", account1)
	fmt.Println("Account 2:", account2)
}

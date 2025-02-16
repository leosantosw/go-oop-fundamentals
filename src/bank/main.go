package main

import (
	"fmt"

	"github.com/bank/accounts"
	"github.com/bank/customers"
)

func main() {
	fmt.Println("Hi, welcome to the bank!")
	customer := customers.Customer{Name: "Leonardo", CPF: "123.456.789.10", Occupation: "Desenvolvedor GO"}
	account1 := accounts.CheckingAccount{Customer: customer, AgencyNumber: 123, AccountNumber: 123456789}
	account1.Deposit(500)

	fmt.Println("Account 1:", account1.GetBalance())
	fmt.Println("Account 1:", account1)
}

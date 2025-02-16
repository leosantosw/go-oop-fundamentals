package main

import (
	"fmt"

	"github.com/bank/accounts"
)

func main() {
	fmt.Println("Hi, welcome to the bank!")
	account1 := accounts.CheckingAccount{"Leonardo", 123, 123456789, 500.00}
	account2 := accounts.CheckingAccount{"Daniel", 123, 123456789, 500.00}

	fmt.Println("Account 1:", account1)
	fmt.Println("Account 2:", account2)
	fmt.Println(account1.Transfer(200, &account2))

	fmt.Println("Account 1:", account1.Balance)
	fmt.Println("Account 2:", account2.Balance)
}

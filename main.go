package main

import (
	"fmt"
)

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	fmt.Println("Hi, welcome to the bank!")
	account1 := CheckingAccount{"Leonardo", 123, 123456789, 500.00}
	account2 := new(CheckingAccount)
	account2.holder = "Leonardo"

	fmt.Println("Account:", account1)
	fmt.Println("Account:", account2)
}

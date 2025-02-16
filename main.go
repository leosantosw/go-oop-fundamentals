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

func (c *CheckingAccount) withdraw(value float64) (string, float64) {
	canWithdraw := value > 0 && value <= c.balance
	if canWithdraw {
		c.balance -= value
		return "successfully", c.balance
	} else {
		return "enough balance", c.balance
	}
}

func (c *CheckingAccount) deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "successfully", c.balance
	} else {
		return "enough deposit", c.balance
	}
}

func main() {
	fmt.Println("Hi, welcome to the bank!")
	// account1 := CheckingAccount{"Leonardo", 123, 123456789, 500.00}
	account2 := new(CheckingAccount)

	account2.holder = "Leonardo"
	account2.agencyNumber = 123
	account2.accountNumber = 123456789
	account2.balance = 500.00

	// fmt.Println("Account:", account1 == *account2)
	fmt.Println("Account:", *account2)

	fmt.Println(account2.withdraw(100.00))

	fmt.Println(account2.deposit(2000.00))
}

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

func (c *CheckingAccount) transfer(valueToTransfer float64, accountDestination *CheckingAccount) bool {
	if valueToTransfer <= c.balance && valueToTransfer > 0 {
		c.withdraw(valueToTransfer)
		accountDestination.deposit(valueToTransfer)
		return true
	}
	return false
}

func main() {
	fmt.Println("Hi, welcome to the bank!")
	account1 := CheckingAccount{"Leonardo", 123, 123456789, 500.00}
	account2 := CheckingAccount{"Daniel", 123, 123456789, 500.00}

	fmt.Println("Account 1:", account1)
	fmt.Println("Account 2:", account2)
	fmt.Println(account1.transfer(-200, &account2))

	fmt.Println("Account 1:", account1.balance)
	fmt.Println("Account 2:", account2.balance)
}

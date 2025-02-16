package accounts

import (
	"github.com/bank/customers"
)

type CheckingAccount struct {
	Customer      customers.Customer
	AgencyNumber  int
	AccountNumber int
	balance       float64
}

func (c *CheckingAccount) Withdraw(amount float64) (string, float64) {
	canWithdraw := amount > 0 && amount <= c.balance
	if canWithdraw {
		c.balance -= amount
		return "successfully", c.balance
	} else {
		return "enough balance", c.balance
	}
}

func (c *CheckingAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "successfully", c.balance
	} else {
		return "enough deposit", c.balance
	}
}

func (c *CheckingAccount) Transfer(amountToTransfer float64, accountDestination *CheckingAccount) bool {
	if amountToTransfer <= c.balance && amountToTransfer > 0 {
		c.Withdraw(amountToTransfer)
		accountDestination.Deposit(amountToTransfer)
		return true
	}
	return false
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}

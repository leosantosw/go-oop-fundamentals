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

func (c *CheckingAccount) Withdraw(value float64) (string, float64) {
	canWithdraw := value > 0 && value <= c.balance
	if canWithdraw {
		c.balance -= value
		return "successfully", c.balance
	} else {
		return "enough balance", c.balance
	}
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "successfully", c.balance
	} else {
		return "enough deposit", c.balance
	}
}

func (c *CheckingAccount) Transfer(valueToTransfer float64, accountDestination *CheckingAccount) bool {
	if valueToTransfer <= c.balance && valueToTransfer > 0 {
		c.Withdraw(valueToTransfer)
		accountDestination.Deposit(valueToTransfer)
		return true
	}
	return false
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}

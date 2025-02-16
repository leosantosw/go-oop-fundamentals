package accounts

import "github.com/bank/customers"

type SavingsAccount struct {
	Customer                               customers.Customer
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withdraw(amount float64) (string, float64) {
	canWithdraw := amount > 0 && amount <= c.balance
	if canWithdraw {
		c.balance -= amount
		return "successfully", c.balance
	} else {
		return "enough balance", c.balance
	}
}

func (c *SavingsAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "successfully", c.balance
	} else {
		return "enough deposit", c.balance
	}
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}

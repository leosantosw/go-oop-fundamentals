package accounts

import "github.com/bank/customers"

type SavingsAccount struct {
	Customer                               customers.Customer
	AgencyNumber, AccountNumber, Operation int
	Balance                                float64
}

package accounts

type CheckingAccount struct {
	Holder        string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) Withdraw(value float64) (string, float64) {
	canWithdraw := value > 0 && value <= c.Balance
	if canWithdraw {
		c.Balance -= value
		return "successfully", c.Balance
	} else {
		return "enough balance", c.Balance
	}
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "successfully", c.Balance
	} else {
		return "enough deposit", c.Balance
	}
}

func (c *CheckingAccount) Transfer(valueToTransfer float64, accountDestination *CheckingAccount) bool {
	if valueToTransfer <= c.Balance && valueToTransfer > 0 {
		c.Withdraw(valueToTransfer)
		accountDestination.Deposit(valueToTransfer)
		return true
	}
	return false
}

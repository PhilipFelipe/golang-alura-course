package entity

import (
	"fmt"
)

type CurrentAccount struct {
	Holder  string
	Agency  int
	Account int
	Balance float64
}

func (c *CurrentAccount) Withdraw(value float64) string {
	isAbleToWithdraw := value > 0 && value <= c.Balance
	if !isAbleToWithdraw {
		return "Not enough balance."
	}
	c.Balance -= value
	return "Withdraw executed successfully!"
}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {
	if value < 0 {
		return "Deposit not executed. Negative value", c.Balance
	}
	c.Balance += value
	return "Deposit executed successfully!", c.Balance
}

func (c *CurrentAccount) Transfer(value float64, destinyAccount *CurrentAccount) bool {
	if value < 0 {
		fmt.Println("Negative value")
		return false
	}
	if value > c.Balance {
		fmt.Println("The origin account does not have the requested value at it's balance")
		return false
	}
	c.Balance -= value
	destinyAccount.Deposit(value)
	return true
}

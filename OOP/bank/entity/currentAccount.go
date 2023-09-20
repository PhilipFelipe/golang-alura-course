package entity

import (
	"fmt"
)

type CurrentAccount struct {
	Holder  Holder
	Agency  int
	Account int
	balance float64 // Inicial maíuscula = público (acesso para todos que importarem)
}

func (c *CurrentAccount) Withdraw(value float64) string {
	isAbleToWithdraw := value > 0 && value <= c.balance
	if !isAbleToWithdraw {
		return "Not enough balance."
	}
	c.balance -= value
	return "Withdraw executed successfully!"
}

func (c *CurrentAccount) Deposit(value float64) (string, float64) {
	if value < 0 {
		return "Deposit not executed. Negative value", c.balance
	}
	c.balance += value
	return "Deposit executed successfully!", c.balance
}

func (c *CurrentAccount) Transfer(value float64, destinyAccount *CurrentAccount) bool {
	if value < 0 {
		fmt.Println("Negative value")
		return false
	}
	if value > c.balance {
		fmt.Println("The origin account does not have the requested value at it's balance")
		return false
	}
	c.balance -= value
	destinyAccount.Deposit(value)
	return true
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}

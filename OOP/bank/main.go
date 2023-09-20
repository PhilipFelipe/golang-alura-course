package main

import (
	"fmt"

	"github.com/PhilipFelipe/golang-alura-course/entity"
)

func payBill(account verifyAccount, value float64) {
	account.Withdraw(value)
}

type verifyAccount interface {
	Withdraw(value float64) string
}

func main() {
	gleisonAccount := entity.SavingAccount{}
	gleisonAccount.Deposit(1000)
	payBill(&gleisonAccount, 150) // o método Withdraw da struct tem um receptor com ponteiro, ex: >>s *SavingAccount<<, por isso o endereço é necessário

	carolAccount := entity.CurrentAccount{}
	carolAccount.Deposit(1050)
	payBill(&carolAccount, 10000)
	fmt.Println(gleisonAccount.GetBalance())
	fmt.Println(carolAccount.GetBalance())
}

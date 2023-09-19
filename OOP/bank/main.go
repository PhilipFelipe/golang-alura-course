package main

import (
	"fmt"

	"github.com/PhilipFelipe/golang-alura-course/entity"
)

func main() {
	felipeAccount := entity.CurrentAccount{Holder: "Felipe", Balance: 500, Agency: 1234, Account: 123321}
	anaAccount := entity.CurrentAccount{Holder: "Ana", Balance: 500, Agency: 4321, Account: 159987}

	fmt.Println(felipeAccount.Balance)
	fmt.Println(anaAccount.Balance)
	status := felipeAccount.Transfer(100, &anaAccount)
	fmt.Println(status)
	fmt.Println(felipeAccount.Balance)
	fmt.Println(anaAccount.Balance)

}

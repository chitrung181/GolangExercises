package main

import (
	"fmt"

	bank "exercises/ch9/9.1/bank"
)

func main() {
	bank.Deposit(100)
	bank.Withdraw(50)
	fmt.Println(bank.Balance())
}

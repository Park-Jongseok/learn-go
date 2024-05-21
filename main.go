package main

import (
	"fmt"
	"learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("John")
	account.Deposit(10)
	fmt.Println(account.Balance())
}

package main

import (
	"fmt"
	"learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("John")
	fmt.Println(account)
}

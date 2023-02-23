package main

import (
	"fmt"
	"github.com/JIHYE/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jihye")

	account.Deposit(10)

	fmt.Println(account)

}

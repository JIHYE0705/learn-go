package main

import (
	"fmt"
	"github.com/JIHYE0705/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jihye")

	account.Deposit(10)

	fmt.Println(account)

}

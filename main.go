package main

import (
	"fmt"
	"github.com/JIHYE0705/learngo/accounts"
	"github.com/JIHYE0705/learngo/mydict"
)

func main() {
	account := accounts.NewAccount("jihye")

	account.Deposit(10)

	fmt.Println(account)

	dictionary := mydict.Dictionary{}
}

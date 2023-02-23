package main

import (
	"fmt"
	"github.com/JIHYE/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("jihye")

	fmt.Println(account)
}

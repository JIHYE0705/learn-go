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

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found:", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}

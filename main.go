package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFaied = errors.New("request failed")

func main() {
	//account := accounts.NewAccount("jihye")
	//
	//account.Deposit(10)
	//
	//fmt.Println(account)
	//
	//dictionary := mydict.Dictionary{}

	//word := "hello"
	//definition := "Greeting"
	//
	//err := dictionary.Add(word, definition)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//hello, _ := dictionary.Search(word)
	//fmt.Println("found:", word, "definition:", hello)
	//
	//err2 := dictionary.Add(word, definition)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//baseWord := "hello"
	//dictionary.Add(baseWord, "First")
	//err := dictionary.Update(baseWord, "Second")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//word, _ := dictionary.Search(baseWord)
	//fmt.Println(word)
	//
	//err1 := dictionary.Delete(baseWord)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//fmt.Println(dictionary.Search(baseWord))

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.redit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		hitURL(url)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	response, err := http.Get(url)
	if err == nil || response.StatusCode >= 400 {
		return errRequestFaied
	}

	return nil
}

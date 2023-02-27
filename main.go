package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

// var errRequestFailed = errors.New("request failed")

func main() {
	/*
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
		baseWord := "hello"
		dictionary.Add(baseWord, "First")
		err := dictionary.Update(baseWord, "Second")
		if err != nil {
			fmt.Println(err)
		}
		word, _ := dictionary.Search(baseWord)
		fmt.Println(word)

		err1 := dictionary.Delete(baseWord)
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println(dictionary.Search(baseWord))
	*/

	results := make(map[string]string)
	c := make(chan requestResult)

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
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

	/*
		channel := make(chan string)
		people := [2]string{"jihye", "jisu"}

		for _, person := range people {
			go isCute(person, channel)
		}

		for i := 0; i < len(people); i++ {
			fmt.Println("Waiting for ", i)
			fmt.Println(<-channel)
		}
	*/
}

func hitURL(url string, c chan<- requestResult) {
	response, err := http.Get(url)
	status := "OK"

	if err != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}

}

/*
	func isCute(person string, channel chan string) {
		time.Sleep(time.Second * 10)

		channel <- person + " is Cute"
	}
*/

/*
	func sexyCount(person string) {
		for i := 0; i < 10; i++ {
			fmt.Println(person, "is sexy", i)
			time.Sleep(time.Second)
		}
	}
*/

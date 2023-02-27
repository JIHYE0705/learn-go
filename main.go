package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=java"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination")

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

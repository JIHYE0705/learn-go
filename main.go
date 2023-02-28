package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"strconv"
)

type extractedJob struct {
	id       string
	title    string
	location string
	summary  string
}

var baseURL = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=java"

func main() {
	totalPages := getPages()

	for i := 1; i <= totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page) + "&recruitSort=relation&recruitPageCount=40" +
		"&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10" +
		"&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n"
	fmt.Println("Requesting", pageURL)

	res, err := http.Get(pageURL)

	checkErr(err)
	checkCode(res)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("value")
		title := card.Find(".job_tit>a").Text()
		location := card.Find(".job_condition>span>a").Text()

		fmt.Println(id, title, location)
	})
}

func cleanString(str string) string {
	return ""
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
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

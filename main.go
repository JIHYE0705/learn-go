package main

import (
	"github.com/JIHYE0705/learngo/scrapper"
	"github.com/labstack/echo"
	"log"
	"os"
	"strings"
)

const FileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer func() {
		err := os.Remove(FileName)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)

	return c.Attachment(FileName, FileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Job struct {
	Title string
	Link  string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains
		colly.AllowedDomains("www.profesia.sk", "profesia.sk"),
		// Set mas depth level
		colly.MaxDepth(0),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".list-row", func(e *colly.HTMLElement) {
		//link := e.Attr("href")

		dataTemp := Job{}

		dataTemp.Title = e.ChildText(".title")
		dataTemp.Link = e.ChildAttr("a", "href")

		// Print link
		fmt.Printf("[%s]: %q -> %s\n", time.Now().Format("2006-01-02 15:04:05"), dataTemp.Title, dataTemp.Link)
	})

	// Start scraping on profesia.sk
	c.Visit("https://www.profesia.sk/praca/?search_anywhere=golang")
}

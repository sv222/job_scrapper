package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Job struct {
	Title string
	Link  string
}

func Parse() {
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

		// find selectors with data
		dataTemp.Title = e.ChildText(".title")
		dataTemp.Link = e.ChildAttr("a", "href")
		path := filepath.Join("data", "log.txt")
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		// filter empty data
		if strings.Contains(strings.ToLower(dataTemp.Title), "go") {
			// Save data to file
			now := time.Now()
			_, err = fmt.Fprintf(file, "[%d/%d/%d] %q -> %s\n", now.Year(), now.Month(), now.Day(), dataTemp.Title, "https://www.profesia.sk"+dataTemp.Link)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	})

	// Start scraping on profesia.sk
	c.Visit("https://www.profesia.sk/praca/?search_anywhere=golang")
}

// ParseAgain function makes iteration for a specific function
func ParseAgain(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

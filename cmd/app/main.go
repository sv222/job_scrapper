package main

import (
	"job_crawler/internal/app"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// TODO:
// 1. Make filtering final output file without duplicates.
// 2. Add new sources, not only profesia.sk.
// 3. Implement html template instead of txt (to be able to follow links).

func main() {

	path := filepath.Join("data", "log.txt")

	// create folder for data scrapping
	os.MkdirAll("data", 0755)

	// create file for data
	file, err := os.Create(path)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	// make first fetching data
	app.Parse()

	// iterate scrapping every N hours.
	go app.ParseAgain(8*time.Hour, app.Parse)

	// handle requests for index page
	http.HandleFunc("/", app.IndexHandler)

	// check health status of server
	http.HandleFunc("/health-check", app.HealthCheckHandler)

	// starting server
	log.Fatal(http.ListenAndServe(":7070", nil))
}

package main

import (
	"fmt"
	"job_crawler/internal/app"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// TODO:
// 1. Make filtering final output file without duplicates.
// 2. Add new source, not only profesia.sk.
// 3. Implement html template instead of txt (to be able to follow links).
var (
	path = filepath.Join("data", "log.txt")
)

func main() {

	file, err := os.Create(path)
	if err != nil {
		log.Println("couldn't create log file'")
	}

	defer file.Close()

	// make first fetching data
	app.Parse()

	// iterate scrapping every N hours.
	go parseAgain(8*time.Hour, app.Parse)

	// handle requests for index page
	http.HandleFunc("/", index)

	// starting server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// root handler
func index(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("couldn't read log file'")
	}

	fmt.Fprint(w, string(data))
}

// parseAgain function makes iteration for a specific function
func parseAgain(d time.Duration, f func()) {
	for _ = range time.Tick(d) {
		f()
	}
}

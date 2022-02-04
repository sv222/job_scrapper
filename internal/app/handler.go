package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// IndexHandler handle index page requests
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	path := filepath.Join("data", "log.txt")

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// pass scrapped data to ResponseWriter
	fmt.Fprint(w, string(data))
}

// HealthCheckHandler is handle checks health of server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}

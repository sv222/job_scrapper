package main

import (
	"fmt"
	"job_crawler/internal/app"
	"log"
	"net/http"
	"os"
)

type Job struct {
	Title string
	Link  string
}

func main() {

	//ticker := time.NewTicker(8 * time.Hour)
	//
	//// for every `tick` that our `ticker`
	//// emits, we print `tock`
	//for _ = range ticker.C {
	//	app.Parse()
	//}

	app.Parse()

	http.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	//tmpl, err := template.ParseFiles("file.txt")
	//if err != nil {
	//	panic(err)
	//}
	//err = tmpl.ExecuteTemplate(w, "file.txt", dataTemp)
	//if err != nil {
	//	panic(err)
	//}
	data, err := os.ReadFile("./data/log.txt")
	if err != nil {
		log.Fatal("couldn't read log file'")
	}

	fmt.Fprint(w, data)
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func righTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Right time: %s</h1>", s)
}

func main() {
	http.HandleFunc("/rightTime", righTime)

	log.Println("Exec...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

package main

import (
	"log"
	"net/http"

	"BDres.com/m/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/some-htmx-endpoint", handlers.HtmxHandler)

	log.Println("Server started on port 6969.")
	if err := http.ListenAndServe(":6969", mux); err != nil {
		log.Fatal(err)
	}
}

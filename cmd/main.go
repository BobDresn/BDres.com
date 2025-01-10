package main

import (
	"html/template"
	"log"
	"net/http"
)

type Movie struct {
	Title    string
	Director string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	Movies := map[string][]Movie{
		"Movies": {
			{Title: "Example 1", Director: "Example Director"},
			{Title: "Example 2", Director: "Example Director"},
			{Title: "Example 3", Director: "Example Director"},
		},
	}
	tmpl.Execute(w, Movies)
}

func HomePostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("HTMX Received")
	log.Print(r.Method)
	//title := r.PostFormValue("title")
	//director := r.PostFormValue("director")
}

func main() {
	mux := http.NewServeMux()

	//THIS WILL BE USED FOR STATIC FILES, NEEDS TO POINT TO CSS FOLDER
	//fs := http.FileServer(http.Dir("css"))

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/add-post/", HomePostHandler)

	log.Println("Server started on port 6969.")
	if err := http.ListenAndServe(":6969", mux); err != nil {
		log.Fatal(err)
	}
}

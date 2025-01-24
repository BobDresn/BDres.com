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
	tmpl.Execute(w, 0)
}

func HomeAddPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("HTMX Received")
	log.Println(r.Header.Get("title"))

	//title := r.PostFormValue("title")
	//director := r.PostFormValue("director")
	//fmt.Println(title, director)
}

func main() {
	mux := http.NewServeMux()

	//THIS WILL BE USED FOR STATIC FILES, NEEDS TO POINT TO CSS FOLDER
	//fs := http.FileServer(http.Dir("css"))

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/add-post/", HomeAddPostHandler)

	log.Println("Server started on port 6969.")
	if err := http.ListenAndServe(":6969", mux); err != nil {
		log.Fatal(err)
	}
}

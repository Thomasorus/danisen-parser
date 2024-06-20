package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server started")

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "style.css")
	})

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			templ404 := template.Must(template.ParseFiles("404.html"))
			templ404.Execute(w, nil)
		} else {
			templ := template.Must(template.ParseFiles("index.html"))
			templ.Execute(w, nil)
		}
	}
	http.HandleFunc("/", homeHandler)

	getRanking := func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		fmt.Println(id)
		templ := template.Must(template.ParseFiles("ranking.html"))
		templ.Execute(w, id)
	}
	http.HandleFunc("/ranking", getRanking)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

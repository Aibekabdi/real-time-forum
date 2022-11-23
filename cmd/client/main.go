package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web/src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		indexTmpl, err := template.ParseFiles("./web/index.html")
		if err != nil {
			log.Fatalf("error occured in parsing html: %v", err.Error())
		} else if err := indexTmpl.Execute(w, nil); err != nil {
			log.Fatalf("error occured in parsing html: %v", err.Error())
		}
	})
}

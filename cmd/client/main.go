package main

import (
	"forum/internal/config"
	"log"
	"net/http"
	"text/template"
)

func main() {
	conf, err := config.NewConfig("./configs/config.json")
	if err != nil {
		log.Fatalf("error occured while trying to parse config: %s", err.Error())
	}
	fs := http.FileServer(http.Dir("./web/src"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		indexTmpl, err := template.ParseFiles("./web/index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		if err := indexTmpl.Execute(w, nil); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	})
	log.Printf("Frontend server is starting at %v", "8081")
	if err := http.ListenAndServe(":"+conf.Client.Port, nil); err != nil {
		log.Fatalln(err)
	}
}

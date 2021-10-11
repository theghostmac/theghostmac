package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"time"
)

var assets map[string]*template.Template

func renderTemplate() {
	if assets == nil {
		assets = make(map[string]*template.Template)
	}
	assets["index"] = template.Must(template.ParseFiles("assets/index.html", "assets/styles.css"))
}

func main() {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("assets/"))
	router.Handle("/", fileServer)
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
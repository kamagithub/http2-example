package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)
}

var httpAddr = flag.String("http", ":8080", "Listen address")

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	log.Fatal(http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", nil))
}

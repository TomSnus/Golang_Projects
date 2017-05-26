package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("webserver/templates/",
			t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	// root
	http.Handle("/", &templateHandler{filename: "chattemplate.html"})
	// start the web server
	if err := http.ListenAndServe(":8070", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	println("Server running.....")
}

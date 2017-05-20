package main

import (
	"log"
	"net/http"
	"sync"
	"path/filepath"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}
// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r  *http.Request) {
	t.once.Do(func() {
		t.templ =  template.Must(template.ParseFiles(filepath.Join("templates",
			t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request) {
		w.Write([]byte("<html><head><title>Webserver</title></head><body>Server Running!</body></html>"))
  })
  // start the web server
  if err := http.ListenAndServe(":8070", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
	println("Server running.....")
}
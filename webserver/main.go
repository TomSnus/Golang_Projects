package main

import (
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request) {
		w.Write([]byte("<html><head><title>Webserver</title></head><body>Server Running!</body></html>"))
  })
  // start the web server
  if err := http.ListenAndServe(":8070", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
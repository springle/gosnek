package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test.html")
	})

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		registerPlayer(w, r)
	})

	if err := http.ListenAndServe(addr(), nil); err != nil {
		log.Fatal(err)
	}
}

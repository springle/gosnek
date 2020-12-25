package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test.html")
	})

	g := makeGame()
	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		registerPlayer(&g, w, r)
	})

	go g.run()
	if err := http.ListenAndServe(getAddrFromEnvironment(), nil); err != nil {
		log.Fatal(err)
	}
}

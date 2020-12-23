package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test.html")
	})

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		joinGame(w, r)
	})

	log.Println("Listening on http://localhost" + addr())
	if err := http.ListenAndServe(addr(), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func joinGame(w http.ResponseWriter, r *http.Request) {
	log.Println("New player joined!")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	client := &Client{conn: conn, send: make(chan []byte, 256)}
	go client.registerWriter()
	client.send <- []byte("hi")
}

func addr() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}

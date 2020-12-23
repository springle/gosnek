package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	timeAllowedWriteToPeer  = 1 * time.Second
	timeAllowedReadFromPeer = 1 * time.Second
	pingFrequency           = (timeAllowedReadFromPeer * 9) / 10
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

var httpToWebsocket = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// registerPlayer launches goroutines for a reader and writer per-client
func registerPlayer(w http.ResponseWriter, r *http.Request) {
	log.Println("New player joined!")
	conn, err := httpToWebsocket.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	client := &Client{conn: conn, send: make(chan []byte, 256)}
	go client.registerWriter()
	client.send <- []byte("hi")
}

// registerWriter sends game state to the client
func (c *Client) registerWriter() {
	keepAlive := time.NewTicker(pingFrequency)
	defer func() {
		keepAlive.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(timeAllowedWriteToPeer))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			} else if w, err := c.conn.NextWriter(websocket.TextMessage); err == nil {
				_, _ = w.Write(message)
				_ = w.Close()
			}
		case <-keepAlive.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(timeAllowedWriteToPeer))
			_ = c.conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}

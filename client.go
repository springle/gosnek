package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	timeAllowedWriteToPeer  = 1 * time.Second
	timeAllowedReadFromPeer = 1 * time.Second
	pingFrequency           = (timeAllowedReadFromPeer * 9) / 10
)

var httpToWebsocket = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// registerPlayer adds a new player to the game,
// then launches goroutines for a new reader and writer
func registerPlayer(g *game, w http.ResponseWriter, r *http.Request) {
	playerName := getNameFromRequest(r)
	g.joinRequests <- joinRequest{playerName}
	conn, _ := httpToWebsocket.Upgrade(w, r, nil)
	client := &Client{g: g, conn: conn, id: -1, send: make(chan GameState)}
	g.clientSet[client] = true
	go client.registerWriter()
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
		case stateOfGame, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(timeAllowedWriteToPeer))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			} else if w, err := c.conn.NextWriter(websocket.TextMessage); err == nil {
				m := ClientMessage{c.id, stateOfGame}
				b, _ := json.Marshal(m)
				_, _ = w.Write(b)
				_ = w.Close()
			}
		case <-keepAlive.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(timeAllowedWriteToPeer))
			_ = c.conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}

func getNameFromRequest(r *http.Request) string {
	keys, _ := r.URL.Query()["name"]
	return keys[0]
}

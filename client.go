package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

// registerPlayer launches goroutines for a reader and writer per-client
func registerPlayer(g *game, w http.ResponseWriter, r *http.Request) {
	playerName := "todo" + strconv.Itoa(rand.Intn(1000))
	_, nameTaken := g.nameToPlayer[playerName]
	conn, err := httpToWebsocket.Upgrade(w, r, nil)
	if nameTaken || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Println(playerName + " joined!")
		g.nameToPlayer[playerName] = &player{
			playerName,
			1,
			East,
			0,
			points{g.chooseEntranceSquare(), nil, nil, 1},
		}
	}

	client := &Client{g: g, conn: conn, name: playerName, send: make(chan GameState)}
	g.clientSet[client] = true
	go client.registerWriter()
	g.broadcast()
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
				m := ClientMessage{c.name, stateOfGame}
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

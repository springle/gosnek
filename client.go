package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

const (
	timeAllowedWriteToPeer      = 10 * time.Second
	timeAllowedReadPongFromPeer = 60 * time.Second
	pingFrequency               = (timeAllowedReadPongFromPeer * 9) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// registerWriter starts a ticker which keeps conn alive.
// When the send channel receives a new message,
// conn's writer sends it to the client.
func (c *Client) registerWriter() {
	ticker := time.NewTicker(pingFrequency)

	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(timeAllowedWriteToPeer))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if w, err := c.conn.NextWriter(websocket.TextMessage); err != nil {
				return
			} else {
				_, _ = w.Write(message)
				if err := w.Close(); err != nil {
					return
				}
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(timeAllowedWriteToPeer))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

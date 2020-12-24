package main

import (
	"github.com/gorilla/websocket"
)

const (
	East  = iota
	North = iota
	West  = iota
	South = iota
)

type game struct {
	clientSet   map[*Client]bool
	playersById map[int]*player
	foodSet     map[Point]bool
	occupiedSet map[Point]int
	boardHeight int
	boardWidth  int
}

type player struct {
	name     string
	id       int
	size     int
	heading  int
	age      int
	occupies points
}

type points struct {
	head Point
	next *points
	prev *points
	len  int
}

type Point struct {
	X int
	Y int
}

type ClientMessage struct {
	PlayerId    int
	StateOfGame GameState
}

type GameState struct {
	StateOfPlayers []PlayerState
	BoardHeight    int
	BoardWidth     int
}

type PlayerState struct {
	Name     string
	Id       int
	Age      int
	Occupies []Point
}

type Client struct {
	g    *game
	conn *websocket.Conn
	id   int
	send chan GameState
}

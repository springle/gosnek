package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	minSnakeLen        = 4
	metronomeFrequency = 1 * time.Second
)

// makeGame initializes a new game
func makeGame() game {
	return game{
		make(map[*Client]bool),
		make(map[int]*player),
		make(map[Point]bool),
		make(map[Point]int),
		5,
		10,
	}
}

// run calls step and broadcast every metronomeFrequency
func (g *game) run() {
	metronome := time.NewTicker(metronomeFrequency)
	defer metronome.Stop()
	for {
		select {
		case <-metronome.C:
			g.step()
			g.broadcast()
		}
	}
}

// chooseEntranceSquare picks a Point along the edge of the board,
// which is not already occupied by another player
func (g *game) chooseEntranceSquare() Point {
	for {
		point := Point{0, rand.Intn(g.boardHeight)}
		_, pointIsOccupied := g.occupiedSet[point]
		if !pointIsOccupied {
			g.occupiedSet[point] = 1
			return point
		}
	}
}

// addPlayer adds a new player to the game
func (g *game) addPlayer(name string) int {
	id, head := g.choosePlayerId(), g.chooseEntranceSquare()
	g.playersById[id] = &player{name, id, 1, East, 0, []Point{head}}
	log.Println(name + " joined (id #" + strconv.Itoa(id) + ")!")
	return id
}

// choosePlayerId chooses the lowest, available playerId
func (g *game) choosePlayerId() int {
	id := 1
	for {
		_, idIsTaken := g.playersById[id]
		if !idIsTaken {
			return id
		}

		id++
	}
}

// broadcast sends the GameState to every Client
func (g *game) broadcast() {
	m := g.message()
	for client := range g.clientSet {
		client.send <- m
	}
}

// message constructs an externally consumable representation of the GameState
func (g *game) message() GameState {
	m := GameState{make([]PlayerState, len(g.playersById)), g.boardHeight, g.boardWidth}

	index := 0
	for _, player := range g.playersById {
		m.StateOfPlayers[index] = PlayerState{player.name, index, player.age, player.occupies}
		index++
	}

	return m
}

// step progresses the game by one time unit
func (g *game) step() {
	// grow players
	for _, player := range g.playersById {
		player.grow(g)
	}

	// handle collisions
	g.occupiedSet = g.occupied()
	for id, player := range g.playersById {
		head := player.occupies[len(player.occupies)-1]
		count, ok := g.occupiedSet[head]
		if ok && count > 1 || g.outOfBounds(head) {
			log.Println(player.name + " has died")
			delete(g.playersById, id)
		}
	}
}

// outOfBounds determines if a Point is on the board
func (g *game) outOfBounds(p Point) bool {
	return p.X < 0 || p.Y < 0 || p.X >= g.boardWidth || p.Y >= g.boardHeight
}

// occupied establishes how many times each Point is occupied by a player
func (g *game) occupied() map[Point]int {
	occupiedSet := make(map[Point]int)
	for _, player := range g.playersById {
		for _, p := range player.occupies {
			count, ok := occupiedSet[p]
			if ok {
				occupiedSet[p] = count + 1
			} else {
				occupiedSet[p] = 1
			}
		}
	}

	return occupiedSet
}

// print clears the console and draws the current board
func (g *game) print() {
	board := make([][]int, g.boardHeight)
	for i := range board {
		board[i] = make([]int, g.boardWidth)
	}

	for _, player := range g.playersById {
		for _, p := range player.occupies {
			board[p.Y][p.X] = player.id
		}
	}

	fmt.Print("\033[H\033[2J")
	for i := range board {
		fmt.Println(board[i])
	}
}

// grow moves player towards heading and might increase the length of occupies by 1
func (p *player) grow(g *game) {
	n := p.nextSquare()
	p.occupies = append(p.occupies, n)
	_, nextSquareHasFood := g.foodSet[n]
	if len(p.occupies) > minSnakeLen {
		if nextSquareHasFood {
			delete(g.foodSet, n)
		}

		p.occupies = p.occupies[1:]
	}
}

// nextSquare determines player's next head location based on heading
func (p *player) nextSquare() Point {
	head := p.occupies[len(p.occupies)-1]
	switch p.heading {
	case East:
		return Point{head.X + 1, head.Y}
	case North:
		return Point{head.X, head.Y - 1}
	case West:
		return Point{head.X - 1, head.Y}
	case South:
		return Point{head.X, head.Y + 1}
	}

	panic("Invalid Direction: " + strconv.Itoa(p.heading))
}

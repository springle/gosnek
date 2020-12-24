package main

import (
	"math/rand"
)

// makeGame initializes a new game
func makeGame() game {
	return game{
		make(map[*Client]bool),
		make(map[string]*player),
		make(map[Point]bool),
		make(map[Point]int),
		100,
		100,
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

// broadcast sends the GameState to every Client
func (g *game) broadcast() {
	m := g.message()
	for client := range g.clientSet {
		client.send <- m
	}
}

// message constructs an externally consumable representation of the GameState
func (g *game) message() GameState {
	m := GameState{
		make([]PlayerState, len(g.nameToPlayer)),
		g.boardHeight,
		g.boardWidth,
	}

	index := 0
	for name, player := range g.nameToPlayer {
		m.StateOfPlayers[index] = PlayerState{
			name,
			index,
			player.age,
			[]Point{player.occupies.head},
		}

		index++
	}

	return m
}

// step progresses the game by one time unit
// TODO: implement
func (g *game) step() {
	for name := range g.nameToPlayer {
		player := g.nameToPlayer[name]
		switch player {
		}
	}
}

// occupiesList converts the Linked List of occupied Points to a vanilla array
// TODO: implement
func (p *player) occupiesList() []Point {
	l := make([]Point, p.occupies.len)
	return l
}

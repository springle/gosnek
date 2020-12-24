package main

import (
	"log"
	"math/rand"
	"strconv"
)

// makeGame initializes a new game
func makeGame() game {
	return game{
		make(map[*Client]bool),
		make(map[int]*player),
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

func (g *game) addPlayer(name string) int {
	id, head := g.choosePlayerId(), g.chooseEntranceSquare()
	g.playersById[id] = &player{
		name,
		id,
		1,
		East,
		0,
		points{head, nil, nil, 1},
	}

	log.Println(name + " joined (id #" + strconv.Itoa(id) + ")!")
	return id
}

// choosePlayerId chooses the lowest, available playerId
func (g *game) choosePlayerId() int {
	id := 0
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
	m := GameState{
		make([]PlayerState, len(g.playersById)),
		g.boardHeight,
		g.boardWidth,
	}

	index := 0
	for _, player := range g.playersById {
		m.StateOfPlayers[index] = PlayerState{
			player.name,
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
	for name := range g.playersById {
		player := g.playersById[name]
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

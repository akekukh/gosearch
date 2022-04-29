package game

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Game -
type Game struct {
	ch       chan string
	wg       *sync.WaitGroup
	scores   map[string]int
	winscore int
}

// NewGame - creates new game with winscore
func NewGame(winscore int) *Game {
	return &Game{
		ch:       make(chan string),
		scores:   make(map[string]int),
		winscore: winscore,
		wg:       new(sync.WaitGroup),
	}
}

// Start - starts the game
func (g *Game) Start() {
	fmt.Println("Start Game")

	rand.Seed(time.Now().Unix())

	g.wg.Add(2)

	go g.player("Andrey")
	go g.player("Dmitriy")

	g.ch <- "begin"

	g.wg.Wait()

	fmt.Println("Scores", g.scores)
}

func (g *Game) player(name string) {
	defer g.wg.Done()

	for val := range g.ch {
		if g.finished() {
			close(g.ch)
			return
		}

		var turn string
		switch val {
		case "begin", "stop", "pong":
			turn = "ping"
		case "ping":
			turn = "pong"
		}

		fmt.Println(name, turn)

		if goal() {
			g.scores[name]++
			fmt.Println(name, "score", g.scores[name])
			g.ch <- "stop"
		} else {
			g.ch <- turn
		}
	}
}

func (g *Game) finished() bool {
	for _, score := range g.scores {
		if score == g.winscore {
			return true
		}
	}

	return false
}

func goal() bool {
	return rand.Intn(5) == 1
}

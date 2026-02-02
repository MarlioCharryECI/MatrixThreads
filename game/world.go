package game

import (
	"fmt"
	"math/rand"
	"matrix_concurrency/config"
	"sync"
)

type Position struct {
	X int
	Y int
}

type World struct {
	Grid   [][]rune
	Mutex  sync.Mutex
	Neo    Position
	Phones []Position
	Agents []Position
}

func NewWorld(cfg config.Config) *World {
	grid := make([][]rune, cfg.Rows)
	for i := range grid {
		grid[i] = make([]rune, cfg.Cols)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	w := &World{Grid: grid}

	w.placeRandomWalls(cfg.NumWalls)
	w.Neo = w.randomEmptyPosition()
	w.Agents = w.randomPositions(cfg.NumAgents)
	w.Phones = w.randomPositions(cfg.NumPhones)

	w.placeEntities()
	return w
}

func (w *World) placeRandomWalls(count int) {
	for i := 0; i < count; i++ {
		p := w.randomEmptyPosition()
		w.Grid[p.X][p.Y] = 'X'
	}
}

func (w *World) randomEmptyPosition() Position {
	for {
		x := rand.Intn(len(w.Grid))
		y := rand.Intn(len(w.Grid[0]))

		if w.Grid[x][y] == '.' {
			return Position{x, y}
		}
	}
}

func (w *World) randomPositions(count int) []Position {
	positions := make([]Position, 0, count)
	for i := 0; i < count; i++ {
		positions = append(positions, w.randomEmptyPosition())
	}
	return positions
}

func (w *World) placeEntities() {
	for i := range w.Grid {
		for j := range w.Grid[i] {
			if w.Grid[i][j] != 'X' {
				w.Grid[i][j] = '.'
			}
		}
	}

	w.Grid[w.Neo.X][w.Neo.Y] = 'N'

	for _, p := range w.Phones {
		w.Grid[p.X][p.Y] = 'C'
	}

	for _, a := range w.Agents {
		w.Grid[a.X][a.Y] = 'A'
	}
}

func (w *World) Print() {
	w.Mutex.Lock()
	defer w.Mutex.Unlock()

	fmt.Println("----------------------")
	for _, row := range w.Grid {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

func IsValid(w *World, p Position) bool {
	return p.X >= 0 && p.X < len(w.Grid) &&
		p.Y >= 0 && p.Y < len(w.Grid[0]) &&
		w.Grid[p.X][p.Y] != 'X'
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const generations = 10

type Cell struct {
	row   int
	col   int
	alive bool
}

type Generation struct {
	size  int
	field [][]Cell
}

func main() {
	var n int

	fmt.Scan(&n)

	g := NewGeneration(n)
	g.init()
	g.evaluate()
}

func NewCell(row, col int) Cell {
	return Cell{row, col, false}
}

func (c *Cell) N(g *Generation) Cell {
	row := c.row - 1
	if row == -1 {
		row = g.size - 1
	}
	col := c.col
	return g.field[row][col]
}

func (c *Cell) NE(g *Generation) Cell {
	row := c.row - 1
	if row == -1 {
		row = g.size - 1
	}
	col := c.col + 1
	if col == g.size {
		col = 0
	}

	return g.field[row][col]
}

func (c *Cell) E(g *Generation) Cell {
	row := c.row
	col := c.col + 1
	if col == g.size {
		col = 0
	}
	return g.field[row][col]
}

func (c *Cell) SE(g *Generation) Cell {
	row := c.row + 1
	if row == g.size {
		row = 0
	}
	col := c.col + 1
	if col == g.size {
		col = 0
	}
	return g.field[row][col]
}

func (c *Cell) S(g *Generation) Cell {
	row := c.row + 1
	if row == g.size {
		row = 0
	}
	col := c.col
	return g.field[row][col]
}

func (c *Cell) SW(g *Generation) Cell {
	row := c.row + 1
	if row == g.size {
		row = 0
	}
	col := c.col - 1
	if col == -1 {
		col = g.size - 1
	}
	return g.field[row][col]
}

func (c *Cell) W(g *Generation) Cell {
	row := c.row
	col := c.col - 1
	if col == -1 {
		col = g.size - 1
	}
	return g.field[row][col]
}

func (c *Cell) NW(g *Generation) Cell {
	row := c.row - 1
	if row == -1 {
		row = g.size - 1
	}
	col := c.col - 1
	if col == -1 {
		col = g.size - 1
	}

	return g.field[row][col]
}

func (c *Cell) neighbors(g *Generation) []Cell {
	return []Cell{c.N(g), c.NE(g), c.E(g), c.SE(g), c.S(g), c.SW(g), c.W(g), c.NW(g)}
}

func NewGeneration(n int) *Generation {

	generation := Generation{size: n}
	generation.field = generation.generateField()

	return &generation
}

func (g *Generation) generateField() [][]Cell {
	field := make([][]Cell, g.size)
	for row := 0; row < g.size; row++ {
		field[row] = make([]Cell, g.size)
		for col := 0; col < g.size; col++ {
			field[row][col] = NewCell(row, col)
		}
	}
	return field
}

func (g *Generation) init() {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	for row, columns := range g.field {
		for col := range columns {
			g.field[row][col].alive = r.Intn(2) == 1
		}
	}
}

func (g *Generation) nextStepField() [][]Cell {
	next := g.generateField()

	for r, row := range g.field {
		for c, cell := range row {
			next[r][c].alive = g.isNextStepCellAlive(cell)
		}
	}

	return next
}

func (g *Generation) evaluate() {
	for i := 0; i < generations; i++ {
		g.field = g.nextStepField()

		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation #%d\n", i+1)
		fmt.Printf("Alive: %d\n", g.aliveCount())
		g.print()
		time.Sleep(500 * time.Millisecond)
	}
}

func (g *Generation) print() {
	for row, columns := range g.field {
		for col := range columns {
			if g.field[row][col].alive {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (g *Generation) aliveCount() int {
	count := 0

	for _, columns := range g.field {
		for _, cell := range columns {
			if cell.alive {
				count++
			}
		}
	}
	return count
}

func (g *Generation) isNextStepCellAlive(c Cell) bool {
	aliveNeighbors := 0

	for _, c := range c.neighbors(g) {
		alive := c.alive
		if alive {
			aliveNeighbors++
		}
	}

	if c.alive {
		return (aliveNeighbors == 2 || aliveNeighbors == 3)
	}
	return (aliveNeighbors == 3)
}

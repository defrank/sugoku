package sudoku

import (
	"errors"
	"fmt"
	"strings"
)

type cell struct {
	value    int
	possible [9]bool
}

func (c *cell) get() int {
	return c.value
}

func (c *cell) set(v int) {
	if v < 1 || 9 < v {
		panic(errors.New("value must be between 1 and 9"))
	}
	c.value = v
}

func (c *cell) notes() (n []int) {
	for i, p := range c.possible {
		if p {
			n = append(n, i+1)
		}
	}
	return
}

func (c *cell) note(value int) {
	c.possible[value] = true
}

func (c *cell) unnote(value int) {
	c.possible[value] = false
}

func (c *cell) String() string {
	if value := c.value; value > 0 {
		return fmt.Sprintf("%d", value)
	}
	notes := c.notes()
	var joinable []string
	for _, n := range notes {
		joinable = append(joinable, fmt.Sprintf("%d", n))
	}
	return "{" + strings.Join(joinable, ",") + "}"
}

type Segment struct {
}

type Box struct {
	X int
	Y int
	V int
}

type Grid struct {
	size  int
	cells [][]cell
}

func NewGrid(n int) (g *Grid) {
	if n < 2 {
		panic(errors.New("size must be larger than 2"))
	}

	cells := make([][]cell, n, n)
	for j := range cells {
		cells[j] = make([]cell, n, n)
		for i := 0; i < n; i++ {
			cells[j][i] = cell{}
		}
	}

	g = &Grid{size: n, cells: cells}
	return
}

func (g *Grid) Get(x, y int) (v int) {
	g.indexBoundsCheck(x, "x")
	g.indexBoundsCheck(y, "y")

	v = g.cells[y][x].value
	g.valueBoundsCheck(v)
	return
}

func (g *Grid) Set(x, y, v int) {
	g.indexBoundsCheck(x, "x")
	g.indexBoundsCheck(y, "y")
	g.valueBoundsCheck(v)

	g.cells[y][x].value = v
}

func (g *Grid) Size() int {
	return g.size
}

func (g *Grid) Iter() <-chan Box {
	ch := make(chan Box, g.size)
	go func() {
		for row := 0; row < g.size; row++ {
			for col := 0; col < g.size; col++ {
				ch <- Box{col, row, g.cells[row][col].get()}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *Grid) indexBoundsCheck(i int, s string) {
	if i < 0 || i > g.size-1 {
		msg := fmt.Sprintf("%q must be between 0 and %d", s, g.size-1)
		panic(errors.New(msg))
	}
}

func (g *Grid) valueBoundsCheck(v int) {
	if v <= 0 || v > g.size {
		msg := fmt.Sprintf("%q must be between 1 and %d", "value", g.size)
		panic(errors.New(msg))
	}
}

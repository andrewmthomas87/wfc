package wfc

import "math/rand"

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionDown
	DirectionUp
)

type Neighbor struct {
	d Direction
	c *Cell
}

func (w *Wave) neighbors(c *Cell) []Neighbor {
	var ns []Neighbor

	if c.X > 0 {
		c := w.Grid[c.Y][c.X-1]
		ns = append(ns, Neighbor{
			d: DirectionLeft,
			c: c,
		})
	}
	if c.X < w.W-1 {
		c := w.Grid[c.Y][c.X+1]
		ns = append(ns, Neighbor{
			d: DirectionRight,
			c: c,
		})
	}
	if c.Y > 0 {
		c := w.Grid[c.Y-1][c.X]
		ns = append(ns, Neighbor{
			d: DirectionDown,
			c: c,
		})
	}
	if c.Y < w.H-1 {
		c := w.Grid[c.Y+1][c.X]
		ns = append(ns, Neighbor{
			d: DirectionUp,
			c: c,
		})
	}

	return ns
}

func (w *Wave) neighborCells(c *Cell) []*Cell {
	var ns []*Cell

	if c.X > 0 {
		c := w.Grid[c.Y][c.X-1]
		ns = append(ns, c)
	}
	if c.X < w.W-1 {
		c := w.Grid[c.Y][c.X+1]
		ns = append(ns, c)
	}
	if c.Y > 0 {
		c := w.Grid[c.Y-1][c.X]
		ns = append(ns, c)
	}
	if c.Y < w.H-1 {
		c := w.Grid[c.Y+1][c.X]
		ns = append(ns, c)
	}

	return ns
}

func initGrid(rand *rand.Rand, statesCount, w, h int) [][]*Cell {
	g := make([][]*Cell, h)
	for y := 0; y < h; y++ {
		g[y] = make([]*Cell, w)

		for x := 0; x < w; x++ {
			states := make([]State, statesCount)
			for s := 0; s < statesCount; s++ {
				states[s] = s
			}

			g[y][x] = &Cell{
				rand:   rand,
				X:      x,
				Y:      y,
				States: states,
			}
		}
	}

	return g
}

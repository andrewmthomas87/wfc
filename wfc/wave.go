package wfc

import (
	"math/rand"
)

type AreStatesCompatibleFn func(s1, s2 State, d Direction) bool

type Wave struct {
	rand *rand.Rand

	weights []int

	W, H int
	Grid [][]*Cell

	compatFn AreStatesCompatibleFn
}

func NewWave(rand *rand.Rand, statesCount int, weights []int, w, h int, compatFn AreStatesCompatibleFn) *Wave {
	return &Wave{
		rand:     rand,
		weights:  weights,
		W:        w,
		H:        h,
		Grid:     initGrid(rand, statesCount, w, h),
		compatFn: compatFn,
	}
}

func (w *Wave) IsContradictory() bool {
	for _, col := range w.Grid {
		for _, c := range col {
			if c.IsContradictory() {
				return true
			}
		}
	}

	return false
}

func (w *Wave) IsCollapsed() bool {
	for _, col := range w.Grid {
		for _, c := range col {
			if !c.IsCollapsed() {
				return false
			}
		}
	}

	return true
}

func (w *Wave) Iterate() {
	cs := w.minEntropyCells()
	c := cs[w.rand.Intn(len(cs))]

	c.observe(w.weights)

	w.propagate(w.neighborCells(c))
}

func (w *Wave) propagate(queue []*Cell) {
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		isDirty := c.update(w.neighbors(c), w.compatFn)
		if isDirty {
			queue = append(queue, w.neighborCells(c)...)
		}
	}
}

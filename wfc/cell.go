package wfc

import "math/rand"

type State = int

type Cell struct {
	rand *rand.Rand

	X, Y   int
	States []State
}

func (c *Cell) Entropy() int {
	return len(c.States)
}

func (c *Cell) IsContradictory() bool {
	return len(c.States) == 0
}

func (c *Cell) IsCollapsed() bool {
	return len(c.States) == 1
}

func (c *Cell) observe(weights []int) {
	var pool []int
	for _, s := range c.States {
		for i := 0; i < weights[s]; i++ {
			pool = append(pool, s)
		}
	}

	i := c.rand.Intn(len(pool))
	s := pool[i]
	c.States = []State{s}
}

func (c *Cell) update(ns []Neighbor, compatFn AreStatesCompatibleFn) bool {
	if c.IsContradictory() || c.IsCollapsed() {
		return false
	}

	isDirty := false
	for _, n := range ns {
		var states []State
		for _, s1 := range c.States {
			isCompatible := false
			for _, s2 := range n.c.States {
				if compatFn(s1, s2, n.d) {
					isCompatible = true
				}
			}

			if isCompatible {
				states = append(states, s1)
			} else {
				isDirty = true
			}
		}

		c.States = states
	}

	return isDirty
}

package wfc

func (w *Wave) minEntropyCells() []*Cell {
	min := -1
	for _, col := range w.Grid {
		for _, c := range col {
			if c.IsCollapsed() {
				continue
			}

			e := c.Entropy()
			if min == -1 || e < min {
				min = e
			}
		}
	}

	var cs []*Cell
	for _, col := range w.Grid {
		for _, c := range col {
			if c.Entropy() == min {
				cs = append(cs, c)
			}
		}
	}

	return cs
}

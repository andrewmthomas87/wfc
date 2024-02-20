package main

import "github.com/andrewmthomas87/wfc/wfc"

type Arrow struct {
	In  wfc.Direction
	Out wfc.Direction
}

var arrows = []Arrow{
	{
		In:  wfc.DirectionLeft,
		Out: wfc.DirectionLeft,
	},
	{
		In:  wfc.DirectionRight,
		Out: wfc.DirectionRight,
	},
	{
		In:  wfc.DirectionDown,
		Out: wfc.DirectionDown,
	},
	{
		In:  wfc.DirectionUp,
		Out: wfc.DirectionUp,
	},
	{
		In:  wfc.DirectionLeft,
		Out: wfc.DirectionDown,
	},
	{
		In:  wfc.DirectionDown,
		Out: wfc.DirectionLeft,
	},
	{
		In:  wfc.DirectionRight,
		Out: wfc.DirectionDown,
	},
	{
		In:  wfc.DirectionDown,
		Out: wfc.DirectionRight,
	},
	{
		In:  wfc.DirectionLeft,
		Out: wfc.DirectionUp,
	},
	{
		In:  wfc.DirectionUp,
		Out: wfc.DirectionLeft,
	},
	{
		In:  wfc.DirectionRight,
		Out: wfc.DirectionUp,
	},
	{
		In:  wfc.DirectionUp,
		Out: wfc.DirectionRight,
	},
	{
		In:  -1,
		Out: -1,
	},
}

var weights = []int{6, 6, 6, 6, 1, 1, 1, 1, 1, 1, 1, 1, 24}

func buildCompatMatrix(arrows []Arrow) [][][4]bool {
	m := make([][][4]bool, len(arrows))
	for i, a1 := range arrows {
		m[i] = make([][4]bool, len(arrows))

		for j, a2 := range arrows {
			for d := wfc.Direction(0); d < 4; d++ {
				var isCompat = false

				if d == a1.Out && d == a2.In {
					isCompat = true
				} else if d != a1.Out && d != a2.In {
					isCompat = true
				}

				m[i][j][d] = isCompat
			}
		}
	}

	for i := range arrows {
		for j := range arrows {
			for d := wfc.Direction(0); d < 4; d++ {
				id := inverseDirection(d)
				m[i][j][d] = m[i][j][d] && m[j][i][id]
			}
		}
	}

	return m
}

func inverseDirection(d wfc.Direction) wfc.Direction {
	switch d {
	case wfc.DirectionLeft:
		return wfc.DirectionRight
	case wfc.DirectionRight:
		return wfc.DirectionLeft
	case wfc.DirectionDown:
		return wfc.DirectionUp
	case wfc.DirectionUp:
		return wfc.DirectionDown
	}

	panic("invalid direction")
}

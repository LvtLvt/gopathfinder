package main

import (
	"math"
	"strings"
)

const (
	SpotPlain = iota
	SpotBlock
	SpotFrom
	SpotTo
	SpotPath
)

var FeatureRunes = map[int]rune{
	SpotPlain: '.',
	SpotBlock: 'x',
	SpotFrom:  'f',
	SpotTo:    't',
	SpotPath:  'o',
}

var RuneFeatures = map[rune]int{
	'.': SpotPlain,
	'x': SpotBlock,
	'f': SpotFrom,
	't': SpotTo,
	'o': SpotPath,
}

type Spot struct {
	Feature int
	X, Y int
	G    Grid
}

func (s Spot) Neighbors() []Noder {
	var ret []Noder
	for _, offset := range [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	} {
		v, ok := s.G[s.X + offset[0]][s.Y + offset[1]]
		if ok && v.Feature != SpotBlock {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s Spot) NeighborCost(to Noder) float64 {
	return 1
}

func (s Spot) NeighborHeuristicCost(to Noder) float64 {
	toS := to.(*Spot)

	absX := math.Abs(float64(s.X - toS.X))
	absY := math.Abs(float64(s.Y - toS.Y))

	return absX + absY
}

type Grid map[int]map[int]*Spot

func (w Grid) FirstOfKind(feature int) *Spot {
	for _, row := range w {
		for _, t := range row {
			if t.Feature == feature {
				return t
			}
		}
	}
	return nil
}

func (g Grid) From() *Spot {
	return g.FirstOfKind(SpotFrom)
}

func (g Grid) To() *Spot {
	return g.FirstOfKind(SpotTo)
}

func (g Grid) setSpot(s *Spot, x, y int) {
	if g[x] == nil {
		g[x] = map[int]*Spot{}
	}
	g[x][y] = s
	s.X = x
	s.Y = y
}

func ParseGrid(grid string) Grid {
	g := Grid{}

	for y, row := range strings.Split(strings.TrimSpace(grid), "\n") {
		for x, raw := range row {
			feature, ok := RuneFeatures[raw]

			if !ok {
				feature = SpotBlock
			}
			g.setSpot(&Spot{
				G: g,
				Feature: feature,
			}, x, y)
		}
	}

	return g
}

func (g Grid) RenderPath() string {
	width := len(g)
	height := len(g[0])

	var rows = make([]string, height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rows[y] += string(FeatureRunes[g[x][y].Feature])
		}
	}

	return strings.Join(rows, "\n")
}
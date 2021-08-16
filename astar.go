package main

import (
	"container/heap"
	"math"
)

type Noder interface {
	Neighbors() []Noder
	NeighborCost(to Noder) float64
	NeighborHeuristicCost(to Noder) float64
}

type node struct {
	noder  Noder
	index  int
	cost float64
	rank   float64
	parent *node
	closed bool
}

type nodeMap map[Noder]*node

func (nm nodeMap) get(noder Noder) *node {
	n, ok := nm[noder]
	if !ok {
		n = &node{
			noder: noder,
			cost: math.MaxFloat64,
		}
		nm[noder] = n
	}
	return n
}

func Find(from, to Noder, debug bool) (path []node, found bool, hist []node) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	var hq []node
	if debug {
		hq = []node{}
	}
	heap.Init(nq)
	toNode := nm.get(to)
	fromNode := nm.get(from)
	fromNode.cost = 0
	fromNode.closed = true
	heap.Push(nq, fromNode)

	for {
		if nq.Len() == 0 {
			return nil, false, hq
		}
		current := heap.Pop(nq).(*node)

		if current == toNode {
			// 현재까지 패스, 거리 반환.
			var paths []node
			curr := current.parent
			for {
				if curr == nil {
					return paths, true, hq
				}
				paths = append(paths, *curr)
				curr = curr.parent
			}
		}

		if debug {
			hq = append(hq, *current)
		}

		for _, neighbor := range current.noder.Neighbors() {
			cost := current.cost + current.noder.NeighborCost(neighbor)
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				neighborNode.cost = cost
				neighborNode.parent = current
				neighborNode.closed = false
			}
			if !neighborNode.closed {
				neighborNode.closed = true
				neighborNode.rank = neighbor.NeighborHeuristicCost(to)
				heap.Push(nq, neighborNode)
			}
		}

	}
}
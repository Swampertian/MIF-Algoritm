package main

import (
	"math"
)

func (g *Graph) closestStaleNeighbor(dataID int) (int, float64) {
	minCost := math.Inf(1)
	best := -1
	for _, e := range g.Edges[dataID] {
		dest := g.Nodes[e.To]
		if len(dest.StalePackets) > 0 && dest.Energy > 0 && e.Cost < minCost {
			minCost = e.Cost
			best = e.To
		}
	}
	return best, minCost
}

func (g *Graph) freshness(t int) int {
	sum := 0
	for _, n := range g.Nodes {
		for _, p := range n.StalePackets {
			sum += t - p.Timestamp
		}
		for _, p := range n.NewPackets {
			sum += t - p.Timestamp
		}
	}
	return sum
}

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

func bottleneckCapacity(dataNode, staleNode *Node, cost float64) int {

	if cost <= 0 {
		return 0
	}
	ePerPacket := cost
	maxSend := int(math.Floor(dataNode.Energy / ePerPacket))
	maxRecv := int(math.Floor(staleNode.Energy / ePerPacket))
	if maxSend < maxRecv {
		return maxSend
	}
	return maxRecv
}

func updateEnergy(sender, receiver *Node, cost float64, packets int) {
	totalCost := cost * float64(packets)
	sender.Energy -= totalCost
	receiver.Energy -= totalCost
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

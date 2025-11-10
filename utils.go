package main

import (
	"math"
)

func (g *Graph) findClosestStaleNode(dataID int) (int, float64) {
	minCost := math.Inf(1)
	closest := -1
	for _, e := range g.Edges[dataID] {
		target := g.Nodes[e.To]
		if target.StalePackets > 0 && e.Cost < minCost {
			minCost = e.Cost
			closest = e.To
		}
	}
	return closest, minCost
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

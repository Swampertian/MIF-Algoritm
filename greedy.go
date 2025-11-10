package main

import "fmt"

func (g *Graph) GreedyMIF() {
	for _, node := range g.Nodes {
		if node.NewPackets <= 0 {
			continue
		}

		for node.NewPackets > 0 {
			staleID, cost := g.findClosestStaleNode(node.ID)
			if staleID == -1 {
				break
			}

			staleNode := g.Nodes[staleID]
			bottle := bottleneckCapacity(node, staleNode, cost)
			if bottle == 0 {
				break
			}

			q := minInt(node.NewPackets, staleNode.StalePackets, bottle)

			// atualiza energias e métricas
			updateEnergy(node, staleNode, cost, q)
			g.Metrics.TotalEnergyConsumed += cost * float64(q)
			g.Metrics.TotalPacketsOffloaded += q

			node.NewPackets -= q
			staleNode.StalePackets -= q

			// remove nós com energia <= 0
			if node.Energy <= 0 {
				node.Energy = 0
			}
			if staleNode.Energy <= 0 {
				staleNode.Energy = 0
			}
		}
	}

	depleted := 0
	for _, n := range g.Nodes {
		if n.Energy <= 0 {
			depleted++
		}
	}
	g.Metrics.EnergyDepletedNodes = depleted
}

// Utils
func minInt(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func (g *Graph) PrintMetrics() {
	fmt.Println("\n===== MÉTRICAS MIF (Greedy) =====")
	fmt.Printf("Energia total consumida: %.2f\n", g.Metrics.TotalEnergyConsumed)
	fmt.Printf("Total de pacotes offloadados (Transferir quando encher): %d\n", g.Metrics.TotalPacketsOffloaded)
	fmt.Printf("Nós com energia esgotada: %d\n", g.Metrics.EnergyDepletedNodes)
	fmt.Println("=================================")
}

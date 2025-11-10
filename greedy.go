package main

import (
	"fmt"
	"math"
)

func (g *Graph) GreedyMIFSimple(t int) {
	g.Metrics.FreshnessBefore = g.freshness(t)

	for _, dn := range g.Nodes {
		for len(dn.NewPackets) > 0 && dn.Energy > 0 {
			staleID, cost := g.closestStaleNeighbor(dn.ID)
			if staleID == -1 {
				break
			}
			sn := g.Nodes[staleID]
			ePerPkt := cost * 2 // tx+rx

			maxSend := int(math.Floor(dn.Energy / ePerPkt))
			maxRecv := int(math.Floor(sn.Energy / ePerPkt))
			q := min3(len(dn.NewPackets), len(sn.StalePackets), maxSend, maxRecv)
			if q <= 0 {
				break
			}

			dn.NewPackets = dn.NewPackets[q:]
			sn.StalePackets = sn.StalePackets[q:]

			dn.Energy -= ePerPkt * float64(q) / 2
			sn.Energy -= ePerPkt * float64(q) / 2

			g.Metrics.TotalEnergyConsumed += ePerPkt * float64(q)
			g.Metrics.TotalPacketsOffloaded += q
		}
	}

	g.Metrics.FreshnessAfter = g.freshness(t)
	g.Metrics.FreshnessGain = g.Metrics.FreshnessBefore - g.Metrics.FreshnessAfter
}

// Utils
func min3(a, b, c, d int) int {
	m := a
	for _, v := range []int{b, c, d} {
		if v < m {
			m = v
		}
	}
	return m
}

func (g *Graph) PrintMetrics() {
	fmt.Println("\n===== MÉTRICAS MIF (Greedy) =====")
	fmt.Printf("Energia total consumida: %.2f\n", g.Metrics.TotalEnergyConsumed)
	fmt.Printf("Total de pacotes offloadados (Transferir quando encher): %d\n", g.Metrics.TotalPacketsOffloaded)
	fmt.Printf("Nós com energia esgotada: %d\n", g.Metrics.EnergyDepletedNodes)

	fmt.Println("\n=========Intervalo de tempo dos dados==============")
	fmt.Printf("Idade acumulada da rede: %d\n", g.Metrics.FreshnessBefore)
	fmt.Printf("Idade atual da rede (Pós-algoritmo): %d\n", g.Metrics.FreshnessAfter)
	fmt.Printf("Queda da média de idade da rede : %d\n ", g.Metrics.FreshnessGain)
	fmt.Println("=================================")
}

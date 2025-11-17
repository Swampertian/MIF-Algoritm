package main

import (
	"fmt"
	"math"
)

func (g *Graph) GreedyMIFSimple(t int) {
	g.Metrics.FreshnessBefore = g.freshness(t)

	// Reiniciar contagem de nós esgotados e marcar nós já esgotados
	g.Metrics.EnergyDepletedNodes = 0
	for _, n := range g.Nodes {
		if n.Energy <= 0 {
			n.Depleted = true
			g.Metrics.EnergyDepletedNodes++
		} else {
			n.Depleted = false
		}
	}

	for _, dn := range g.Nodes {
		for len(dn.NewPackets) > 0 && dn.Energy > 0 {
			staleID, cost := g.closestStaleNeighbor(dn.ID)
			if staleID == -1 {
				break
			}
			sn := g.Nodes[staleID]
			ePerNode := cost
			ePerPkt := ePerNode * 2

			maxSend := int(math.Floor(dn.Energy / ePerNode))
			maxRecv := int(math.Floor(sn.Energy / ePerNode))
			q := min3(len(dn.NewPackets), len(sn.StalePackets), maxSend, maxRecv)
			if q <= 0 {
				if maxSend == 0 {
					break
				}
				continue
			}

			dn.NewPackets = dn.NewPackets[q:]
			sn.StalePackets = sn.StalePackets[q:]
			// consumir energia por nó
			consumedPerNode := ePerNode * float64(q)
			dn.Energy -= consumedPerNode
			sn.Energy -= consumedPerNode

			if dn.Energy <= 0 {
				if dn.Energy < 0 {
					dn.Energy = 0
				}
				if !dn.Depleted {
					dn.Depleted = true
					g.Metrics.EnergyDepletedNodes++
				}
			}
			if sn.Energy <= 0 {
				if sn.Energy < 0 {
					sn.Energy = 0
				}
				if !sn.Depleted {
					sn.Depleted = true
					g.Metrics.EnergyDepletedNodes++
				}
			}

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

// countEnergyDepletedNodes removed; contagem agora é incremental durante as operações

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

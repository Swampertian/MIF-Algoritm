package main

import "fmt"

func (g *Graph) GreedyMIF() {
	for _, node := range g.Nodes {
		// só processa nós com pacotes novos
		if node.NewPackets <= 0 {
			continue
		}

		for node.NewPackets > 0 {
			staleID, cost := g.findClosestStaleNode(node.ID)
			if staleID == -1 {
				fmt.Printf("Node %d não encontrou stale nodes disponíveis\n", node.ID)
				break
			}

			staleNode := g.Nodes[staleID]

			// calcula a capacidade do caminho
			bottle := bottleneckCapacity(node, staleNode, cost)
			if bottle == 0 {
				break
			}

			// define quantos pacotes transferir
			q := minInt(node.NewPackets, staleNode.StalePackets, bottle)

			// transfere pacotes e atualiza energias
			updateEnergy(node, staleNode, cost, q)
			node.NewPackets -= q
			staleNode.StalePackets -= q

			fmt.Printf("Offload %d pacotes de %d → %d (custo=%.2f)\n",
				q, node.ID, staleNode.ID, cost)

			// remove nós sem energia (energia <= 0)
			if node.Energy <= 0 {
				node.Energy = 0
				break
			}
			if staleNode.Energy <= 0 {
				staleNode.Energy = 0
			}
		}
	}
}

// utilitário para mínimo entre três inteiros
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

package main

import (
	"fmt"
)

func main() {
	g := &Graph{
		Nodes: map[int]*Node{
			1: {ID: 1, Energy: 100, NewPackets: 3},
			2: {ID: 2, Energy: 80, StalePackets: 2},
			3: {ID: 3, Energy: 70, StalePackets: 1},
		},
		Edges: map[int][]Edge{
			1: {{To: 2, Cost: 5}, {To: 3, Cost: 8}},
			2: {{To: 1, Cost: 5}, {To: 3, Cost: 6}},
			3: {{To: 1, Cost: 8}, {To: 2, Cost: 6}},
		},
	}

	fmt.Println(g)
}

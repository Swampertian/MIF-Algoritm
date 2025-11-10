package main

func main() {
	g := &Graph{
		Nodes: map[int]*Node{
			1: {
				ID:     1,
				Energy: 100,
				NewPackets: []Packet{
					{Timestamp: 10}, {Timestamp: 10}, {Timestamp: 10},
				},
			},
			2: {
				ID:     2,
				Energy: 80,
				StalePackets: []Packet{
					{Timestamp: 1}, {Timestamp: 2},
				},
			},
			3: {
				ID:     3,
				Energy: 80,
				StalePackets: []Packet{
					{Timestamp: 3}, {Timestamp: 4},
				},
			},
		},
		Edges: map[int][]Edge{
			1: {{To: 2, Cost: 5}, {To: 3, Cost: 8}},
			2: {{To: 1, Cost: 5}},
			3: {{To: 1, Cost: 8}},
		},
	}

	g.GreedyMIFSimple(12)
	g.PrintMetrics()
}

package main

type Node struct {
	ID           int
	Energy       float64
	NewPackets   int
	StalePackets int
}

type Edge struct {
	To   int
	Cost float64
}

type Graph struct {
	Nodes   map[int]*Node
	Edges   map[int][]Edge
	Metrics Metrics
}

type Metrics struct {
	TotalEnergyConsumed   float64
	TotalPacketsOffloaded int
	EnergyDepletedNodes   int
}

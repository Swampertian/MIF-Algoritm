package main

type Packet struct {
	Timestamp int
}
type Node struct {
	ID           int
	Energy       float64
	NewPackets   []Packet
	StalePackets []Packet
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
	FreshnessBefore       int
	FreshnessAfter        int
	FreshnessGain         int
}

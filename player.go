package main

type Mover interface {
	// GetMove(ir InputReader, b Board) int
}

type Player struct {
	Id     string
	Name   string
	Marker string
}

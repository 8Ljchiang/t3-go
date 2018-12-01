package main

type Mover interface {
	// GetMove(ir InputReader, b Board) int
}

type Player struct {
	Id     string
	Name   string
	Marker string
}

func (p Player) GetMove(b Board) (move Move) {
	// position = ir.GetInput(p.Name + " - Please enter a position (1-" + convertIntToString(b.Width*b.Height) + "): ")
	// return position
	return Move{p.Marker, 1, p.Id}
}

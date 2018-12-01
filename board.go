package board

type Move struct {
	Mark     string
	Position int
	PlayerId string
}

type Board struct {
	Size  int
	Moves []Move
}

func (b *Board) AddMove(move Move) {
	b.Moves = append(b.Moves, move)
}

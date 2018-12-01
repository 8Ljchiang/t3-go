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
	if len(b.Moves) < b.Size*b.Size {
		b.Moves = append(b.Moves, move)
	}
}

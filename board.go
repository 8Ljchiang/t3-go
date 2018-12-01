package board

const (
	ErrBoardFull           = BoardErr("board is full")
	ErrBoardPositionFilled = BoardErr("board position is filled")
)

type BoardErr string

func (e BoardErr) Error() string {
	return string(e)
}

type Move struct {
	Mark     string
	Position int
	PlayerId string
}

type Board struct {
	Size  int
	Moves []Move
}

func (b *Board) AddMove(move Move) (Move, error) {
	if len(b.Moves) < b.Size*b.Size {
		if isValidPosition(b, move.Position) {
			b.Moves = append(b.Moves, move)
			return move, nil
		} else {
			return move, ErrBoardPositionFilled
		}
	}
	return move, ErrBoardFull
}

func isValidPosition(b *Board, position int) bool {
	for _, move := range b.Moves {
		if move.Position == position {
			return false
		}
	}
	return true
}

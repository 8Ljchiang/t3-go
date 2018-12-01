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
		b.Moves = append(b.Moves, move)
		return move, nil
	}
	return move, ErrBoardFull
}

package main

import "strconv"

const (
	ErrBoardFull            = BoardErr("board is full")
	ErrBoardPositionInvalid = BoardErr("board position is invalid")
	ADJMENT_VAL             = 1
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
	Id    string
	Size  int
	Moves []Move
}

func (b *Board) AddMove(move Move) (Move, error) {
	if len(b.Moves) < b.Size*b.Size {
		if isValidPosition(b, move.Position) {
			b.Moves = append(b.Moves, move)
			return move, nil
		} else {
			return move, ErrBoardPositionInvalid
		}
	}
	return move, ErrBoardFull
}

// Not tested
func (b Board) GetPositionMark(position int) string {
	for _, move := range b.Moves {
		if move.Position == position {
			return move.Mark
		}
	}
	return strconv.Itoa(position)
}

func getPositionMark(b Board, position int) string {
	for _, move := range b.Moves {
		if move.Position == position {
			return move.Mark
		}
	}
	return strconv.Itoa(position)
}

func (b Board) GetTakenPositions() (positions []int) {
	for _, move := range b.Moves {
		positions = append(positions, move.Position)
	}
	return positions
}

func (b Board) GetEmptyPositions() (positions []int) {
	takenPositions := b.GetTakenPositions()
	for i := 1; i <= b.Size*b.Size; i++ {
		if !containsPosition(takenPositions, i) {
			positions = append(positions, i)
		}
	}
	return positions
}

func adjustIndexToPosition(index int) int {
	return index + ADJMENT_VAL
}

func isValidPosition(b *Board, position int) bool {
	if position < 0 || position > (b.Size*b.Size) {
		return false
	}

	for _, move := range b.Moves {
		if move.Position == position {
			return false
		}
	}
	return true
}

func containsPosition(positions []int, position int) bool {
	for _, p := range positions {
		if p == position {
			return true
		}
	}
	return false
}

func getTakenPositions(b Board) (positions []int) {
	for _, move := range b.Moves {
		positions = append(positions, move.Position)
	}
	return positions
}

func getEmptyPositions(b Board) (positions []int) {
	takenPositions := getTakenPositions(b)
	for i := 1; i <= b.Size*b.Size; i++ {
		if !containsPosition(takenPositions, i) {
			positions = append(positions, i)
		}
	}
	return positions
}

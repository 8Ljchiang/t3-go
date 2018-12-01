package main

import (
	"reflect"
	"testing"
)

func TestAddMove(t *testing.T) {
	t.Run("add move to empty board", func(t *testing.T) {
		board := createNewBoard()
		move := Move{"X", 1, "playerId-1"}
		_, err := board.AddMove(move)

		got := board.Moves
		want := []Move{move}

		assertError(t, err, nil)
		assertMoves(t, got, want)
	})

	t.Run("add move to a full board", func(t *testing.T) {
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 2, "playerId-2"}
		move3 := Move{"X", 3, "playerId-1"}
		move4 := Move{"O", 4, "playerId-2"}
		move5 := Move{"X", 5, "playerId-1"}
		move6 := Move{"O", 6, "playerId-2"}
		move7 := Move{"X", 7, "playerId-1"}
		move8 := Move{"O", 8, "playerId-2"}
		move9 := Move{"X", 9, "playerId-1"}
		moves := []Move{move1, move2, move3, move4, move5, move6, move7, move8, move9}
		board := createFullBoard(moves)
		_, err := board.AddMove(Move{"O", 10, "playerId-2"})

		got := board.Moves
		want := moves

		assertError(t, err, ErrBoardFull)
		assertMoves(t, got, want)
	})

	t.Run("add move to a position that is already taken", func(t *testing.T) {
		board := createNewBoard()
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 1, "playerId-2"}
		_, err1 := board.AddMove(move1)
		_, err2 := board.AddMove(move2)

		got := board.Moves
		want := []Move{move1}

		assertError(t, err1, nil)
		assertError(t, err2, ErrBoardPositionInvalid)
		assertMoves(t, got, want)
	})

	t.Run("adding a move to a postion greater than the board size", func(t *testing.T) {
		board := createNewBoard()
		move1 := Move{"X", 10, "playerId-1"}
		_, err := board.AddMove(move1)

		got := board.Moves
		want := []Move{}

		assertError(t, err, ErrBoardPositionInvalid)
		assertMoves(t, got, want)
	})

	t.Run("adding a move to a position less than 0", func(t *testing.T) {
		board := createNewBoard()
		move1 := Move{"X", -1, "playerId-1"}
		_, err := board.AddMove(move1)

		got := board.Moves
		want := []Move{}

		assertError(t, err, ErrBoardPositionInvalid)
		assertMoves(t, got, want)
	})

	t.Run("get taken positions from board", func(t *testing.T) {
		board := createNewBoard()
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 2, "playerId-2"}
		board.AddMove(move1)
		board.AddMove(move2)
		takenPositions := board.GetTakenPositions()

		got := takenPositions
		want := []int{1, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get empty positions from empty board", func(t *testing.T) {
		board := createNewBoard()
		emptyPostions := board.GetEmptyPositions()

		got := emptyPostions
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}

func createNewBoard() Board {
	return Board{3, []Move{}}
}

func createFullBoard(moves []Move) Board {
	board := Board{3, moves}
	return board
}

func assertMoves(t *testing.T, got, want []Move) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

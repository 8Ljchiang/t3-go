package main

import (
	"reflect"
	"testing"
)

func TestAddMove(t *testing.T) {
	t.Run("add move to empty board", func(t *testing.T) {
		board := createNewBoard([]Move{})
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
		board := createFullBoard()
		_, err := board.AddMove(Move{"O", 10, "playerId-2"})

		got := board.Moves
		want := moves

		assertError(t, err, ErrBoardFull)
		assertMoves(t, got, want)
	})

	t.Run("add move to a position that is already taken", func(t *testing.T) {
		board := createNewBoard([]Move{})
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
		board := createNewBoard([]Move{})
		move1 := Move{"X", 10, "playerId-1"}
		_, err := board.AddMove(move1)

		got := board.Moves
		want := []Move{}

		assertError(t, err, ErrBoardPositionInvalid)
		assertMoves(t, got, want)
	})

	t.Run("adding a move to a position less than 0", func(t *testing.T) {
		board := createNewBoard([]Move{})
		move1 := Move{"X", -1, "playerId-1"}
		_, err := board.AddMove(move1)

		got := board.Moves
		want := []Move{}

		assertError(t, err, ErrBoardPositionInvalid)
		assertMoves(t, got, want)
	})
}

func TestGetEmptyPositions(t *testing.T) {
	t.Run("get empty positions from empty board", func(t *testing.T) {
		board := createNewBoard([]Move{})

		got := board.GetEmptyPositions()
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		assertPositions(t, got, want)
	})

	t.Run("get empty positions from board with 2 moves", func(t *testing.T) {
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 2, "playerId-2"}
		moves := []Move{move1, move2}
		board := createNewBoard(moves)

		got := board.GetEmptyPositions()
		want := []int{3, 4, 5, 6, 7, 8, 9}

		assertPositions(t, got, want)
	})

	t.Run("get empty positions from full board", func(t *testing.T) {
		board := createFullBoard()
		emptyPositions := board.GetEmptyPositions()

		got := len(emptyPositions)
		want := 0

		if got != want {
			t.Errorf("got %d positions want %d", got, want)
		}
	})
}

func TestGetTakenPositions(t *testing.T) {
	t.Run("get taken positions from board", func(t *testing.T) {
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 2, "playerId-2"}
		moves := []Move{move1, move2}
		board := createNewBoard(moves)
		board.AddMove(move1)
		board.AddMove(move2)
		takenPositions := board.GetTakenPositions()

		got := takenPositions
		want := []int{1, 2}

		assertPositions(t, got, want)
	})
}

/** Helper Function tests **/
func TestGetTakenPositionsHelper(t *testing.T) {

	t.Run("get taken positions from board using helper func", func(t *testing.T) {
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 2, "playerId-2"}
		moves := []Move{move1, move2}
		board := createNewBoard(moves)
		board.AddMove(move1)
		board.AddMove(move2)

		got := getTakenPositions(board)
		want := []int{1, 2}

		assertPositions(t, got, want)
	})
}

func TestGetEmptyPositionsHelper(t *testing.T) {
	t.Run("get empty positions from empty board using helper func", func(t *testing.T) {
		board := createNewBoard([]Move{})

		got := board.GetEmptyPositions()
		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		assertPositions(t, got, want)
	})

	t.Run("get empty positions from board with 2 moves using helper func", func(t *testing.T) {
		move1 := Move{"X", 1, "playerId-1"}
		move2 := Move{"O", 2, "playerId-2"}
		moves := []Move{move1, move2}
		board := createNewBoard(moves)

		got := getEmptyPositions(board)
		want := []int{3, 4, 5, 6, 7, 8, 9}

		assertPositions(t, got, want)
	})

	t.Run("get emptPositions from full board using helper func", func(t *testing.T) {
		board := createFullBoard()
		emptyPositions := getEmptyPositions(board)

		got := len(emptyPositions)
		want := 0

		if got != want {
			t.Errorf("got %d positions want %d", got, want)
		}
	})
}

func TestGetPositionMark(t *testing.T) {
	t.Run("get the value from an open position", func(t *testing.T) {
		board := createNewBoard([]Move{})

		got := board.GetPositionMark(1)
		want := "1"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("get the value from a filled position", func(t *testing.T) {
		move := Move{"X", 1, "playerId-1"}
		board := createNewBoard([]Move{move})

		got := board.GetPositionMark(1)
		want := "X"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestGetPositionMarkHelper(t *testing.T) {
	t.Run("get the value from an open position", func(t *testing.T) {
		board := createNewBoard([]Move{})

		got := getPositionMark(board, 1)
		want := "1"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("get the value from a filled position", func(t *testing.T) {
		move := Move{"X", 1, "playerId-1"}
		board := createNewBoard([]Move{move})

		got := getPositionMark(board, 1)
		want := "X"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func createNewBoard(moves []Move) Board {
	return Board{"1", 3, moves}
}

func createFullBoard() Board {
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
	board := createNewBoard(moves)
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

func assertPositions(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

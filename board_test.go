package board

import (
	"reflect"
	"testing"
)

func TestAddMove(t *testing.T) {
	t.Run("add move to empty board", func(t *testing.T) {
		board := Board{3, []Move{}}
		move := Move{"X", 1, "playerId-1"}
		board.AddMove(move)

		got := board.Moves
		want := []Move{move}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
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
		board := Board{3, []Move{move1, move2, move3, move4, move5, move6, move7, move8, move9}}
		board.AddMove(Move{"O", 10, "playerId-2"})

		got := board.Moves
		want := []Move{move1, move2, move3, move4, move5, move6, move7, move8, move9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("add move to a position that is already taken", func(t *testing.T) {

	})

	t.Run("adding a move to an invalid position", func(t *testing.T) {

	})
}

func createNewBoard() Board {
	return Board{3, []Move{}}
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
	board := Board{3, []Move{move1, move2, move3, move4, move5, move6, move7, move8, move9}}
	return board
}

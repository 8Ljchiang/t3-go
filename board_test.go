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

	})

	t.Run("add move to a position that is already taken", func(t *testing.T) {

	})
}

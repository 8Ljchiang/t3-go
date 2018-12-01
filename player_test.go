package main

import (
	"testing"
)

func TestGetMove(t *testing.T) {
	t.Run("get move from player", func(t *testing.T) {
		board := Board{"1", 3, []Move{}}
		player := Player{"playerId-1", "John", "X"}
		move := player.GetMove(board)

		got := move
		want := Move{"X", 1, "playerId-1"}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

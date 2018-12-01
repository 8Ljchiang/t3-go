package main

import "testing"

func TestCycleActivePlayerIndex(t *testing.T) {
	t.Run("it increments the index", func(t *testing.T) {
		playerCount := 2
		currentIndex := 0

		got := cycleActivePlayerIndex(currentIndex, playerCount)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("it resets the index to 0 when reaching the max player count", func(t *testing.T) {
		playerCount := 2
		currentIndex := 1

		got := cycleActivePlayerIndex(currentIndex, playerCount)
		want := 0

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestContainsPattern(t *testing.T) {
	t.Run("should return false for an empty board", func(t *testing.T) {
		moves := []Move{}
		board := Board{Id: "1", Size: 3, Moves: moves}

		got := containsWinningPattern(board, "X")
		want := false

		if got != want {
			t.Errorf("unwanted match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 1,5,9", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 1, playerId}
		move5 := Move{playerMarker, 5, playerId}
		move9 := Move{playerMarker, 9, playerId}
		moves := []Move{move1, move5, move9}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 3,5,7", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 3, playerId}
		move2 := Move{playerMarker, 5, playerId}
		move3 := Move{playerMarker, 7, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 1,2,3", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 1, playerId}
		move2 := Move{playerMarker, 2, playerId}
		move3 := Move{playerMarker, 3, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 4,5,6", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 4, playerId}
		move2 := Move{playerMarker, 5, playerId}
		move3 := Move{playerMarker, 6, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 7,8,9", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 7, playerId}
		move2 := Move{playerMarker, 8, playerId}
		move3 := Move{playerMarker, 9, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 1,4,7", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 1, playerId}
		move2 := Move{playerMarker, 4, playerId}
		move3 := Move{playerMarker, 7, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 2,5,8", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 2, playerId}
		move2 := Move{playerMarker, 5, playerId}
		move3 := Move{playerMarker, 8, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})

	t.Run("should return true for a board that has a matching pattern 3,6,9", func(t *testing.T) {
		playerId := "P1"
		playerMarker := "X"
		move1 := Move{playerMarker, 3, playerId}
		move2 := Move{playerMarker, 6, playerId}
		move3 := Move{playerMarker, 9, playerId}
		moves := []Move{move1, move2, move3}
		board := Board{"1", 3, moves}

		got := containsWinningPattern(board, playerMarker)
		want := true

		if got != want {
			t.Errorf("no pattern match found")
		}
	})
}

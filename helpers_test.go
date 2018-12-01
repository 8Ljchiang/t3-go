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
	t.Run("should return false for a board that has no matching patterns", func(t *testing.T) {

	})

	t.Run("should return true for a board that has a matching pattern", func(t *testing.T) {

	})
}

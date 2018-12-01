package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add board to empty bstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		boardStore := BoardStore{map[string]Board{}}

		_, addErr := boardStore.Add(board)
		got, gotErr := boardStore.Get(boardId)
		want := board

		assertError(t, addErr, nil)
		assertError(t, gotErr, nil)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("add board to store where board already exists", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		collection := map[string]Board{"1": board}
		boardStore := BoardStore{collection}

		_, err := boardStore.Add(board)
		want := ErrPlayerAlreadyExists

		assertError(t, err, want)
	})
}

func TestGet(t *testing.T) {
	t.Run("get board from bstore that exists", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		collection := map[string]Board{"1": board}
		boardStore := BoardStore{collection}

		got, err := boardStore.Get(playerId)
		want := player

		assertError(t, err, nil)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get board from bstore that doesn't exist", func(t *testing.T) {
		playerStore := BoardStore{map[string]Board{}}

		_, err := BoardStore.Get("1")
		want := ErrPlayerNotFound

		assertError(t, err, want)
	})
}

func TestUpate(t *testing.T) {
	t.Run("update board that exists in bstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard()
		updatedBoard := Board{Id: playerId, Size: 3, Moves: []Move{Move{"X", 1, "playerId-1"}}}
		boardStore := BoardStore{map[string]Player{"1": board}}

		_, err := boardStore.Update(updatedBoard)
		brd, _ := boardStore.Get(boardId)

		got := brd
		want := updatedBoard

		assertError(t, err, nil)
		assertPlayer(t, got, want)
	})

	t.Run("update board that doesn't exist in bstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard()
		boardStore := BoardStore{map[string]Board{}}

		_, updateErr := boardStore.Update(board)
		_, getErr := boardStore.Get(boardId)

		assertError(t, updateErr, ErrPlayerNotFound)
		assertError(t, getErr, ErrPlayerNotFound)
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove board that exists in pstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard()
		boardStore := BoardStore{map[string]board{"1": board}}

		boardStore.Remove(boardId)

		_, err := boardStore.Get(board.Id)

		assertError(t, err, ErrPlayerNotFound)
	})
}

func assertBoard(t *testing.T, got, want Board) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

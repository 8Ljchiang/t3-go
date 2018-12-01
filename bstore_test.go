package main

import (
	"reflect"
	"testing"
)

func TestBStoreAdd(t *testing.T) {
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
		board := createNewBoard([]Move{})
		collection := map[string]Board{"1": board}
		boardStore := BoardStore{collection}

		_, err := boardStore.Add(board)
		want := ErrBoardAlreadyExists

		assertError(t, err, want)
	})
}

func TestBStoreGet(t *testing.T) {
	t.Run("get board from bstore that exists", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		collection := map[string]Board{"1": board}
		boardStore := BoardStore{collection}

		got, err := boardStore.Get(boardId)
		want := board

		assertError(t, err, nil)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get board from bstore that doesn't exist", func(t *testing.T) {
		boardStore := BoardStore{map[string]Board{}}

		_, err := boardStore.Get("1")
		want := ErrBoardNotFound

		assertError(t, err, want)
	})
}

func TestBStoreUpate(t *testing.T) {
	t.Run("update board that exists in bstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		updatedBoard := Board{Id: boardId, Size: 3, Moves: []Move{Move{"X", 1, "playerId-1"}}}
		boardStore := BoardStore{map[string]Board{"1": board}}

		_, err := boardStore.Update(updatedBoard)
		brd, _ := boardStore.Get(boardId)

		got := brd
		want := updatedBoard

		assertError(t, err, nil)
		assertBoard(t, got, want)
	})

	t.Run("update board that doesn't exist in bstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		boardStore := BoardStore{map[string]Board{}}

		_, updateErr := boardStore.Update(board)
		_, getErr := boardStore.Get(boardId)

		assertError(t, updateErr, ErrBoardNotFound)
		assertError(t, getErr, ErrBoardNotFound)
	})
}

func TestBStoreRemove(t *testing.T) {
	t.Run("remove board that exists in pstore", func(t *testing.T) {
		boardId := "1"
		board := createNewBoard([]Move{})
		boardStore := BoardStore{map[string]Board{"1": board}}

		boardStore.Remove(boardId)

		_, err := boardStore.Get(board.Id)

		assertError(t, err, ErrBoardNotFound)
	})
}

func assertBoard(t *testing.T, got, want Board) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

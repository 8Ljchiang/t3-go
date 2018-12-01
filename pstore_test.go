package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("add player to empty pstore", func(t *testing.T) {

	})

	t.Run("add player to store where player already exists", func(t *testing.T) {

	})
}

func TestGet(t *testing.T) {
	t.Run("get player from pstore that exists", func(t *testing.T) {
		playerId := "1"
		player := Player{Id: playerId, Name: "John", Marker: "X"}
		collection := map[string]Player{
			"1": player,
		}
		playerStore := PlayerStore{collection}

		got, err := playerStore.Get(playerId)
		want := player

		assertError(t, err, nil)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get player from pstore that doesn't exist", func(t *testing.T) {
		playerStore := PlayerStore{map[string]Player{}}

		_, err := playerStore.Get("1")
		want := ErrPlayerNotFound

		assertError(t, err, want)
	})
}

func TestUpate(t *testing.T) {
	t.Run("update player that exists in pstore", func(t *testing.T) {

	})

	t.Run("update player that doesn't exist in pstore", func(t *testing.T) {

	})
}

func TestRemove(t *testing.T) {
	t.Run("remove player that exists in pstore", func(t *testing.T) {

	})

	t.Run("remove player that doesn't exist in pstore", func(t *testing.T) {

	})
}

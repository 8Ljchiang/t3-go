package main

import (
	"reflect"
	"testing"
)

func TestPStoreAdd(t *testing.T) {
	t.Run("add player to empty pstore", func(t *testing.T) {
		playerId := "1"
		player := Player{Id: playerId, Name: "John", Marker: "X"}
		playerStore := PlayerStore{map[string]Player{}}

		_, addErr := playerStore.Add(player)
		got, gotErr := playerStore.Get(playerId)
		want := player

		assertError(t, addErr, nil)
		assertError(t, gotErr, nil)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("add player to store where player already exists", func(t *testing.T) {
		playerId := "1"
		player := Player{Id: playerId, Name: "John", Marker: "X"}
		collection := map[string]Player{
			"1": player,
		}
		playerStore := PlayerStore{collection}

		_, err := playerStore.Add(player)
		want := ErrPlayerAlreadyExists

		assertError(t, err, want)
	})
}

func TestPStoreGet(t *testing.T) {
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

func TestPStoreUpate(t *testing.T) {
	t.Run("update player that exists in pstore", func(t *testing.T) {
		playerId := "1"
		player := Player{Id: playerId, Name: "John", Marker: "X"}
		updatedPlayer := Player{Id: playerId, Name: "Mike", Marker: "O"}
		playerStore := PlayerStore{map[string]Player{
			"1": player,
		}}

		_, err := playerStore.Update(updatedPlayer)
		plyr, _ := playerStore.Get(playerId)

		got := plyr
		want := updatedPlayer

		assertError(t, err, nil)
		assertPlayer(t, got, want)
	})

	t.Run("update player that doesn't exist in pstore", func(t *testing.T) {
		playerId := "1"
		player := Player{Id: playerId, Name: "John", Marker: "X"}
		playerStore := PlayerStore{map[string]Player{}}

		_, updateErr := playerStore.Update(player)
		_, getErr := playerStore.Get(playerId)

		assertError(t, updateErr, ErrPlayerNotFound)
		assertError(t, getErr, ErrPlayerNotFound)
	})
}

func TestPStoreRemove(t *testing.T) {
	t.Run("remove player that exists in pstore", func(t *testing.T) {
		playerId := "1"
		player := Player{Id: playerId, Name: "John", Marker: "X"}
		playerStore := PlayerStore{map[string]Player{"1": player}}

		playerStore.Remove(playerId)

		_, err := playerStore.Get(player.Id)

		assertError(t, err, ErrPlayerNotFound)
	})
}

func assertPlayer(t *testing.T, got, want Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

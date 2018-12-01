package main

import "testing"

type BoardStoreMock struct{}

func (bsm BoardStoreMock) Get(boardId string) (Board, error) {
	board1Id := "B1"
	board1 := Board{Id: board1Id, Size: 3, Moves: []Move{}}
	return board1, nil
}

type FullBoardStoreMock struct{}

func (fbsm FullBoardStoreMock) Get(boardId string) (Board, error) {
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
	board1Id := "B1"
	board1 := Board{Id: board1Id, Size: 3, Moves: moves}
	return board1, nil
}

func TestGetGameStatus(t *testing.T) {
	t.Run("should return the status of IN_PROGRESS when board is not full", func(t *testing.T) {
		boardStoreMock := BoardStoreMock{}
		newGame := createNewGame()

		got := getGameStatus(&newGame, boardStoreMock)
		want := STATUS_IN_PROGRESS

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("should return the status of DRAW when the board is full and no winner", func(t *testing.T) {
		boardStoreMock := FullBoardStoreMock{}
		newGame := createNewGame()

		got := getGameStatus(&newGame, boardStoreMock)
		want := STATUS_DRAW

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("should return the status of WINNER when there is a matching pattern", func(t *testing.T) {
		boardStoreMock := FullBoardStoreMock{}
		newGame := createNewGame()

		got := getGameStatus(&newGame, boardStoreMock)
		want := STATUS_WINNER

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func createNewGame() Game {
	player1Id := "P1"
	player2Id := "P2"
	board1Id := "B1"
	game1Id := "G1"
	newGame := Game{Id: game1Id, BoardId: board1Id, Players: []string{player1Id, player2Id}, ActivePlayerIndex: 0}
	return newGame
}

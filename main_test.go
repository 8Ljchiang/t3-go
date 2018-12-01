package main

import "testing"

type BoardStoreMock struct{}

func (bsm BoardStoreMock) Get(boardId string) (Board, error) {
	board1Id := "B1"
	board1 := Board{Id: board1Id, Size: 3, Moves: []Move{}}
	return board1, nil
}

type WinBoardStoreMock struct{}

func (wbsm WinBoardStoreMock) Get(boardId string) (Board, error) {
	p1 := "P1"
	p2 := "P2"

	move1 := Move{M_1, 1, p1}
	move2 := Move{M_2, 2, p2}
	move3 := Move{M_1, 3, p1}
	move4 := Move{M_2, 4, p2}
	move5 := Move{M_1, 5, p1}
	move6 := Move{M_2, 6, p2}
	move7 := Move{M_1, 7, p1}
	move8 := Move{M_2, 8, p2}
	move9 := Move{M_1, 9, p1}
	moves := []Move{move1, move2, move3, move4, move5, move6, move7, move8, move9}
	board1Id := "B1"
	board1 := Board{Id: board1Id, Size: 3, Moves: moves}
	return board1, nil
}

type DrawBoardStoreMock struct{}

func (dbsm DrawBoardStoreMock) Get(boardId string) (Board, error) {
	p1 := "P1"
	p2 := "P2"
	move1 := Move{M_1, 1, p1}
	move2 := Move{M_1, 2, p1}
	move3 := Move{M_2, 3, p2}
	move4 := Move{M_2, 4, p2}
	move5 := Move{M_2, 5, p1}
	move6 := Move{M_1, 6, p2}
	move7 := Move{M_1, 7, p1}
	move8 := Move{M_1, 8, p1}
	move9 := Move{M_2, 9, p2}
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
		boardStoreMock := DrawBoardStoreMock{}
		newGame := createNewGame()

		got := getGameStatus(&newGame, boardStoreMock)
		want := STATUS_DRAW

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("should return the status of WINNER when there is a matching pattern", func(t *testing.T) {
		boardStoreMock := WinBoardStoreMock{}
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	STATUS_NEW         = "NEW"
	STATUS_IN_PROGRESS = "IN_PROGRESS"
	STATUS_DRAW        = "DRAW"
	STATUS_WINNER      = "WINNER"
)

type Game struct {
	Id                string
	BoardId           string
	Players           []string
	ActivePlayerIndex int
}

func main() {
	player1Id := "P1"
	player2Id := "P2"
	player1 := Player{Id: player1Id, Name: "John", Marker: "X"}
	player2 := Player{Id: player2Id, Name: "Tom", Marker: "O"}
	playerStore := PlayerStore{map[string]Player{}}
	playerStore.Add(player1)
	playerStore.Add(player2)

	board1Id := "B1"
	board1 := Board{Id: board1Id, Size: 3, Moves: []Move{}}
	boardStore := BoardStore{map[string]Board{}}
	boardStore.Add(board1)

	game1Id := "G1"
	newGame := Game{Id: game1Id, BoardId: board1Id, Players: []string{player1Id, player2Id}, ActivePlayerIndex: 0}

	reader := bufio.NewReader(os.Stdin)

	playGame(reader, &newGame, &boardStore, &playerStore)
}

func playGame(inputReader *bufio.Reader, game *Game, bs *BoardStore, ps *PlayerStore) {
	fmt.Print("Welcome to Tic Tac Toe\n\n")
	fmt.Print("1. Enter a position 1-9.\n")
	fmt.Print("2. Match 3 in a row to win.\n")

	board, bsErr := bs.Get(game.BoardId)
	if bsErr == nil {
		renderBoard(board)
	}

	for {
		status := getGameStatus(game, bs)
		playRound(inputReader, status, game, bs, ps)

		if status == STATUS_DRAW || status == STATUS_WINNER {
			break
		}
	}
}

func playRound(inputReader *bufio.Reader, status string, game *Game, bs *BoardStore, ps *PlayerStore) {
	processGameStatus(inputReader, status, game, bs, ps)
}

func getGameStatus(game *Game, bs Store) string {
	board, err := bs.Get(game.BoardId)

	if err == nil {
		if len(getEmptyPositions(board)) == 0 {
			return STATUS_DRAW
		} else if containsWinningPattern(board, M_1) || containsWinningPattern(board, M_2) {
			return STATUS_WINNER
		}
	}
	return STATUS_IN_PROGRESS
}

func processGameStatus(inputReader *bufio.Reader, status string, game *Game, bs *BoardStore, ps *PlayerStore) {
	switch status {
	case STATUS_DRAW:
		fmt.Print("The game state is " + STATUS_DRAW + "\n")
	case STATUS_IN_PROGRESS:
		fmt.Print("The game state is " + STATUS_IN_PROGRESS + "\n")
		board, bsErr := bs.Get(game.BoardId)
		if bsErr == nil {
			currentPlayerId := game.Players[game.ActivePlayerIndex]
			currentPlayer, psErr := ps.Get(currentPlayerId)
			input := getInput(inputReader, currentPlayer.Name+" ("+currentPlayer.Marker+") "+"enter a position:")
			position, parseError := strconv.Atoi(input)
			if psErr == nil && isValidPosition(&board, position) {
				if parseError == nil {
					fmt.Printf("input: %s, position: %d\n", input, position)
					newMove := Move{Mark: currentPlayer.Marker, Position: position, PlayerId: currentPlayerId}
					newMoves := append(board.Moves, newMove)
					newBoard := Board{Id: board.Id, Size: board.Size, Moves: newMoves}
					game.ActivePlayerIndex = cycleActivePlayerIndex(game.ActivePlayerIndex, len(game.Players))
					_, bsUErr := bs.Update(newBoard)
					if bsUErr == nil {
						clearConsole()
						fmt.Print("Tic Tac Toe\n\n")
						renderBoard(newBoard)
					}
				}
			} else {
				fmt.Print("Position is not valid\n")
			}
		}
	case STATUS_WINNER:
		fmt.Print("The game state is " + STATUS_WINNER + "\n")
	default:
		fmt.Print("The game state is UNKNOWN\n")
	}
}

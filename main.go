package main

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
	board1 := Board{Id: board1Id, Moves: []Move{}}
	boardStore := BoardStore{map[string]Board{}}
	boardStore.Add(board1)

	game1Id := "G1"
	newGame := Game{Id: game1Id, BoardId: board1Id, Players: []string{player1Id, player2Id}, ActivePlayerIndex: 0}

	play(newGame)
}

func play(game Game) {

}

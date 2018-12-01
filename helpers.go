package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func renderBoard(board Board) {
	fmt.Print("\n")
	fmt.Printf(" %s    %s    %s\n", board.GetPositionMark(1), board.GetPositionMark(2), board.GetPositionMark(3))
	fmt.Print("\n")
	fmt.Printf(" %s    %s    %s\n", board.GetPositionMark(4), board.GetPositionMark(5), board.GetPositionMark(6))
	fmt.Print("\n")
	fmt.Printf(" %s    %s    %s\n", board.GetPositionMark(7), board.GetPositionMark(8), board.GetPositionMark(9))
	fmt.Print("\n")
}

func clearConsole() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func getInput(inputReader *bufio.Reader, prompt string) (input string) {
	fmt.Print(prompt)
	input, _ = inputReader.ReadString('\n')
	return strings.TrimRight(input, "\n")
}

func cycleActivePlayerIndex(activePlayerIndex int, playerCount int) int {
	if activePlayerIndex < playerCount-1 {
		return activePlayerIndex + 1
	}
	return 0
}

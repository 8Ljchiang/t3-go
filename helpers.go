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

func containsWinningPattern(board Board, mark string) bool {
	winningPatterns := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 6, 9},
		[]int{1, 5, 9},
		[]int{3, 5, 7},
	}
	markerPositions := []int{}

	for _, move := range board.Moves {
		if move.Mark == mark {
			markerPositions = append(markerPositions, move.Position)
		}
	}

	for _, pattern := range winningPatterns {
		matchingPositions := []int{}
		for _, winningPos := range pattern {
			if containsPosition(markerPositions, winningPos) {
				matchingPositions = append(matchingPositions, winningPos)
			}
		}
		if len(matchingPositions) == len(pattern) {
			return true
		}
	}
	return false
}

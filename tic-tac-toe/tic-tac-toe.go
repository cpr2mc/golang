// https://www.youtube.com/watch?v=MnpjkEBGhtY

package main

import "fmt"

xs := []int{1, 2, 3}
ys := []int{1, 2, 3}

gameBoard := map[string]string{
	"(1, 1)": "_",
	"(2, 1)": "_",
	"(3, 1)": "_",
	"(1, 2)": "_",
	"(2, 2)": "_",
	"(3, 2)": "_",
	"(1, 3)": "_",
	"(2, 3)": "_",
	"(3, 3)": "_",
}

func main() {
	
	printGameBoard()
	updateBoard := getPlayerOne()
	
}

func printGameBoard() {
	for _, y := range ys {
		for _, x := range xs {
			key := fmt.Sprintf("(%d, %d)", x, y)
			val := gameBoard[key]
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func updateGameBoard(update string) {
	// this will update game board
}

func getPlayerOne() string {
	var playerOneMove string
	fmt.Println("Player 1: Enter your move (x, y)")
	fmt.Scan(&playerOneMove)
	return playerOneMove
}

func getPlayerTwo() string {
	var playerTwoMove string
	fmt.Println("Player 2: Enter your move (x, y)")
	fmt.Scan(&playerTwoMove)
	return playerTwoMove
}

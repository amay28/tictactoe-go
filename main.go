package main

import (
	"fmt"
	"os"
)

// Board : Tictactoe matrix
type Board struct {
	board [3][3]string
	count int
}

func printBoard(matrix [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s\t", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func checkRow(board [3][3]string, row int, player string) bool {
	flag := true
	for i := 0; i < 3; i++ {
		if board[row][i] != player {
			flag = false
			break
		}
	}
	return flag
}

func checkCol(board [3][3]string, col int, player string) bool {
	flag := true
	for i := 0; i < 3; i++ {
		if board[i][col] != player {
			flag = false
			break
		}
	}
	return flag
}

func checkOffDiagonal(board [3][3]string, player string) bool {
	flag := true
	for i := 0; i < 3; i++ {
		if board[i][2-i] != player {
			flag = false
			break
		}
	}
	return flag
}

func checkDiagonal(board [3][3]string, player string) bool {
	flag := true
	for i := 0; i < 3; i++ {
		if board[i][i] != player {
			flag = false
			break
		}
	}
	return flag
}

func check(board [3][3]string, player string, row int, col int) bool {
	if checkRow(board, row, player) {
		return true
	}
	if checkCol(board, col, player) {
		return true
	}
	if row == col {
		if checkDiagonal(board, player) {
			return true
		}
	}
	if row-col == 2 || row-col == -2 {
		if checkOffDiagonal(board, player) {
			return true
		}
	}
	return false
}

func main() {
	board := Board{
		board: [3][3]string{[3]string{"-", "-", "-"}, [3]string{"-", "-", "-"}, [3]string{"-", "-", "-"}},
		count: 0,
	}
	printBoard(board.board)
	player := 0
	for {
		var index int
	Input:
		fmt.Printf("Player %d Enter the index\n\n", player+1)
		fmt.Scan(&index)
		row := int((index - 1) / 3)
		col := (index - 1) % 3
		if board.board[row][col] != "-" {
			fmt.Println("Index already filled")
			goto Input
		}
		playerChar := "X"
		if player == 0 {
			board.board[row][col] = "X"
		} else {
			board.board[row][col] = "O"
			playerChar = "O"
		}
		board.count = board.count + 1
		res := check(board.board, playerChar, row, col)
		if res {
			fmt.Printf("Player %d is the winner\n", player+1)
			os.Exit(0)
		}
		if board.count == 9 {
			fmt.Println("Draw Match")
			os.Exit(0)
		}
		player = (player + 1) % 2
		printBoard(board.board)
	}
}

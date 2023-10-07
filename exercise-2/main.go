package main

import "fmt"

// We'll build a simple tic tac toe game
// The game will be played in the terminal

// The game will be played by two players
// The players will take turns to play
// The game will be played on a 3x3 board

// The game will be played until there is a winner or a draw

type Game struct {
	Board        [3][3]string
	CurrentPlayer string
}

func NewGame() *Game {
	return &Game{
		Board: [3][3]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
		CurrentPlayer: "X",
	}
}

func main() {
	currentGame := NewGame()


	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("-----------------------")

	for i := 0; i < 9; i++ {
		GetMove(currentGame)

		if CheckWinner(currentGame) {
			fmt.Println("Winner is", currentGame.CurrentPlayer)
			break
		}
		SwitchPlayer(currentGame)
	}
	PrintBoard(currentGame)
}


func GetMove(game *Game) {
	var row, col int
	for {
		PrintBoard(game)
		fmt.Printf("%s's turn. Enter your move (row col): ", game.CurrentPlayer)
		_, err := fmt.Scan(&row, &col)

		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || game.Board[row][col] != " " {
			fmt.Println("Invalid move. Try again.")
			continue
		}
		game.Board[row][col] = game.CurrentPlayer
		break
	}
}

func SwitchPlayer(game *Game) {
	if game.CurrentPlayer == "X" {
		game.CurrentPlayer = "O"
	} else {
		game.CurrentPlayer = "X"
	}
}

func PrintBoard(game *Game) {
	for _, row := range game.Board {
		for _, cell := range row {
			fmt.Printf("|%s", cell)
		}
		fmt.Println("|")
	}
}

func CheckWinner(game *Game) (bool) {
	if(game.Board[0][0]==game.Board[0][1] && game.Board[0][1]==game.Board[0][2] && game.Board[0][0]!=" "){
		return true
	}
	if(game.Board[1][0]==game.Board[1][1] && game.Board[1][1]==game.Board[1][2] && game.Board[1][0]!=" "){
		return true
	}
	if(game.Board[2][0]==game.Board[2][1] && game.Board[2][1]==game.Board[2][2] && game.Board[2][0]!=" "){
		return true
	}
	if(game.Board[0][0]==game.Board[1][0] && game.Board[1][0]==game.Board[2][0] && game.Board[0][0]!=" "){
		return true
	}
	if(game.Board[1][0]==game.Board[1][1] && game.Board[1][1]==game.Board[1][2] && game.Board[1][0]!=" "){
		return true
	}
	if(game.Board[2][0]==game.Board[2][1] && game.Board[2][1]==game.Board[2][2] && game.Board[2][0]!=" "){
		return true
	}
	if(game.Board[0][0]==game.Board[1][1] && game.Board[1][1]==game.Board[2][2] && game.Board[0][0]!=" "){
		return true
	}
	if(game.Board[0][2]==game.Board[1][1] && game.Board[1][1]==game.Board[2][0] && game.Board[0][2]!=" "){
		return true
	}

	return false

}
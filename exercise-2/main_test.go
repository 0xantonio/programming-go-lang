package main

import (
	"testing"
)

func TestCheckWinner(t *testing.T) {
	tests := []struct {
		name          string
		game          *Game
		expectedIsWon bool
	}{
		{
			name: "Row win",
			game: &Game{
				Board: [3][3]string{
					{"X", "X", "X"},
					{" ", " ", " "},
					{" ", " ", " "},
				},
				CurrentPlayer: "X",
			},
			expectedIsWon: true,
		},
		{
			name: "Diagonal win",
			game: &Game{
				Board: [3][3]string{
					{"X", " ", " "},
					{" ", "X", " "},
					{" ", " ", "X"},
				},
				CurrentPlayer: "X",
			},
			expectedIsWon: true,
		},
		{
			name: "Couner diagonal win",
			game: &Game{
				Board: [3][3]string{
					{" ", " ", "X"},
					{" ", "X", " "},
					{"X", " ", " "},
				},
				CurrentPlayer: "X",
			},
			expectedIsWon: true,
		},
		{
			name: "No win",
			game: &Game{
				Board: [3][3]string{
					{"X", " ", " "},
					{" ", " ", " "},
					{" ", " ", " "},
				},
				CurrentPlayer: "X",
			},
			expectedIsWon: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isWon := CheckWinner(tt.game)
			if isWon != tt.expectedIsWon {
				t.Errorf("Expected isWon to be %v for %v, got %v", tt.expectedIsWon, tt.name, isWon)
			}
		})
	}
}

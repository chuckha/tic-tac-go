package tictactoe

import (
	"testing"
)

var testCase = []struct {
	index    int
	expected []int
}{
	{0, []int{0, 1, 2}},
	{3, []int{0, 3, 6}},
	{6, []int{0, 4, 8}},
	{7, []int{2, 4, 6}},
}

func TestGenerateWinCombos(t *testing.T) {
	combos := GenerateWinCombos(3)
	for _, tc := range testCase {
		if tc.index > len(combos) {
			t.Errorf("Combos did not provide enough combos\n")
		}
		if !slicesEqual(combos[tc.index], tc.expected) {
			t.Errorf("Expectred %v. Got %v.", tc.expected, combos[tc.index])
		}
	}
}

var winnerTestCase = []struct {
	ttt    TicTacToe
	winner int
}{
	{
		ttt: TicTacToe{
			1, 2, 1,
			2, 2, 1,
			2, 1, 1,
		},
		winner: 1,
	},
	{
		ttt: TicTacToe{
			1, 0, 0,
			0, 0, 0,
			0, 0, 0,
		},
		winner: 0,
	},
	{
		ttt: TicTacToe{
			1, 1, 1,
			2, 2, 1,
			2, 2, 0,
		},
		winner: 1,
	},
}

func TestWinner(t *testing.T) {
	for _, wtc := range winnerTestCase {
		winner := Winner(wtc.ttt)
		if winner != wtc.winner {
			t.Errorf("Expected %v. Got %v.", wtc.winner, winner)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

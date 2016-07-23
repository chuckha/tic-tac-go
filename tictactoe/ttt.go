package tictactoe

import (
	"bytes"
	"fmt"
)

// TicTacToe is the representation of the game and its rules.
type TicTacToe []int

const (
	// Size is the width & height of the board
	Size = 3
)

// New creates a TicTacToe
func New() TicTacToe {
	return make([]int, Size*Size)
}

// Set a value on the board with some rules
func (t TicTacToe) Set(idx, val int) error {
	if idx > len(t) {
		return fmt.Errorf("Move cannot be outside of board")
	}
	if t[idx] != 0 {
		return fmt.Errorf("Cannot set an already set spot")
	}
	t[idx] = val
	return nil
}

// Get returns the value of a spot on the board
func (t TicTacToe) Get(idx int) int {
	return t[idx]
}

func (t TicTacToe) String() string {
	var buffer bytes.Buffer
	for i := range t {
		if i%Size == 0 {
			buffer.WriteString("\n")
		}
		str := "-"
		if t[i] != 0 {
			str = fmt.Sprintf("%d", t[i])
		}

		buffer.WriteString(str)
	}
	buffer.WriteString("\n")
	return buffer.String()
}

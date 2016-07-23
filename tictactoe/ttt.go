package tictactoe

import (
	"bytes"
	"fmt"
)

// TicTacToe is the representation of the game and its rules
type TicTacToe []int

// TODO: parameterize Size instead of making it a const
const (
	Size = 3
)

// New creates a new TicTacToe game struct
func New() TicTacToe {
	return make([]int, Size*Size)
}

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

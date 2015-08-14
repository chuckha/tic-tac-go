package state

import (
	"fmt"
)

type State struct {
	board []string
}

func NewState() (*State, error) {
	b := make([]string, 9)
	for i := range b {
		b[i] = "."
	}
	return &State{
		board: b,
	}, nil
}
func NewStateWithArray(initial []string) (*State, error) {
	return &State{
		board: initial,
	}, nil
}

func (s *State) Update(newState *State) error {
	for i, v := range newState.board {
		if v != "" {
			if s.board[i] != "." {
				return fmt.Errorf("Space is already taken")
			}
			s.board[i] = newState.board[i]
		}
	}
	return nil
}

func (s *State) String() string {
	return fmt.Sprintf(" %s | %s | %s \n---|---|---\n %s | %s | %s \n---|---|---\n %s | %s | %s\n",
		s.board[0],
		s.board[1],
		s.board[2],
		s.board[3],
		s.board[4],
		s.board[5],
		s.board[6],
		s.board[7],
		s.board[8],
	)
}
func (s *State) movesRemaining() int {
	remaining := 0
	for _, v := range s.board {
		if v == "." {
			remaining += 1
		}
	}
	return remaining
}

func (s *State) IsOver() bool {
	winStates := [][]int{
		[]int{0, 1, 2},
		[]int{0, 4, 8},
		[]int{0, 3, 6},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
		[]int{2, 4, 6},
	}
	// did anyone win?
	for _, winState := range winStates {

		if s.board[winState[0]] == s.board[winState[1]] &&
			s.board[winState[1]] == s.board[winState[2]] &&
			s.board[winState[0]] != "." {
			return true
		}
	}
	// can we take more turns?
	if s.movesRemaining() == 0 {
		return true
	}

	return false
}

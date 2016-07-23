package tictactoe

import (
	"testing"
)

func TestSet(t *testing.T) {
	ttt := New()
	ttt.Set(0, 1)
	if ttt[0] != 1 {
		t.Errorf("0th index should be 1")
	}
}

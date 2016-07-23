package player

import (
	"testing"
)

func TestPlayerNumberCannotBeZero(t *testing.T) {
	_, err := New(0)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}

func TestPlayerId(t *testing.T) {
	p, _ := New(1)
	if p.Id() != 1 {
		t.Error("Did not get player id correctly\n")
	}
}

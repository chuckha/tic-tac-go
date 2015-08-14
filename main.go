package main

import (
	"fmt"
	"github.com/chuckha/tictactoe/state"
	"math/rand"
)

func game() {
	// while game is not over
	// get user input
	// update state
}

type Player interface {
	GetSymbol() string
	GetInput() int
	ValidInput() error
}
type ComputerInput struct{}

func (c *ComputerInput) GetSymbol() string {
	return "X"
}

func (c *ComputerInput) GetInput() int {
	return rand.Intn(9)
}
func (c *ComputerInput) ValidInput() error {
	return nil
}

type UserInput struct {
	input int
}

func (u *UserInput) GetSymbol() string {
	return "O"
}
func (u *UserInput) GetInput() int {
	fmt.Println("What is your input?")
	fmt.Scanf("%d\n", &u.input)
	return u.input
}
func (u *UserInput) ValidInput() error {
	if u.input < 0 || u.input > 8 {
		return fmt.Errorf("Input must be between 0 and 9")
	}
	return nil
}

func main() {
	s, _ := state.NewState()
	u := &UserInput{}
	inputs := []Player{&UserInput{}, &ComputerInput{}}
	turn := 0
	for !s.IsOver() {
		fmt.Println(s)
		input := inputs[turn].GetInput()
		err := u.ValidInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		initial := make([]string, 9)
		initial[input] = inputs[turn].GetSymbol()
		fmt.Println(initial)
		newState, _ := state.NewStateWithArray(initial)
		err = s.Update(newState)
		if err != nil {
			// if you get here you're cheating :)
			continue
		}
		turn += 1
		turn %= 2
	}
	fmt.Println(s)
}

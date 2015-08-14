package main

import (
	"fmt"
	"github.com/chuckha/tictactoe/player"
	"github.com/chuckha/tictactoe/state"
	"math/rand"
)

func main() {
	s, _ := state.NewState()
	u := &UserInput{}
	inputs := []player.Player{&player.UserInput{}, &player.ComputerInput{}}
	turn := 0
	for !s.IsOver() {
		fmt.Println(s)
		input := inputs[turn].GetInput()
		err := u.ValidInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// TODO: pull out 9 to somewhere else
		initial := make([]string, 9)
		initial[input] = inputs[turn].GetSymbol()
		fmt.Println(initial)

		// TODO: catch errors that might happen here
		newState, _ := state.NewStateWithArray(initial)
		err = s.Update(newState)
		if err != nil {
			// if you get here you're cheating :)
			continue
		}
		turn += 1

		// TODO: make this mod number of players
		turn %= 2
	}
	fmt.Println(s)
}

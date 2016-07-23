package main

import (
	"fmt"
	"github.com/chuckha/tic-tac-go/cli"
)

func main() {
	game, err := cli.New()
	if err != nil {
		panic(err)
	}

	for game.InProgress() {
		fmt.Println("Current board:")
		fmt.Println(game)
		err := game.NextTurn()
		if err != nil {
			fmt.Println("Illegal move!")
			fmt.Println(err)
			continue
		}
	}
	winner := game.Winner()
	fmt.Printf("Game over! %d wins\n", winner)
	fmt.Println("Final board:")
	fmt.Println(game)
}

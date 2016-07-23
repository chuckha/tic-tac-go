package input

import (
	"fmt"
	"github.com/chuckha/tic-tac-go/player"
	"github.com/chuckha/tic-tac-go/tictactoe"
)

func Instructions(ttt tictactoe.TicTacToe, p *player.Player) {
	fmt.Println("Player %d input your move:")
	fmt.Println(ttt)
}

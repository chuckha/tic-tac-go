package cli

import (
	"fmt"
	"github.com/chuckha/tic-tac-go/input"
	"github.com/chuckha/tic-tac-go/player"
	"github.com/chuckha/tic-tac-go/tictactoe"
)

// Game ties together all the pieces.
// Game expects IO to happen on the command line
type Game struct {
	ttt                tictactoe.TicTacToe
	players            []*player.Player
	currentPlayerIndex int
}

// New creates a new Game.
func New() (*Game, error) {
	player1, err := player.New(1)
	if err != nil {
		return nil, err
	}
	player2, err := player.New(2)
	if err != nil {
		return nil, err
	}
	return &Game{
		ttt: tictactoe.New(),
		players: []*player.Player{
			player1,
			player2,
		},
		currentPlayerIndex: 0,
	}, nil
}

// NextTurn gets the user input and returns an error if the input is not valid
func (g *Game) NextTurn() error {
	fmt.Printf("Get input from %d", g.players[g.currentPlayerIndex].Id())
	id, index, err := input.GetInput(g.players[g.currentPlayerIndex])
	if err != nil {
		return err
	}
	err = g.ttt.Set(index, id)
	if err != nil {
		return err
	}
	g.updateCurrentPlayerIndex()
	return nil
}

func (g *Game) updateCurrentPlayerIndex() {
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % 2
}

// InProgress tells us if a Game is currently underway or if the game is over
func (g *Game) InProgress() bool {
	return tictactoe.Winner(g.ttt) == 0
}

func (g *Game) String() string {
	return g.ttt.String()
}

// Winner tells us who won.
func (g *Game) Winner() int {
	return tictactoe.Winner(g.ttt)
}

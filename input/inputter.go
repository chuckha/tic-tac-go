package input

import (
	"fmt"
	"github.com/chuckha/tic-tac-go/player"
)

// GetInput gets input from somehere and returns some input
// Get input takes a player and returns a value, an index and an error
func GetInput(p *player.Player) (int, int, error) {
	// get the user input
	var userInputIndex int
	fmt.Scanln(&userInputIndex)

	err := ValidateInput(userInputIndex)
	if err != nil {
		return 0, 0, err
	}

	return p.Id(), userInputIndex, nil
}

package input

import (
	"fmt"
	"github.com/chuckha/tic-tac-go/player"
)

// GetInput gets input from somewhere and returns some values.
// GetInput returns the index of the board to set, the value to set it to and any error encountered.
func GetInput(p *player.Player) (int, int, error) {
	var userInputIndex int
	fmt.Scanln(&userInputIndex)

	err := ValidateInput(userInputIndex)
	if err != nil {
		return 0, 0, err
	}

	return p.Id(), userInputIndex, nil
}

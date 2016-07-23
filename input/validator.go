package input

import (
	"fmt"
)

// The user input needs to be a number > 0
func ValidateInput(input int) error {
	if input < 0 {
		return fmt.Errorf("Input must be > 0")
	}
	return nil
}

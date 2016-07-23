package input

import (
	"testing"
)

func TestValidator(t *testing.T) {
	err := ValidateInput(-1)
	if err == nil {
		t.Errorf("Expected an error with input of -1")
	}
}

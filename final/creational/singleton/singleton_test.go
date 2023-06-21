package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	firstCall := GetNewGame()
	secondCall := GetNewGame()

	if firstCall != secondCall {
		t.Error("pointers must be equal!")
	}

	firstCall.SetDifficulty(3)
	if firstCall != secondCall {
		t.Error("pointers must be equal!")
	}
}

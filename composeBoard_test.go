package gofen

import (
	"testing"
)

func TestComposeBoardArrayToString(t *testing.T) {
	initialBoardArray := []string{
		"R", "N", "B", "Q", "K", "B", "N", "R",
		"P", "P", "P", "P", "P", "P", "P", "P",
		".", ".", ".", ".", ".", ".", ".", ".",
		".", ".", ".", ".", ".", ".", ".", ".",
		".", ".", ".", ".", ".", ".", ".", ".",
		".", ".", ".", ".", ".", ".", ".", ".",
		"p", "p", "p", "p", "p", "p", "p", "p",
		"r", "n", "b", "q", "k", "b", "n", "r",
	}
	expected := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

	if ret := composeBoardArrayToString(initialBoardArray); ret != expected {
		t.Errorf("mapPositionToBoardIndex() = %q, want %q", ret, expected)
	}
}

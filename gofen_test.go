package gofen

import (
	"reflect"
	"testing"
)

func TestNewFen(t *testing.T) {
	expected := Fen{
		fenString:          "rnbqkbnr w KQkq - 0 1",
		board:              []string{"r", "n", "b", "q", "k", "b", "n", "r"},
		colorToMove:        "w",
		availableCastlings: "KQkq",
		halfMoveClock:      0,
		fullMoveClock:      1,
	}

	if ret := NewFen("rnbqkbnr w KQkq - 0 1"); !reflect.DeepEqual(ret, expected) {
		t.Errorf("NewFen() = %q, want %q", ret, expected)
	}
}
func TestAddMove(t *testing.T) {
	expected := "rnbqkbnr/pppppppp/8/8/P7/8/.PPPPPPP/RNBQKBNR w KQkq - 0 1"
	f := NewFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	f.AddMove("a2a4")

	if ret := f.fenString; ret != expected {
		t.Errorf("AddMove() = %q, want %q", ret, expected)
	}
}

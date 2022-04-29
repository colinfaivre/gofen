package gofen

import (
	"reflect"
	"testing"
)

func TestGetRanks(t *testing.T) {
	expected := []string{"rnbqkbnr", "pppppppp", "8", "8", "8", "8", "PPPPPPPP", "RNBQKBNR"}

	if ret := getRanks("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"); !reflect.DeepEqual(ret, expected) {
		t.Errorf("getRanks() = %q, want %q", ret, expected)
	}
}
func TestGetRankCells(t *testing.T) {
	expected := []string{"r", "n", "b", "q", "k", "b", "n", "r"}

	if ret := getRankCells("rnbqkbnr"); !reflect.DeepEqual(ret, expected) {
		t.Errorf("getRankCells() = %q, want %q", ret, expected)
	}
}
func TestParseBoard(t *testing.T) {
	expected := []string{".", ".", ".", ".", ".", ".", ".", ".", "p", "p", "p", "p", "p", "p", "p", "p", "r", "n", "b", "q", "k", "b", "n", "r"}

	if ret := parseBoard("rnbqkbnr/pppppppp/8"); !reflect.DeepEqual(ret, expected) {
		t.Errorf("parseBoard() = %q, want %q", ret, expected)
	}
}
func TestConvertNumbersToPoints(t *testing.T) {
	expected := "........"

	if ret := convertNumbersToPoints("8"); !reflect.DeepEqual(ret, expected) {
		t.Errorf("convertNumbersToPoints() = %q, want %q", ret, expected)
	}

	expected2 := "..r..r.."

	if ret := convertNumbersToPoints("2r2r2"); !reflect.DeepEqual(ret, expected2) {
		t.Errorf("convertNumbersToPoints() = %q, want %q", ret, expected2)
	}
}

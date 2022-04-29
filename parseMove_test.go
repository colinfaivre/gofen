package gofen

import (
	"testing"
)

func TestMapPositionToBoardIndex(t *testing.T) {
	expected := 0
	expected2 := 63

	if ret := mapPositionToBoardIndex("a1"); ret != expected {
		t.Errorf("mapPositionToBoardIndex() = %q, want %q", ret, expected)
	}
	if ret := mapPositionToBoardIndex("h8"); ret != expected2 {
		t.Errorf("mapPositionToBoardIndex() = %q, want %q", ret, expected2)
	}
}

func TestParseMove(t *testing.T) {
	expected1 := 8
	expected2 := 16

	if ret1, ret2 := parseMove("a2a3"); ret1 != expected1 || ret2 != expected2 {
		t.Errorf("mapPositionToBoardIndex() = %q %q, want %q %q", ret1, ret2, expected1, expected2)
	}
}
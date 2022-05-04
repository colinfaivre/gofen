package gofen

import (
	"strconv"
	"strings"
)

type Fen struct {
	fenString          string
	board              []string
	colorToMove        string
	availableCastlings string
	halfMoveClock      int
	fullMoveClock      int
}

func NewFen(fen string) Fen {
	fenString := fen
	board := parseBoard(strings.Split(fen, " ")[0])
	colorToMove := strings.Split(fen, " ")[1]
	availableCastlings := strings.Split(fen, " ")[2]
	halfMoveClock, _ := strconv.Atoi(strings.Split(fen, " ")[4])
	fullMoveClock, _ := strconv.Atoi(strings.Split(fen, " ")[5])

	return Fen{fenString, board, colorToMove, availableCastlings, halfMoveClock, fullMoveClock}
}

func (f *Fen) AddMove(move string) {
	fromIndex, toIndex := parseMove(move)
	pieceToMove := f.board[fromIndex]
	f.board[fromIndex] = "."
	f.board[toIndex] = pieceToMove

	f.fenString = composeBoardArrayToString(f.board) + " " + f.colorToMove + " " + f.availableCastlings + " - " + strconv.Itoa(f.halfMoveClock) + " " + strconv.Itoa(f.fullMoveClock)
}

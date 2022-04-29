package gofen

import (
	"strconv"
	"strings"
)

func composeBoardArrayToString(board []string) string {
	eightRankString := strings.Join(board[56:], "")
	sevenRankString := strings.Join(board[48:56], "")
	sixRankString := strings.Join(board[40:48], "")
	fiveRankString := strings.Join(board[32:40], "")
	fourRankString := strings.Join(board[24:32], "")
	threeRankString := strings.Join(board[16:24], "")
	twoRankString := strings.Join(board[8:16], "")
	oneRankString := strings.Join(board[0:8], "")

	boardStrings := []string{
		eightRankString,
		sevenRankString,
		sixRankString,
		fiveRankString,
		fourRankString,
		threeRankString,
		twoRankString,
		oneRankString,
	}

	return convertPointsToNumbers(strings.Join(boardStrings, "/"))
}

func convertPointsToNumbers(boardString string) string {
	for i := 8; i >= 2; i-- {
		boardString = strings.Replace(boardString, strings.Repeat(".", i), strconv.Itoa(i), 8)
	}

	return boardString
}

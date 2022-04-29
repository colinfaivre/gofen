package gofen

import (
	"strconv"
	"strings"
)

func parseBoard(board string) []string {
	var boardCells []string
	var rankList = getRanks(board)

	for i := len(rankList) - 1; i >= 0; i-- {
		boardCells = append(boardCells, getRankCells(rankList[i])...)
	}

	return boardCells
}

func getRanks(board string) []string {
	return strings.Split(board, "/")
}

func getRankCells(rank string) []string {
	rank = convertNumbersToPoints(rank)
	return strings.Split(rank, "")
}

func convertNumbersToPoints(rank string) string {
	for i := 1; i <= 8; i++ {
		rank = strings.Replace(rank, strconv.Itoa(i), strings.Repeat(".", i), 8)
	}

	return rank
}

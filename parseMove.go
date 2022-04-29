package gofen

import (
	"strconv"
)

func parseMove(move string) (int, int) {
	from := move[0:2]
	to := move[2:]

	return mapPositionToBoardIndex(from), mapPositionToBoardIndex(to)
}

func mapPositionToBoardIndex(position string) int {
	filesLetter := [8]string{"a", "b", "c", "d", "e", " f", "g", "h"}
	indexMap := map[string]int{}
	counter := 0

	for rank := 1; rank <= 8; rank++ {
		for fileIndex := 0; fileIndex < len(filesLetter); fileIndex++ {
			indexMap[filesLetter[fileIndex]+strconv.Itoa(rank)] = counter
			counter++
		}
	}

	return indexMap[position]
}

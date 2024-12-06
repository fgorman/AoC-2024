package solution4

import (
	"fmt"
	"strings"
)

const TEST = "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

const XMAS = "XMAS"

func Solution(in string) {
	fmt.Println("Running solution for day 4.")

	wordSearch := buildWordSearch(in)

	numXmas := countXmas(wordSearch)
	fmt.Println("Count of XMAS in word search:", numXmas)

	numMas := countMas(wordSearch)
	fmt.Println("Count of X-MAS in word search:", numMas)
}

func buildWordSearch(in string) (out []string) {
	trimmed := strings.TrimSpace(in)
	out = strings.Split(trimmed, "\n")

	return
}

func countXmas(wordSearch []string) (count int) {
	directions := [][]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	for lineIdx, line := range wordSearch {
		for chIdx, ch := range line {
			if ch != 'X' {
				continue
			}

			for _, direction := range directions {
				if searchForXmas(wordSearch, lineIdx, chIdx, 0, direction[0], direction[1]) {
					count++
				}
			}
		}
	}
	return
}

func searchForXmas(wordSearch []string, row, col, index, xDir, yDir int) bool {
	lenSearch := len(wordSearch)
	lenLine := len(wordSearch[0])
	lenXmas := len(XMAS)

	if lenXmas == index {
		return true
	}

	if row >= lenSearch || row < 0 {
		return false
	}

	if col >= lenLine || col < 0 {
		return false
	}

	if wordSearch[row][col] != XMAS[index] {
		return false
	}

	if searchForXmas(wordSearch, row+xDir, col+yDir, index+1, xDir, yDir) {
		return true
	}

	return false
}

func countMas(wordSearch []string) (count int) {
	lenSearch := len(wordSearch)
	lenLine := len(wordSearch[0])

	directions := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	for lineIdx, line := range wordSearch {
		for chIdx, ch := range line {
			if ch != 'A' {
				continue
			}

			found := false

			for _, direction := range directions {

				row := lineIdx + direction[0]
				col := chIdx + direction[1]
				oppRow := lineIdx + -1*direction[0]
				oppCol := chIdx + -1*direction[1]

				if (row < 0 || row >= lenSearch) || (oppRow < 0 || oppRow >= lenSearch) {
					found = false
					break
				}

				if (col < 0 || col >= lenLine) || (oppCol < 0 || oppCol >= lenLine) {
					found = false
					break
				}

				if wordSearch[row][col] != 'M' && wordSearch[row][col] != 'S' {
					found = false
					break
				}

				if wordSearch[row][col] == 'M' && wordSearch[oppRow][oppCol] != 'S' {
					found = false
					break
				} else if wordSearch[row][col] == 'S' && wordSearch[oppRow][oppCol] != 'M' {
					found = false
					break
				}

				found = true
			}

			if found {
				count++
			}
		}
	}
	return
}

package solution6

import (
	"fmt"
	"strings"
)

const TEST = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

type Pos struct {
	Row int
	Col int
}

type PosDir struct {
	Position Pos
	Dir      rune
}

func Solution(in string) {
	fmt.Println("Running solution for day 6.")

	//in = TEST

	in = strings.TrimSpace(in)
	grid := strings.Split(in, "\n")

	startingPos := getGuardStartingPosition(grid)

	visited := getGuardVisited(grid, startingPos)
	fmt.Println("Number of unique visited positions:", len(visited))

	numPlacements := determineNumberOfPositionsForNewObstruction(grid, startingPos, visited)
	fmt.Println("Number of possible obstruction placements:", numPlacements)
}

func replaceAtIndex(in string, idx int, ch rune) string {
	rs := []rune(in)
	rs[idx] = ch
	return string(rs)
}

func getGuardStartingPosition(grid []string) (pos Pos) {
	for row, line := range grid {
		for col, ch := range line {
			if ch == '^' {
				pos.Row = row
				pos.Col = col
				grid[row] = replaceAtIndex(grid[row], col, '.')
				return
			}
		}
	}
	panic("No starting position found")
}

func getGuardVisited(grid []string, start Pos) (visited map[Pos]int) {
	visited = map[Pos]int{}

	turns := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	moves := map[rune][]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}

	numRows := len(grid)
	numCols := len(grid[0])

	currPos := start
	currDir := '^'
	currMove := moves[currDir]

	visited[currPos]++

	for {
		nextPosRow := currPos.Row + currMove[0]
		nextPosCol := currPos.Col + currMove[1]

		if nextPosRow < 0 || nextPosCol < 0 || nextPosRow >= numRows || nextPosCol >= numCols {
			break
		}

		if grid[nextPosRow][nextPosCol] == '#' {
			currDir = turns[currDir]
			currMove = moves[currDir]
			continue
		}

		currPos.Row = nextPosRow
		currPos.Col = nextPosCol

		visited[currPos]++
	}

	return
}

func isCyclic(grid []string, start Pos) bool {
	visited := map[PosDir]int{}

	turns := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	moves := map[rune][]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}

	numRows := len(grid)
	numCols := len(grid[0])

	currPos := start
	currDir := '^'
	currMove := moves[currDir]

	visited[PosDir{currPos, currDir}]++

	for {
		currPosDir := PosDir{currPos, currDir}

		if visited[currPosDir] > 1 {
			return true
		}

		nextPosRow := currPos.Row + currMove[0]
		nextPosCol := currPos.Col + currMove[1]

		if nextPosRow < 0 || nextPosCol < 0 || nextPosRow >= numRows || nextPosCol >= numCols {
			break
		}

		if grid[nextPosRow][nextPosCol] == '#' {
			currDir = turns[currDir]
			currMove = moves[currDir]
			continue
		}

		currPos.Row = nextPosRow
		currPos.Col = nextPosCol

		currPosDir = PosDir{currPos, currDir}

		visited[currPosDir]++
	}

	return false
}

func determineNumberOfPositionsForNewObstruction(grid []string, start Pos, possible map[Pos]int) (num int) {
	for pos := range possible {
		grid[pos.Row] = replaceAtIndex(grid[pos.Row], pos.Col, '#')

		if isCyclic(grid, start) {
			num++
		}

		grid[pos.Row] = replaceAtIndex(grid[pos.Row], pos.Col, '.')
	}

	return
}

package solution10

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Row int
	Col int
}

const TEST = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Solution(in string) {
	topoMap := getTopologicalMap(in)
	trailheads := getTrailheads(topoMap)

	allTrailheadsScore := getScoreForAllTrailheads(topoMap, trailheads)
	fmt.Println("Total score for all trailheads:", allTrailheadsScore)

	allTrailheadsScoreDistinct := getScoreForAllTrailheadsDistinct(topoMap, trailheads)
	fmt.Println("Total score for all distinct trails:", allTrailheadsScoreDistinct)
}

func getTopologicalMap(in string) (out [][]int) {
	out = [][]int{}

	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		nums := []int{}
		numsStr := strings.Split(line, "")
		for _, num := range numsStr {
			if num == "." {
				nums = append(nums, -1)
				continue
			}
			num, _ := strconv.Atoi(num)
			nums = append(nums, num)
		}
		out = append(out, nums)
	}

	return
}

func getTrailheads(topoMap [][]int) (trailheads [][]int) {
	trailheads = [][]int{}

	for rowIdx, row := range topoMap {
		for colIdx, height := range row {
			if height == 0 {
				trailheads = append(trailheads, []int{rowIdx, colIdx})
			}
		}
	}

	return
}

func getScoreForTrailhead(topoMap [][]int, trailhead []int) (score int) {
	q := []Point{}
	visited := map[Point]bool{}

	numRows := len(topoMap)
	numCols := len(topoMap[0])

	q = append(q, Point{trailhead[0], trailhead[1]})

	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]

		currHeight := topoMap[curr.Row][curr.Col]

		if currHeight == 9 && !visited[curr] {
			score++
			visited[curr] = true
			continue
		}

		for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			row := curr.Row + dir[0]
			col := curr.Col + dir[1]

			if row < 0 || row >= numRows || col < 0 || col >= numCols {
				continue
			}

			newHeight := topoMap[row][col]

			if newHeight == currHeight+1 {
				q = append(q, Point{row, col})
			}
		}
	}
	return
}

func getScoreForAllTrailheads(topoMap [][]int, trailheads [][]int) (score int) {
	for _, trailhead := range trailheads {
		score += getScoreForTrailhead(topoMap, trailhead)
	}
	return
}

func getScoreForTrailheadDistinct(topoMap [][]int, trailhead []int) (score int) {
	q := []Point{}

	numRows := len(topoMap)
	numCols := len(topoMap[0])

	q = append(q, Point{trailhead[0], trailhead[1]})

	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]

		currHeight := topoMap[curr.Row][curr.Col]

		if currHeight == 9 {
			score++
			continue
		}

		for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			row := curr.Row + dir[0]
			col := curr.Col + dir[1]

			if row < 0 || row >= numRows || col < 0 || col >= numCols {
				continue
			}

			newHeight := topoMap[row][col]

			if newHeight == currHeight+1 {
				q = append(q, Point{row, col})
			}
		}
	}
	return
}

func getScoreForAllTrailheadsDistinct(topoMap [][]int, trailheads [][]int) (score int) {
	for _, trailhead := range trailheads {
		score += getScoreForTrailheadDistinct(topoMap, trailhead)
	}
	return
}

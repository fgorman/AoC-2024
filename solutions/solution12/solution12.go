package solution12

import (
	"fmt"
	"strings"
)

const TEST = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

type Loc struct {
	Row int
	Col int
}

type Dir Loc

type Area struct {
	Locs      map[Loc]bool
	Area      int
	Perimeter int
	Sides     int
}

func Solution(in string) {
	region := getRegion(in)
	areas := getAllAreas(region)

	costOfAreas := getCostOfAreas(areas)
	fmt.Println("Total cost of fenced in areas:", costOfAreas)

	costOfAreasDiscounted := getCostOfAreasDiscounted(areas)
	fmt.Println("Total cost of fenced in areas with discount:", costOfAreasDiscounted)
}

func getRegion(in string) (out [][]string) {
	out = [][]string{}

	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")

	for _, line := range lines {
		plotChs := strings.Split(line, "")
		out = append(out, plotChs)
	}

	return
}

func getTotalPlot(region [][]string, plotCh string, startRow, startCol int) (area Area) {
	dirsToNumFences := map[int]int{
		0: 4,
		1: 3,
		2: 2,
		3: 1,
		4: 0,
	}

	q := [][]int{}
	reachable := map[Loc]bool{}

	numRows := len(region)
	numCols := len(region[0])

	q = append(q, []int{startRow, startCol})

	plotArea := 0
	plotPerimiter := 0

	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]

		currRow := curr[0]
		currCol := curr[1]

		if reachable[Loc{currRow, currCol}] {
			continue
		}

		reachable[Loc{currRow, currCol}] = true
		plotArea++

		dirs := 0

		for _, dist := range []Dir{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			nextRow := currRow + dist.Row
			nextCol := currCol + dist.Col

			if nextRow < 0 || nextRow >= numRows || nextCol < 0 || nextCol >= numCols {
				continue
			}

			nextCh := region[nextRow][nextCol]

			if nextCh == plotCh {
				q = append(q, []int{nextRow, nextCol})
				dirs++
			}
		}

		plotPerimiter += dirsToNumFences[dirs]
	}

	area.Area = plotArea
	area.Perimeter = plotPerimiter
	area.Locs = reachable
	area.Sides = 0

	for loc := range area.Locs {
		row := loc.Row
		col := loc.Col

		for _, dir := range [][]int{{-1, -1}, {1, -1}, {-1, 1}, {1, 1}} {
			rowDiff := row + dir[0]
			colDiff := col + dir[1]

			if !reachable[Loc{rowDiff, col}] && !reachable[Loc{row, colDiff}] {
				area.Sides++
			}

			if reachable[Loc{rowDiff, col}] && reachable[Loc{row, colDiff}] && !reachable[Loc{rowDiff, colDiff}] {
				area.Sides++
			}
		}
	}

	return
}

func getAllAreas(region [][]string) (areas []Area) {
	areas = []Area{}

	visited := make([][]bool, len(region))
	for idx := range region {
		visited[idx] = make([]bool, len(region[idx]))
	}

	for rowIdx, row := range region {
		for colIdx, plotCh := range row {
			if visited[rowIdx][colIdx] {
				continue
			}

			totalPlot := getTotalPlot(region, plotCh, rowIdx, colIdx)

			for pl := range totalPlot.Locs {
				if !visited[pl.Row][pl.Col] {
					visited[pl.Row][pl.Col] = true
				}
			}

			areas = append(areas, totalPlot)
		}
	}
	return
}

func getCostOfAreas(areas []Area) (cost int) {
	for _, area := range areas {
		cost += area.Area * area.Perimeter
	}
	return
}

func getCostOfAreasDiscounted(areas []Area) (cost int) {
	for _, area := range areas {
		cost += area.Sides * area.Area
	}

	return
}

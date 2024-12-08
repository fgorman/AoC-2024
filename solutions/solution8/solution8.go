package solution8

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-set"
)

const TEST = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

type Location struct {
	Row int
	Col int
}

func Solution(in string) {
	//in = TEST

	numRows, numCols, antennas := getAntennaLocations(in)

	antinodeLocs := getNumDistinctAntinodeLocs(antennas, numRows, numCols, false)
	fmt.Println("Number of disticnt antinode locations:", antinodeLocs)

	antinodeLocs = getNumDistinctAntinodeLocs(antennas, numRows, numCols, true)
	fmt.Println("Number of distinct antinode locations on line:", antinodeLocs)
}

func getAntennaLocations(in string) (numRows, numCols int, locations map[rune][]Location) {
	locations = map[rune][]Location{}

	in = strings.TrimSpace(in)
	rows := strings.Split(in, "\n")
	numRows = len(rows)
	numCols = len(rows[0])
	for rowIdx, row := range rows {
		for colIdx, ch := range []rune(row) {
			if ch == '.' {
				continue
			}

			if locations[ch] == nil {
				locations[ch] = []Location{{rowIdx, colIdx}}
			} else {
				locations[ch] = append(locations[ch], Location{rowIdx, colIdx})
			}
		}
	}
	return
}

func getAllAntennaPairs(antennas []Location) [][]Location {
	pairs := [][]Location{}
	for idx := 0; idx < len(antennas); idx++ {
		for jdx := idx + 1; jdx < len(antennas); jdx++ {
			pairs = append(pairs, []Location{antennas[idx], antennas[jdx]})
		}
	}
	return pairs
}

func getAntinodeLocations(antenna1, antenna2 Location) []Location {
	x1, y1 := antenna1.Row, antenna1.Col
	x2, y2 := antenna2.Row, antenna2.Col

	xdiff := x1 - x2
	ydiff := y1 - y2

	return []Location{
		{x1 + xdiff, y1 + ydiff},
		{x2 - xdiff, y2 - ydiff},
	}
}

func getAntinodeLocationsOnLine(antenna1, antenna2 Location, numRows, numCols int) []Location {
	x1, y1 := antenna1.Row, antenna1.Col
	x2, y2 := antenna2.Row, antenna2.Col

	xdiff := x1 - x2
	ydiff := y1 - y2

	locs := []Location{}
	newX := x1 + xdiff
	newY := y1 + ydiff
	for newX >= 0 && newX < numRows && newY >= 0 && newY < numCols {
		locs = append(locs, Location{newX, newY})

		newX += xdiff
		newY += ydiff
	}

	newX = x2 - xdiff
	newY = y2 - ydiff
	for newX >= 0 && newX < numRows && newY >= 0 && newY < numCols {
		locs = append(locs, Location{newX, newY})

		newX -= xdiff
		newY -= ydiff
	}

	return locs
}

func getNumDistinctAntinodeLocs(antennas map[rune][]Location, numRows, numCols int, onLine bool) int {
	antinodes := set.New[Location](0)

	for _, antenna := range antennas {
		pairs := getAllAntennaPairs(antenna)
		for _, pair := range pairs {
			var newAntinodes []Location
			if onLine {
				antinodes.Insert(pair[0])
				antinodes.Insert(pair[1])
				newAntinodes = getAntinodeLocationsOnLine(pair[0], pair[1], numRows, numCols)
			} else {
				newAntinodes = getAntinodeLocations(pair[0], pair[1])
			}

			for _, antinode := range newAntinodes {
				if antinode.Row >= 0 && antinode.Row < numRows &&
					antinode.Col >= 0 && antinode.Col < numCols {
					antinodes.Insert(antinode)
				}
			}
		}
	}

	return antinodes.Size()
}

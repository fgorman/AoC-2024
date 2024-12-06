package solutions

import (
	"aoc/solutions/solution1"
	"aoc/solutions/solution10"
	"aoc/solutions/solution11"
	"aoc/solutions/solution12"
	"aoc/solutions/solution13"
	"aoc/solutions/solution14"
	"aoc/solutions/solution15"
	"aoc/solutions/solution16"
	"aoc/solutions/solution17"
	"aoc/solutions/solution18"
	"aoc/solutions/solution19"
	"aoc/solutions/solution2"
	"aoc/solutions/solution20"
	"aoc/solutions/solution21"
	"aoc/solutions/solution22"
	"aoc/solutions/solution23"
	"aoc/solutions/solution24"
	"aoc/solutions/solution25"
	"aoc/solutions/solution3"
	"aoc/solutions/solution4"
	"aoc/solutions/solution5"
	"aoc/solutions/solution6"
	"aoc/solutions/solution7"
	"aoc/solutions/solution8"
	"aoc/solutions/solution9"
)

func RunSolution(day int, in string) {
	if day < 1 || day > 25 {
		panic("Day must be within 1 and 25")
	}

	solutionsMap := map[int]func(string){
		1:  solution1.Solution,
		2:  solution2.Solution,
		3:  solution3.Solution,
		4:  solution4.Solution,
		5:  solution5.Solution,
		6:  solution6.Solution,
		7:  solution7.Solution,
		8:  solution8.Solution,
		9:  solution9.Solution,
		10: solution10.Solution,
		11: solution11.Solution,
		12: solution12.Solution,
		13: solution13.Solution,
		14: solution14.Solution,
		15: solution15.Solution,
		16: solution16.Solution,
		17: solution17.Solution,
		18: solution18.Solution,
		19: solution19.Solution,
		20: solution20.Solution,
		21: solution21.Solution,
		22: solution22.Solution,
		23: solution23.Solution,
		24: solution24.Solution,
		25: solution25.Solution,
	}

	solutionsMap[day](in)
}

package solutions

func RunSolution(day int, in string) {
	if day < 1 || day > 25 {
		panic("Day must be within 1 and 25")
	}

	solutions := map[int]func(string){
		1: solution1,
		2: solution2,
		3: solution3,
		4: solution4,
		5: solution5,
		6: solution6,
		7: solution7,
		8: solution8,
		9: solution9,
		10: solution10,
		11: solution11,
		12: solution12,
		13: solution13,
		14: solution14,
		15: solution15,
		16: solution16,
		17: solution17,
		18: solution18,
		19: solution19,
		20: solution20,
		21: solution21,
		22: solution22,
		23: solution23,
		24: solution24,
		25: solution25,
		
	}

	solutions[day](in)
}
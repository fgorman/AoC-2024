package solution7

import (
	"fmt"
	"strconv"
	"strings"
)

const TEST = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

type Equation struct {
	Res int
	Ops []int
}

func Solution(in string) {
	fmt.Println("Running solution for day 7.")

	//in = TEST

	equations := getEquations(in)

	totalCalibRes := getTotalCalibrationResult(equations, []rune{'+', '*'})
	fmt.Println("Total calibration result:", totalCalibRes)

	totalCalibResWithConcat := getTotalCalibrationResult(equations, []rune{'+', '*', '|'})
	fmt.Println("Total calibration result with concat:", totalCalibResWithConcat)
}

func getEquations(in string) (out []Equation) {
	out = []Equation{}

	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")

	for _, line := range lines {
		resAndNums := strings.Split(line, ": ")
		res, _ := strconv.Atoi(resAndNums[0])
		numsStr := strings.Split(resAndNums[1], " ")
		nums := []int{}
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		out = append(out, Equation{res, nums})
	}
	return
}

// Taken from https://prtamil.github.io/posts/cartesian-product-go/
func product(input []rune, n int) [][]rune {
	//Make atlease single array else won't go into loop
	prod := make([][]rune, 1)
	for i := 1; i <= n; i++ {
		//next Array is stores intermediate results
		next := make([][]rune, 0)
		for _, x := range prod {
			for _, y := range input {
				//t = [x+[y]]
				//x = [1]
				//y = 2
				// t = [1,2]
				t := make([]rune, 0)
				t = append(t, x...)
				t = append(t, y)
				//append to next 2d array
				next = append(next, [][]rune{t}...)
			}
		}
		//Assign intermediate next to prod so next loop will have new items
		//
		prod = next
	}
	return prod
}

func doOperation(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	case '|':
		res, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
		return res
	default:
		return a
	}
}

func isCorrectOperators(expected int, operands []int, operators []rune) bool {
	testRes := operands[0]
	for idx := 1; idx < len(operands); idx++ {
		testRes = doOperation(testRes, operands[idx], operators[idx-1])
	}

	return testRes == expected
}

func getTotalCalibrationResult(equations []Equation, operators []rune) (sum int) {
	for _, equation := range equations {
		combs := product(operators, len(equation.Ops)-1)
		for _, comb := range combs {
			if isCorrectOperators(equation.Res, equation.Ops, comb) {
				sum += equation.Res
				break
			}
		}
	}
	return
}

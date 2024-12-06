package solution3

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const TEST = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func Solution(in string) {
	fmt.Println("Running solution for day 3.")

	mulSum := getMulSum(in)
	fmt.Println("Sum of mul operations:", mulSum)

	doMulSum := getDoMulSum(in)
	fmt.Println("Sum of mul operations in do's:", doMulSum)
}

func parseMulOperation(op string) int {
	numsStr := op[4 : len(op)-1]
	numsSplt := strings.Split(numsStr, ",")

	v1, _ := strconv.Atoi(numsSplt[0])
	v2, _ := strconv.Atoi(numsSplt[1])

	return v1 * v2
}

func getMulSum(in string) int {
	var sum int = 0

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := re.FindAllString(in, -1)

	for _, match := range matches {
		sum += parseMulOperation(match)
	}

	return sum
}

func getDoMulSum(in string) int {
	const DONT_LEN = len("don't()")
	const DO_LEN = len("do()")

	var sum int = 0

	var enabled bool = true

	for idx, ch := range in {
		if ch == 'd' {
			dontEnd := idx + DONT_LEN
			doEnd := idx + DO_LEN
			if dontEnd < len(in) && in[idx:dontEnd] == "don't()" {
				enabled = false
			} else if doEnd < len(in) && in[idx:doEnd] == "do()" {
				enabled = true
			}
		} else if ch == 'm' {
			re := regexp.MustCompile(`mul\(\d+,\d+\)`)
			match := re.FindStringIndex(in[idx:])

			if match[0]+idx == idx && enabled {
				sum += parseMulOperation(in[match[0]+idx : match[1]+idx])
			}
		}
	}

	return sum
}

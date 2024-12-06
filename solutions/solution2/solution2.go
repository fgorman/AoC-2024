package solution2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const TEST = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

func Solution(in string) {
	fmt.Println("Runnig solution for day 2.")

	reports := getReports(in)

	numSafe := determineNumSafe(reports)
	fmt.Println("Number of safe reports:", numSafe)

	numSafewithDampener := determineNumSafeWithDampener(reports)
	fmt.Println("Number of safe reports (after dampener):", numSafewithDampener)
}

func getReports(in string) [][]int {
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")

	reports := make([][]int, len(lines))

	for reportIdx, report := range lines {
		nums := strings.Split(report, " ")
		reportNums := make([]int, len(nums))
		for numIdx, num := range nums {
			n, _ := strconv.Atoi(num)
			reportNums[numIdx] = n
		}
		reports[reportIdx] = reportNums
	}

	return reports
}

func isSafe(report []int) bool {
	var firstCmp bool = report[0] < report[1]
	firstDiff := math.Abs(float64(report[0] - report[1]))
	if firstDiff < 1 || firstDiff > 3 {
		return false
	}

	for num := 1; num < len(report)-1; num++ {
		cmp := report[num] < report[num+1]
		diff := math.Abs(float64(report[num] - report[num+1]))
		if (cmp != firstCmp) || (diff < 1 || diff > 3) {
			return false
		}
	}

	return true
}

func determineNumSafe(reports [][]int) int {
	var numSafe int = 0

	for reportNum := 0; reportNum < len(reports); reportNum++ {
		if isSafe(reports[reportNum]) {
			numSafe++
		}
	}

	return numSafe
}

func determineNumSafeWithDampener(reports [][]int) int {
	var numSafe int = 0

	for reportNum := 0; reportNum < len(reports); reportNum++ {
		report := reports[reportNum]
		if isSafe(reports[reportNum]) {
			numSafe++
		} else {
			for idx := range report {
				reportCopy := make([]int, len(report))
				copy(reportCopy, report)
				newSlice := append(reportCopy[:idx], reportCopy[idx+1:]...)
				if isSafe(newSlice) {
					numSafe++
					break
				}
			}
		}
	}

	return numSafe
}

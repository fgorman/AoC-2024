package solutions

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func solution1(in string) {
	fmt.Println("Running solution for day 1.")
	
	left, right := parseInput(in)

	part1(left, right)

	part2(left, right)
}

func parseInput(in string) (left, right []int) {
	in = strings.TrimSpace(in)
	lines := strings.Split(in, "\n")
	left = make([]int, len(lines))
	right = make([]int, len(lines))

	for i, line := range lines {
		nums := strings.Split(line, "   ")

		if len(nums) < 2 {
			continue
		}

		convL, _ := strconv.Atoi(nums[0])
		convR, _ := strconv.Atoi(nums[1])

		left[i] = convL
		right[i] = convR
	}

	sort.Ints(left)
	sort.Ints(right)

	return
}

func part1(left, right []int) {
	var totalDistance int = 0
	for idx := 0; idx < len(left); idx++ {
		totalDistance += int(math.Abs(float64(left[idx] - right[idx])))
	}

	fmt.Println("Total distance:", totalDistance)
}

func part2(left, right []int) {
	numMap := make(map[int]int)

	for _, num := range left {
		numMap[num]++
	}

	var similarityScore int = 0
	for _, num := range right {
		similarityScore += num * numMap[num]
	}

	fmt.Println("Similarity score:", similarityScore)
}
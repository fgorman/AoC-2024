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

	left, right := getLeftAndRight(in)

	totalDistance := getTotalDistance(left, right)
	fmt.Println("Total distance:", totalDistance)

	similarityScore := getSimilarityScore(left, right)
	fmt.Println("Similarity score:", similarityScore)
}

func getLeftAndRight(in string) (left, right []int) {
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

func getTotalDistance(left, right []int) int {
	var totalDistance int = 0
	for idx := 0; idx < len(left); idx++ {
		totalDistance += int(math.Abs(float64(left[idx] - right[idx])))
	}

	return totalDistance
}

func getSimilarityScore(left, right []int) int {
	numMap := make(map[int]int)

	for _, num := range left {
		numMap[num]++
	}

	var similarityScore int = 0
	for _, num := range right {
		similarityScore += num * numMap[num]
	}

	return similarityScore
}

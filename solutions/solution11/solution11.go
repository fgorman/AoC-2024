package solution11

import (
	"fmt"
	"strconv"
	"strings"
)

const TEST = "125 17"

func Solution(in string) {
	stones := getInitialStones(in)

	numStones := numStonesAfterChanges(stones, 25)
	fmt.Println("Number of stones after 25 changes:", numStones)

	numStones = numStonesAfterChanges(stones, 75)
	fmt.Println("Number of stones after 75 changes:", numStones)
}

func getNumDigits(num int) int {
	if num == 0 {
		return 1
	}

	var count int
	for num != 0 {
		num /= 10
		count++
	}

	return count
}

func getInitialStones(in string) (stones map[int]int) {
	stones = map[int]int{}

	in = strings.TrimSpace(in)
	numsStrs := strings.Split(in, " ")
	for _, numStr := range numsStrs {
		num, _ := strconv.Atoi(numStr)
		stones[num]++
	}
	return
}

func changeStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	numDigits := getNumDigits(stone)
	if numDigits%2 == 0 {
		mid := numDigits / 2
		numStr := strconv.Itoa(stone)
		chs := []rune(numStr)
		left, _ := strconv.Atoi(string(chs[:mid]))
		right, _ := strconv.Atoi(string(chs[mid:]))
		return []int{left, right}
	}

	return []int{stone * 2024}
}

func numStonesAfterChanges(stones map[int]int, numSplits int) (numStones int) {
	for numSplits > 0 {
		newStones := map[int]int{}
		for stoneNum, count := range stones {
			newNums := changeStone(stoneNum)
			for _, newNum := range newNums {
				newStones[newNum] += count
			}
		}

		stones = newStones

		numSplits--
	}

	for _, count := range stones {
		numStones += count
	}
	return
}

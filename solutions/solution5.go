package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set"
)

const TEST5 = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func solution5(in string) {
	fmt.Println("Running solution for day 5.")

	//in = TEST5

	in = strings.TrimSpace(in)
	splt := strings.Split(in, "\n\n")

	pageRules := getPageRules(splt[0])
	updates := getUpdates(splt[1])

	midSum := midSumFromCorrect(pageRules, updates)
	fmt.Println("Sum of middle elements of correct updates:", midSum)

	midSumFixed := midSumFromFixed(pageRules, updates)
	fmt.Println("Sum of middle elements of fixed updates:", midSumFixed)

}

func getPageRules(str string) (rules map[int]*set.Set[int]) {
	rules = make(map[int]*set.Set[int])

	ruleStrs := strings.Split(str, "\n")

	for _, ruleStr := range ruleStrs {
		pages := strings.Split(ruleStr, "|")
		p1, _ := strconv.Atoi(pages[0])
		p2, _ := strconv.Atoi(pages[1])

		if rules[p1] == nil {
			rules[p1] = set.New[int](0)
		}

		rules[p1].Insert(p2)
	}

	return
}

func getUpdates(str string) (updates [][]int) {
	updateStrs := strings.Split(str, "\n")

	for _, updateStr := range updateStrs {
		var update []int
		numStrs := strings.Split(updateStr, ",")
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			update = append(update, num)
		}
		updates = append(updates, update)
	}
	return
}

func isCorrectUpdate(rules map[int]*set.Set[int], update []int) bool {
	for idx, num := range update {
		for jdx := idx + 1; jdx < len(update); jdx++ {
			if rules[num] == nil {
				return false
			}

			if !rules[num].Contains(update[jdx]) {
				return false
			}
		}
	}

	return true
}

func midSumFromCorrect(rules map[int]*set.Set[int], updates [][]int) (sum int) {
	for _, update := range updates {
		if isCorrectUpdate(rules, update) {
			sum += update[len(update)/2]
		}
	}
	return
}

func fixIncorrectUpdate(rules map[int]*set.Set[int], update []int) []int {
	slices.SortFunc(update, func(a, b int) int {
		if rules[a] == nil {
			return 1
		}

		if rules[a].Contains(b) {
			return -1
		}

		return 1
	})

	return update
}

func midSumFromFixed(rules map[int]*set.Set[int], updates [][]int) (sum int) {
	for _, update := range updates {
		if isCorrectUpdate(rules, update) {
			continue
		}

		corrected := fixIncorrectUpdate(rules, update)

		sum += corrected[len(corrected)/2]
	}
	return
}

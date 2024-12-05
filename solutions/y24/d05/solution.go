package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"slices"
)

const YEAR = "24"
const DAY = "05"

type rule struct {
	x, y int
}

func parse(input []string) ([]rule, [][]int) {
	var rules []rule
	var updates [][]int

	first := true
	for _, row := range input {
		if row == "" {
			first = false
			continue
		}
		nums := util.GetIntsFromString(row)
		if first {
			rules = append(rules, rule{nums[0], nums[1]})
		} else {
			updates = append(updates, nums)
		}
	}
	return rules, updates
}

func part1(path string) {
	rules, updates := parse(util.ToStringSlice(path))

	score := 0
	for _, pages := range updates {
		correct := true
	out:
		for i, page := range pages {
			for _, rule := range rules {
				if page == rule.x {
					if slices.Contains(pages[:i], rule.y) {
						correct = false
						break out
					}
				}
			}
		}
		if correct {
			score += pages[len(pages)/2]
		}
	}

	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	rules, updates := parse(util.ToStringSlice(path))

	score := 0
	for p, pages := range updates {
		fixed := false
		for i := 0; i < len(pages); i++ {
		out:
			for _, rule := range rules {
				if updates[p][i] == rule.y {
					if slices.Contains(updates[p][i:], rule.x) {
						updates[p] = append(util.RemoveElement(updates[p], i), rule.y)
						i--
						fixed = true
						break out
					}
				}
			}
		}
		if fixed {
			score += updates[p][len(pages)/2]
		}
	}

	fmt.Printf("part2: %d\n", score)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

package main

import (
	"fmt"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "04"

func part1(path string) {
	input := util.ToStringSlice(path)
	score := 0
	for _, r := range input {
		nums := util.GetIntsFromString(r)
		x1, y1, x2, y2 := nums[0], nums[1], nums[2], nums[3]
		if x1 <= x2 && y1 >= y2 || x2 <= x1 && y2 >= y1 {
			score += 1
		}
	}
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToStringSlice(path)
	score := 0
	for _, r := range input {
		nums := util.GetIntsFromString(r)
		x1, y1, x2, y2 := nums[0], nums[1], nums[2], nums[3]
		if max(y1, y2)-min(x1, x2) <= (y1-x1)-(x2-y2) {
			score += 1
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

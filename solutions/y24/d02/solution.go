package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "02"

func part1(path string) {
	input := util.ToStringSlice(path)
	score := 0
	for _, r := range input {
		nums := util.GetIntsFromString(r)
		if safe(nums) {
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
		for i := 0; i < len(nums); i++ {
			if safe(util.RemoveElement(nums, i)) {
				score += 1
				break
			}
		}
	}
	fmt.Printf("part2: %d\n", score)
}

func safe(nums []int) bool {
	increase := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		diff := util.Abs(nums[i] - nums[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
		if increase && nums[i] < nums[i-1] {
			return false
		}
		if !increase && nums[i] > nums[i-1] {
			return false
		}
	}
	return true
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

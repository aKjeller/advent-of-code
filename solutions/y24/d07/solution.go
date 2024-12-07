package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "07"

func part1(path string) {
	score := solve(util.ToStringSlice(path), []uint8{'+', '*'})
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	score := solve(util.ToStringSlice(path), []uint8{'+', '*', '|'})
	fmt.Printf("part2: %d\n", score)
}

func solve(input []string, ops []uint8) int {
	score := 0
	for _, row := range input {
		nums := util.GetIntsFromString(row)
		if permuteWrap(nums[1:], ops, nums[0]) {
			score += nums[0]
		}
	}
	return score
}

func permuteWrap(nums []int, ops []uint8, ans int) bool {
	return permute(nums, ops, ans, -1, 0, 0)
}

func permute(nums []int, ops []uint8, ans, i int, op uint8, count int) bool {
	if i == len(nums) {
		return count == ans
	}

	if count > ans {
		return false
	}

	if op == '+' {
		count += nums[i]
	} else if op == '*' {
		count *= nums[i]
	} else if op == '|' {
		count = util.Concatenate(count, nums[i])
	}

	for _, o := range ops {
		if permute(nums, ops, ans, i+1, o, count) {
			return true
		}
	}

	return false
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

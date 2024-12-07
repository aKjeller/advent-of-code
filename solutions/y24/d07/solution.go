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

func solve(input []string, types []uint8) int {
	score := 0
	for _, row := range input {
		nums := util.GetIntsFromString(row)

		total := nums[0]
		nums = nums[1:]

		var ops []string
		permute(types, len(nums)-1, "", &ops)
		for _, op := range ops {
			sum := calc(op, nums, total)
			if sum == total {
				score += total
				break
			}
		}
	}
	return score
}

func calc(ops string, nums []int, max int) int {
	t := nums[0]
	for i, _ := range ops {
		if ops[i] == '+' {
			t += nums[i+1]
		} else if ops[i] == '*' {
			t *= nums[i+1]
		} else {
			// int -> string -> int = 3.0s
			//t = util.ParseInt(fmt.Sprintf("%d%d", t, nums[i+1]))

			// concatenate = 1.1s
			t = util.Concatenate(t, nums[i+1])
		}
		if t > max {
			return t
		}
	}
	return t
}

func permute(types []uint8, i int, op string, ops *[]string) {
	if i == 0 {
		*ops = append(*ops, op)
		return
	}
	for _, t := range types {
		permute(types, i-1, op+string(t), ops)
	}
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

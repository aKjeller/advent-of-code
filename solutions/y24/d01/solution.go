package main

import (
	"fmt"
	"sort"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "01"

func part1(path string) {
	input := util.ToStringSlice(path)

	var left []int
	var right []int
	for _, r := range input {
		nums := util.GetIntsFromString(r)
		left = append(left, nums[0])
		right = append(right, nums[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	score := 0
	for i := 0; i < len(left); i++ {
		score += util.Abs(left[i] - right[i])
	}

	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToStringSlice(path)

	var left []int
	m := make(map[int]int)
	for _, r := range input {
		nums := util.GetIntsFromString(r)
		left = append(left, nums[0])
		m[nums[1]]++
	}

	score := 0
	for _, i := range left {
		score += i * m[i]
	}

	fmt.Printf("part2: %d\n", score)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

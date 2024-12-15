package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"math"
	"strings"
)

const YEAR = "24"
const DAY = "13"

func part1(path string) {
	input := util.ToString(path)
	score := 0
	for _, e := range strings.Split(input, "\n\n") {
		score += solve(util.GetFloatsFromString(e))
	}
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToString(path)
	score := 0
	for _, e := range strings.Split(input, "\n\n") {
		nums := util.GetFloatsFromString(e)
		nums[4], nums[5] = nums[4]+10000000000000, nums[5]+10000000000000
		score += solve(nums)
	}
	fmt.Printf("part2: %d\n", score)
}

func solve(nums []float64) int {
	a1, a2, b1, b2, tot1, tot2 := nums[0], nums[1], nums[2], nums[3], nums[4], nums[5]
	det := (a1 * b2) - (b1 * a2)
	a := math.Round(b2/det*tot1 + -b1/det*tot2)
	b := math.Round(-a2/det*tot1 + a1/det*tot2)

	tickets := 0
	if a*a1+b*b1 == tot1 && a*a2+b*b2 == tot2 {
		tickets += int(a)*3 + int(b)
	}

	return tickets
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

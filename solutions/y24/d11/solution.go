package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"strconv"
)

const YEAR = "24"
const DAY = "11"

func part1(path string) {
	nums := util.GetIntsFromString(util.ToString(path))
	for range 25 {
		var next []int
		for _, n := range nums {
			a, b := eval(n)
			next = append(next, a)
			if b > -1 {
				next = append(next, b)
			}
		}
		nums = next
	}
	fmt.Printf("part1: %d\n", len(nums))
}

func part2(path string) {
	nums := util.GetIntsFromString(util.ToString(path))
	score := 0
	cache := make(map[[2]int]int)
	for _, n := range nums {
		score += recurse(n, 0, cache)
	}
	fmt.Printf("part2: %d\n", score)
}

func eval(n int) (int, int) {
	if n == 0 {
		return 1, -1
	}
	if len(strconv.Itoa(n))%2 == 0 {
		s := strconv.Itoa(n)
		return util.ParseInt(s[:len(s)/2]), util.ParseInt(s[len(s)/2:])
	}
	return n * 2024, -1
}

func recurse(n, depth int, cache map[[2]int]int) int {
	if depth == 75 {
		return 1
	}
	if v, ok := cache[[2]int{n, depth}]; ok {
		return v
	}

	a, b := eval(n)
	av := recurse(a, depth+1, cache)
	cache[[2]int{a, depth + 1}] = av
	if b > -1 {
		bv := recurse(b, depth+1, cache)
		cache[[2]int{b, depth + 1}] = bv
		return av + bv
	}
	return av
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

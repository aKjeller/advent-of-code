package main

import (
	"fmt"
	"sort"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "01"

func part1(path string) {
	input := util.ToIntSlice(path)

	t := 0
	c := 0
	for _, v := range input {
		if v == nil {
			t = max(t, c)
			c = 0
		} else {
			c += *v
		}
	}

	fmt.Println("part1: ", t)
}

func part2(path string) {
	input := util.ToIntSlice(path)

	var elfs []int
	c := 0
	for _, v := range input {
		if v == nil {
			elfs = append(elfs, c)
			c = 0
		} else {
			c += *v
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))
	res := elfs[0] + elfs[1] + elfs[2]
	fmt.Println("part2: ", res)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

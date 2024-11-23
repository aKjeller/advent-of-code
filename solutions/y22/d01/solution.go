package main

import (
	"fmt"
	"sort"
	"strconv"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "01"

func part1(input []string) {
	t := 0
	c := 0
	for _, e := range input {
		if e == "" {
			t = max(t, c)
			c = 0
		} else {
			v, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			c += v
		}
	}

	fmt.Println("part1: ", t)
}

func part2(input []string) {
	var elfs []int

	c := 0
	for _, e := range input {
		if e == "" {
			elfs = append(elfs, c)
			c = 0
		} else {
			v, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			c += v
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))
	res := elfs[0] + elfs[1] + elfs[2]
	fmt.Println("part2: ", res)
}

func main() {
	part1(util.ToStringSlice(util.ExamplePath(YEAR, DAY)))
	part1(util.ToStringSlice(util.InputPath(YEAR, DAY)))

	part2(util.ToStringSlice(util.ExamplePath(YEAR, DAY)))
	part2(util.ToStringSlice(util.InputPath(YEAR, DAY)))
}

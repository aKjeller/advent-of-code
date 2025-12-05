package main

import (
	"fmt"
	"sort"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "05"

type interval struct {
	start, end int
}

func part1(path string) {
	input := util.ToStringSlice(path)

	var ranges []interval

	var startId int
	for i, r := range input {
		if r == "" {
			startId = i + 1
			break
		}
		parts := strings.Split(r, "-")
		ranges = append(ranges, interval{util.ParseInt(parts[0]), util.ParseInt(parts[1])})
	}

	sum := 0
	for i := startId; i < len(input); i++ {
		for _, r := range ranges {
			id := util.ParseInt(input[i])
			if id >= r.start && id <= r.end {
				sum += 1
				break
			}
		}
	}

	fmt.Println("part1: ", sum)
}

type item struct {
	id      int
	isStart bool
}

func part2(path string) {
	input := util.ToStringSlice(path)

	var items []item
	for _, r := range input {
		if r == "" {
			break
		}
		parts := strings.Split(r, "-")
		items = append(items, item{util.ParseInt(parts[0]), true})
		items = append(items, item{util.ParseInt(parts[1]), false})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].id != items[j].id {
			return items[i].id < items[j].id
		}
		return items[i].isStart && !items[j].isStart
	})

	var sum, c, prevC, i int
	for _, item := range items {
		if item.isStart {
			c += 1
		} else {
			c -= 1
		}

		if c > 0 && prevC == 0 {
			i = item.id
		}
		if c == 0 && prevC > 0 {
			sum += item.id - i + 1
		}

		prevC = c
	}

	fmt.Println("part2: ", sum)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

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

func part2(path string) {
	input := util.ToStringSlice(path)

	var ranges []interval
	for _, r := range input {
		if r == "" {
			break
		}
		parts := strings.Split(r, "-")
		ranges = append(ranges, interval{util.ParseInt(parts[0]), util.ParseInt(parts[1])})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	sum, start, end := 0, ranges[0].start, ranges[0].end
	for _, item := range ranges {
		if item.start > end {
			sum += end - start + 1
			start = item.start
		}
		end = max(end, item.end)
	}
	sum += end - start + 1

	fmt.Println("part2: ", sum)
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

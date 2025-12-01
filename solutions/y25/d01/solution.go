package main

import (
	"fmt"
	"strconv"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "01"

func part1(path string) {
	input := util.ToStringSlice(path)
	pw := 0

	dial := 50
	for _, line := range input {
		step := 1
		if line[0] == 'L' {
			step = -1
		}
		n, _ := strconv.Atoi(line[1:])
		for range n {
			dial += step
			if dial == -1 {
				dial = 99
			}
			if dial == 100 {
				dial = 0
			}
		}
		if dial == 0 {
			pw += 1
		}
	}

	fmt.Println("part1: ", pw)
}

func part2(path string) {
	input := util.ToStringSlice(path)
	pw := 0

	dial := 50
	for _, line := range input {
		step := 1
		if line[0] == 'L' {
			step = -1
		}
		n, _ := strconv.Atoi(line[1:])

		for range n {
			dial += step
			if dial == -1 {
				dial = 99
			}
			if dial == 100 {
				dial = 0
			}
			if dial == 0 {
				pw += 1
			}
		}
	}

	fmt.Println("part2: ", pw)
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

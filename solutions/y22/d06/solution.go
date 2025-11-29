package main

import (
	"fmt"
	"slices"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "06"

func part1(path string) {
	seq := util.ToString(path)
	score := findStart(seq, 4)
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	seq := util.ToString(path)
	score := findStart(seq, 14)
	fmt.Printf("part2: %d\n", score)
}

func findStart(seq string, size int) int {
	b := make([]rune, size)
	bi := 0

	for i, s := range seq {
		b[bi] = s
		if i >= size-1 && allDifferent(b) {
			return i + 1
		}
		if bi == size-1 {
			bi = 0
		} else {
			bi++
		}
	}
	return -1
}

func allDifferent(b []rune) bool {
	for i, elem := range b {
		if slices.Contains(b[i+1:], elem) {
			return false
		}
	}
	return true
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

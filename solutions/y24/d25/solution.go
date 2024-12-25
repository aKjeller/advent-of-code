package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"strings"
)

const YEAR = "24"
const DAY = "25"

func part1(path string) {
	input := strings.Split(util.ToString(path), "\n\n")

	score := 0
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if matches(input[i], input[j]) {
				score += 1
			}
		}
	}

	fmt.Printf("part1: %d\n", score)
}

func matches(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == '#' && b[i] == '#' {
			return false
		}
	}
	return true
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
}

package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"strings"
)

const YEAR = "24"
const DAY = "19"

func part1(path string) {
	input := util.ToStringSlice(path)
	towels := strings.Split(input[0], ", ")
	score := 0
	cache := make(map[string]int)
	for _, p := range input[2:] {
		if findPatterns(p, towels, cache) > 0 {
			score++
		}
	}
	fmt.Printf("part1: %d\n", score)
}

func findPatterns(desired string, towels []string, cache map[string]int) int {
	if v, ok := cache[desired]; ok {
		return v
	}

	if desired == "" {
		return 1
	}

	pos := 0
	for _, towel := range towels {
		if strings.HasPrefix(desired, towel) {
			pos += findPatterns(strings.TrimPrefix(desired, towel), towels, cache)
		}
	}

	cache[desired] = pos
	return pos
}

func part2(path string) {
	input := util.ToStringSlice(path)
	towels := strings.Split(input[0], ", ")
	score := 0
	cache := make(map[string]int)
	for _, p := range input[2:] {
		score += findPatterns(p, towels, cache)
	}
	fmt.Printf("part2: %d\n", score)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

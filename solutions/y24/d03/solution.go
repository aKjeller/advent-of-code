package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"regexp"
	"strings"
)

const YEAR = "24"
const DAY = "03"

func part1(path string) {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(util.ToString(path), -1)

	score := 0
	for _, m := range matches {
		score += util.ParseInt(m[1]) * util.ParseInt(m[2])
	}

	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := r.FindAllStringSubmatch(util.ToString(path), -1)

	score := 0
	enabled := true
	for _, m := range matches {
		if strings.HasPrefix(m[0], "mul") && enabled {
			score += util.ParseInt(m[1]) * util.ParseInt(m[2])
		} else if m[0] == "do()" {
			enabled = true
		} else if m[0] == "don't()" {
			enabled = false
		}
	}

	fmt.Printf("part2: %d\n", score)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

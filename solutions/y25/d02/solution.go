package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "02"

func part1(path string) {
	input := strings.ReplaceAll(util.ToString(path), "\n", "")
	ranges := strings.Split(input, ",")

	score := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, end := util.ParseInt(parts[0]), util.ParseInt(parts[1])

		for i := start; i <= end; i++ {
			b := strconv.Itoa(i)
			left, right := b[:len(b)/2], b[len(b)/2:]
			if left == right {
				score += i
			}
		}
	}

	fmt.Println("part1: ", score)
}

func part2(path string) {
	input := strings.ReplaceAll(util.ToString(path), "\n", "")
	ranges := strings.Split(input, ",")

	score := 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, end := util.ParseInt(parts[0]), util.ParseInt(parts[1])

		for i := start; i <= end; i++ {
			b := strconv.Itoa(i)
			for c := 1; c <= len(b)/2; c++ {
				equal := true

				first := []byte(b[:c])
				for chunk := range slices.Chunk([]byte(b), c) {
					if !slices.Equal(first, chunk) {
						equal = false
						break
					}
				}

				if equal {
					score += i
					break
				}
			}
		}
	}

	fmt.Println("part2: ", score)
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

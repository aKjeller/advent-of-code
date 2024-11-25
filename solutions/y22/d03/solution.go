package main

import (
	"fmt"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "03"

func part1(path string) {
	input := util.ToStringSlice(path)
	score := 0
	for _, r := range input {
		a := r[0 : len(r)/2]
		b := r[len(r)/2:]

	out:
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				if a[i] == b[j] {
					score += int(getPriority(a[i]))
					break out
				}
			}
		}
	}

	fmt.Printf("part1: %d\n", score)
}

func getPriority(c byte) byte {
	if c > 96 {
		return c - 96
	}
	return c - 65 + 27
}

func part2(path string) {
	input := util.ToStringSlice(path)
	score := 0

	for n := 0; n < len(input); n = n + 3 {
		a, b, c := input[n], input[n+1], input[n+2]
	out:
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(b); j++ {
				for k := 0; k < len(c); k++ {
					if a[i] == b[j] && a[i] == c[k] {
						score += int(getPriority(a[i]))
						break out
					}
				}
			}
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

package main

import (
	"fmt"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "12"

func part1(path string) {
	input := strings.Split(util.ToString(path), "\n\n")

	var presents []int
	var nextIndex int
	for i, line := range input {
		if strings.Contains(line, "x") {
			nextIndex = i
			break
		}
		presents = append(presents, strings.Count(line, "#"))
	}

	part1 := 0
	for tree := range strings.SplitSeq(strings.TrimSpace(input[nextIndex]), "\n") {
		parts := strings.Split(tree, ":")

		nPresents := util.GetIntsFromString(parts[1])

		area := 0
		for i := range presents {
			area += nPresents[i] * presents[i]
		}

		dim := strings.Split(parts[0], "x")
		if area <= util.ParseInt(dim[0])*util.ParseInt(dim[1]) {
			part1 += 1
		}
	}

	fmt.Println("part1: ", part1)
}

func main() {
	part1(util.InputPath(YEAR, DAY))
}

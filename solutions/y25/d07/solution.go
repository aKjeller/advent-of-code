package main

import (
	"fmt"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "07"

func part1(path string) {
	input := util.ToGrid(path)
	splits := 0
	for ri, row := range input {
		for ci := range row {
			if input[ri][ci] == 'S' {
				input[ri+1][ci] = '|'
			} else if input[ri][ci] == '^' && input[ri-1][ci] == '|' {
				input[ri][ci-1] = '|'
				input[ri][ci+1] = '|'
				splits += 1
			} else if ri > 0 && input[ri][ci] == '.' && input[ri-1][ci] == '|' {
				input[ri][ci] = '|'
			}
		}
	}
	fmt.Println("part1: ", splits)
}

func part2(path string) {
	input := util.ToGrid(path)

	var timelines [][]int
	for _, row := range input {
		timelines = append(timelines, make([]int, len(row)))
	}

	for ri, row := range input {
		for ci := range row {
			if input[ri][ci] == 'S' {
				timelines[ri+1][ci] = 1
			} else if input[ri][ci] == '^' {
				timelines[ri][ci-1] += timelines[ri-1][ci]
				timelines[ri][ci+1] += timelines[ri-1][ci]
			} else if ri > 0 && input[ri][ci] == '.' {
				timelines[ri][ci] += timelines[ri-1][ci]
			}
		}
	}

	sum := 0
	for _, timeline := range timelines[len(timelines)-1] {
		sum += timeline
	}

	fmt.Println("part2: ", sum)
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "10"

func part1(path string) {
	m := util.ToGrid(path)

	score := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '0' {
				trails := make(map[[2]int]int)
				travel(m, i, j, '0', trails)
				score += len(trails)
			}
		}
	}

	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	m := util.ToGrid(path)

	trails := make(map[[2]int]int)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '0' {
				travel(m, i, j, '0', trails)
			}
		}
	}

	score := 0
	for _, v := range trails {
		score += v
	}

	fmt.Printf("part2: %d\n", score)
}

func travel(m [][]uint8, x, y int, level uint8, trails map[[2]int]int) {
	if level == '9' {
		trails[[2]int{x, y}]++
		return
	}

	walk(m, x+1, y, level+1, trails)
	walk(m, x-1, y, level+1, trails)
	walk(m, x, y+1, level+1, trails)
	walk(m, x, y-1, level+1, trails)
}

func walk(m [][]uint8, x, y int, level uint8, trails map[[2]int]int) {
	if x >= 0 && x < len(m) && y >= 0 && y < len(m[0]) {
		if m[x][y] == level {
			travel(m, x, y, level, trails)
		}
	}
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

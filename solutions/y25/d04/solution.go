package main

import (
	"fmt"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "04"

type direction [2]int

var (
	north     direction = [2]int{-1, 0}
	south     direction = [2]int{1, 0}
	west      direction = [2]int{0, -1}
	east      direction = [2]int{0, 1}
	northEast direction = [2]int{-1, 1}
	southEast direction = [2]int{1, 1}
	northWest direction = [2]int{-1, -1}
	southWest direction = [2]int{1, -1}
)

var directions = [8]direction{north, south, west, east, northEast, southEast, northWest, southWest}

func part1(path string) {
	m := util.ToGrid(path)

	sum := 0
	for i := range m {
		for j := range len(m[i]) {
			if isFree(m, i, j) {
				sum += 1
			}
		}
	}

	fmt.Println("part1: ", sum)
}

func part2(path string) {
	m := util.ToGrid(path)

	oldSum := -1
	sum := 0
	for oldSum != sum {
		oldSum = sum
		for i := range m {
			for j := range len(m[i]) {
				if isFree(m, i, j) {
					m[i][j] = '.'
					sum += 1
				}
			}
		}
	}

	fmt.Println("part2: ", sum)
}

func isFree(m [][]uint8, i, j int) bool {
	tile := m[i][j]
	if tile == '.' {
		return false
	}
	free := 0
	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= len(m) || y < 0 || y >= len(m[i]) || m[x][y] == '.' {
			free += 1
		}
	}
	return free > 4
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "09"

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

var dirMap = map[string]direction{
	"U": north,
	"R": east,
	"D": south,
	"L": west,
}

type point struct {
	x, y int
}

func part1(path string) {
	input := util.ToStringSlice(path)
	score := solve(input, make([]point, 2))
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToStringSlice(path)
	score := solve(input, make([]point, 10))
	fmt.Printf("part2: %d\n", score)
}

func solve(op []string, points []point) int {
	visited := map[point]bool{points[0]: true}

	for _, line := range op {
		parts := strings.Split(line, " ")
		dir := dirMap[parts[0]]
		n, _ := strconv.Atoi(parts[1])
		for range n {
			points[0] = point{points[0].x + dir[0], points[0].y + dir[1]}
			for i := 1; i < len(points); i++ {
				points[i] = move(points[i-1], points[i])
			}
			visited[points[len(points)-1]] = true
		}
	}

	return len(visited)
}

func move(h, a point) point {
	if util.Abs(h.x-a.x) > 1 || util.Abs(h.y-a.y) > 1 {
		for _, direction := range directions {
			b := point{a.x + direction[0], a.y + direction[1]}
			if util.Abs(b.x-h.x)+util.Abs(b.y-h.y) == 1 {
				return b
			}
		}
		for _, direction := range directions {
			b := point{a.x + direction[0], a.y + direction[1]}
			if util.Abs(b.x-h.x)+util.Abs(b.y-h.y) == 2 {
				return b
			}
		}
	}
	return a
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

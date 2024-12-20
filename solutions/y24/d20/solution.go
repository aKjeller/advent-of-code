package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "20"

func part1(path string) {
	fmt.Printf("part1: %d\n", solve(path, 2))
}

func part2(path string) {
	fmt.Printf("part2: %d\n", solve(path, 20))
}

func solve(inputPath string, cheatTime int) int {
	track := createTrack(inputPath)

	score := 0
	for pos, cost := range track {
		for x := -cheatTime; x <= cheatTime; x++ {
			for y := -cheatTime; y <= cheatTime; y++ {
				manhattan := util.Abs(x) + util.Abs(y)
				if manhattan <= cheatTime {
					cheat := vertex{x: pos.x + x, y: pos.y + y}
					if cheatCost, ok := track[cheat]; ok && cheatCost+manhattan <= cost-100 {
						score += 1
					}
				}
			}
		}
	}
	return score
}

type direction [2]int

var (
	NORTH direction = [2]int{-1, 0}
	SOUTH direction = [2]int{1, 0}
	WEST  direction = [2]int{0, -1}
	EAST  direction = [2]int{0, 1}
)

var directions = [4]direction{NORTH, SOUTH, WEST, EAST}

type vertex struct {
	x int
	y int
}

func createTrack(path string) map[vertex]int {
	input := util.ToStringSlice(path)

	var m [][]uint8
	var end vertex
	for i := 0; i < len(input); i++ {
		var row []uint8
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'E' {
				end = vertex{x: i, y: j}
			}
			row = append(row, input[i][j])
		}
		m = append(m, row)
	}
	dist := make(map[vertex]int)
	walkTrack(end, 0, m, dist)
	return dist
}

func walkTrack(curr vertex, depth int, m [][]uint8, dist map[vertex]int) {
	dist[curr] = depth
	for _, d := range directions {
		next := vertex{x: curr.x + d[0], y: curr.y + d[1]}
		if dist[next] == 0 && m[next.x][next.y] != '#' {
			walkTrack(next, depth+1, m, dist)
		}
	}
}

func main() {
	part1(util.InputPath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

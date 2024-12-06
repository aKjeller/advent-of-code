package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "06"

func part1(path string) {
	input := util.ToStringSlice(path)
	m, x, y := parse(input)
	visited, _ := walk(m, x, y)
	fmt.Printf("part1: %d\n", len(visited))
}

func walk(m [][]uint8, x, y int) (map[[2]int]int, error) {
	visited := make(map[[2]int]int)
	dx, dy := 0, -1
	for {
		visited[[2]int{x, y}] += 1
		if visited[[2]int{x, y}] > 4 {
			return visited, fmt.Errorf("loop")
		}

		if (x+dx < 0 || x+dx > len(m[0])-1) || (y+dy < 0 || y+dy > len(m)-1) {
			return visited, nil
		}

		if m[y+dy][x+dx] == '#' {
			dx, dy = -dy, dx
		} else {
			x, y = x+dx, y+dy
		}
	}
}

func part2(path string) {
	input := util.ToStringSlice(path)
	m, x, y := parse(input)

	score := 0
	visited, _ := walk(m, x, y)
	for k, _ := range visited {
		j, i := k[0], k[1]
		m[i][j] = '#'
		_, err := walk(m, x, y)
		if err != nil {
			score++
		}
		m[i][j] = '.'
	}

	fmt.Printf("part2: %d\n", score)
}

func parse(input []string) (m [][]uint8, x, y int) {
	for i := 0; i < len(input); i++ {
		var row []uint8
		for j := 0; j < len(input[i]); j++ {
			row = append(row, input[i][j])
			if input[i][j] == '^' {
				x, y = j, i
			}
		}
		m = append(m, row)
	}
	return m, x, y
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

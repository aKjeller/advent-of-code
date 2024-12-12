package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "12"

type plot struct {
	x     int
	y     int
	plant uint8
}

func part1(path string) {
	m := util.ToGrid(path)

	visited := make(map[plot]bool)

	var regions [][]plot
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if visited[plot{i, j, m[i][j]}] {
				continue
			}
			regions = append(regions, fill(m, i, j, m[i][j], visited))
		}
	}

	score := 0
	for _, region := range regions {
		for _, p := range region {
			score += getFences(m, p) * len(region)
		}
	}

	fmt.Printf("part1: %d\n", score)
}

func getFences(m [][]uint8, p plot) int {
	fences := 4
	if p.x+1 >= 0 && p.x+1 < len(m) && p.y >= 0 && p.y < len(m[0]) && m[p.x+1][p.y] == p.plant {
		fences -= 1
	}
	if p.x-1 >= 0 && p.x-1 < len(m) && p.y >= 0 && p.y < len(m[0]) && m[p.x-1][p.y] == p.plant {
		fences -= 1
	}
	if p.x >= 0 && p.x < len(m) && p.y+1 >= 0 && p.y+1 < len(m[0]) && m[p.x][p.y+1] == p.plant {
		fences -= 1
	}
	if p.x >= 0 && p.x < len(m) && p.y-1 >= 0 && p.y-1 < len(m[0]) && m[p.x][p.y-1] == p.plant {
		fences -= 1
	}
	return fences
}

func fill(m [][]uint8, x, y int, plant uint8, visited map[plot]bool) []plot {
	var plots []plot

	if x < 0 || x >= len(m) || y < 0 || y >= len(m[0]) || m[x][y] != plant {
		return plots
	}

	p := plot{x, y, m[x][y]}
	if visited[p] {
		return plots
	}

	visited[p] = true

	plots = append(plots, p)

	plots = append(plots, fill(m, x+1, y, plant, visited)...)
	plots = append(plots, fill(m, x-1, y, plant, visited)...)
	plots = append(plots, fill(m, x, y+1, plant, visited)...)
	plots = append(plots, fill(m, x, y-1, plant, visited)...)

	return plots
}

func part2(path string) {
	score := 0
	fmt.Printf("part2: %d\n", score)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

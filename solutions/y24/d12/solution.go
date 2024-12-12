package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"slices"
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
			score += len(getFences(m, make(map[[2]int][]dir), p)) * len(region)
		}
	}

	fmt.Printf("part1: %d\n", score)
}

type dir int

const (
	NO dir = iota
	U
	D
	L
	R
)

func getFences(m [][]uint8, fences map[[2]int][]dir, p plot) map[[2]int][]dir {
	if p.x+1 < 0 || p.x+1 >= len(m) || p.y < 0 || p.y >= len(m[0]) || m[p.x+1][p.y] != p.plant {
		fences[[2]int{p.x + 1, p.y}] = append(fences[[2]int{p.x + 1, p.y}], D)
	}
	if p.x-1 < 0 || p.x-1 >= len(m) || p.y < 0 || p.y >= len(m[0]) || m[p.x-1][p.y] != p.plant {
		fences[[2]int{p.x - 1, p.y}] = append(fences[[2]int{p.x - 1, p.y}], U)
	}
	if p.x < 0 || p.x >= len(m) || p.y+1 < 0 || p.y+1 >= len(m[0]) || m[p.x][p.y+1] != p.plant {
		fences[[2]int{p.x, p.y + 1}] = append(fences[[2]int{p.x, p.y + 1}], R)
	}
	if p.x < 0 || p.x >= len(m) || p.y-1 < 0 || p.y-1 >= len(m[0]) || m[p.x][p.y-1] != p.plant {
		fences[[2]int{p.x, p.y - 1}] = append(fences[[2]int{p.x, p.y - 1}], L)
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
		fences := make(map[[2]int][]dir)
		for _, p := range region {
			getFences(m, fences, p)
		}

		var xMin, xMax, yMin, yMax int
		for k, _ := range fences {
			xMin = min(xMin, k[0])
			yMin = min(yMin, k[1])
			xMax = max(xMax, k[0])
			yMax = max(yMax, k[1])
		}

		length := 0
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				if slices.Contains(fences[[2]int{i, j}], U) {
					for slices.Contains(fences[[2]int{i, j}], U) {
						j++
					}
					length++
				}
			}
		}
		for i := xMin; i <= xMax; i++ {
			for j := yMin; j <= yMax; j++ {
				if slices.Contains(fences[[2]int{i, j}], D) {
					for slices.Contains(fences[[2]int{i, j}], D) {
						j++
					}
					length++
				}
			}
		}

		for j := yMin; j <= yMax; j++ {
			for i := xMin; i <= xMax; i++ {
				if slices.Contains(fences[[2]int{i, j}], R) {
					for slices.Contains(fences[[2]int{i, j}], R) {
						i++
					}
					length++
				}
			}
		}
		for j := yMin; j <= yMax; j++ {
			for i := xMin; i <= xMax; i++ {
				if slices.Contains(fences[[2]int{i, j}], L) {
					for slices.Contains(fences[[2]int{i, j}], L) {
						i++
					}
					length++
				}
			}
		}

		score += length * len(region)
	}
	fmt.Printf("part2: %d\n", score)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

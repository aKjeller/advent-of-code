package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"sync"
	"sync/atomic"
)

const YEAR = "24"
const DAY = "06"

func part1(path string) {
	input := util.ToStringSlice(path)
	m, x, y := parse(input)
	visited, _ := walk(m, x, y)
	fmt.Printf("part1: %d\n", len(visited))
}

func part2(path string) {
	input := util.ToStringSlice(path)
	m, x, y := parse(input)

	wg := sync.WaitGroup{}
	var score atomic.Int32

	visited, _ := walk(m, x, y)
	for coord, _ := range visited {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t := util.DeepCopy(m)
			t[coord[0]][coord[1]] = '#'
			if _, err := walk(t, x, y); err != nil {
				score.Add(1)
			}
		}()
	}

	wg.Wait()

	fmt.Printf("part2: %d\n", score.Load())
}

func walk(m [][]uint8, x, y int) (map[[2]int]int, error) {
	visited := make(map[[2]int]int)
	dx, dy := -1, 0
	for {
		visited[[2]int{x, y}] += 1
		if visited[[2]int{x, y}] > 4 {
			return visited, fmt.Errorf("loop")
		}
		if (x+dx < 0 || x+dx > len(m)-1) || (y+dy < 0 || y+dy > len(m[0])-1) {
			return visited, nil
		}

		if m[x+dx][y+dy] == '#' {
			dx, dy = dy, -dx
		} else {
			x, y = x+dx, y+dy
		}
	}
}

func parse(input []string) (m [][]uint8, x, y int) {
	for i := 0; i < len(input); i++ {
		var row []uint8
		for j := 0; j < len(input[i]); j++ {
			row = append(row, input[i][j])
			if input[i][j] == '^' {
				x, y = i, j
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

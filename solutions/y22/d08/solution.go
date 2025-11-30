package main

import (
	"fmt"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "08"

type tree struct {
	x, y, height int
}

func part1(path string) {
	input := util.ToIntGrid(path)

	rows := getTreeRows(input)

	visible := make(map[tree]bool)
	for _, row := range rows {
		height := -1
		for _, tree := range row {
			if tree.height > height {
				visible[tree] = true
			}
			height = max(height, tree.height)
		}
	}

	score := len(visible)
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToIntGrid(path)

	rows := getTreeRows(input)

	scenic := make(map[tree]int)
	for _, row := range rows {
		for i, tree := range row {
			visible := findVisibleTrees(row[i+1:], tree.height)
			if _, ok := scenic[tree]; !ok {
				scenic[tree] = visible
			} else {
				scenic[tree] *= visible
			}
		}
	}

	score := 0
	for _, v := range scenic {
		score = max(score, v)
	}

	fmt.Printf("part2: %d\n", score)
}

func getTreeRows(input [][]int) (rows [][]tree) {
	for i := range input {
		var a, b, c, d []tree
		for j := range input[i] {
			a = append(a, tree{i, j, input[i][j]})
			b = append(b, tree{j, i, input[j][i]})
			c = append(c, tree{len(input) - i - 1, len(input[i]) - j - 1, input[len(input)-i-1][len(input[i])-j-1]})
			d = append(d, tree{len(input[i]) - j - 1, len(input) - i - 1, input[len(input[i])-j-1][len(input)-i-1]})
		}
		rows = append(rows, a, b, c, d)
	}
	return
}

func findVisibleTrees(trees []tree, height int) int {
	visible := 0
	for _, tree := range trees {
		visible++
		if tree.height >= height {
			return visible
		}
	}
	return visible
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

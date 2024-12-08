package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "08"

func part1(path string) {
	input := util.ToStringSlice(path)
	antennas := parse(input)

	nodes := make(map[[2]int]bool)
	for i, antenna := range antennas {
		for _, second := range antennas[i+1:] {
			if antenna[2] == second[2] {
				a1, b1 := getNode(antenna[0], antenna[1], second[0], second[1])
				a2, b2 := getNode(second[0], second[1], antenna[0], antenna[1])
				nodes[[2]int{a1, b1}] = true
				nodes[[2]int{a2, b2}] = true
			}
		}
	}

	score := 0
	for n, _ := range nodes {
		if n[0] >= 0 && n[0] < len(input) && n[1] >= 0 && n[1] < len(input[0]) {
			score += 1
		}
	}
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToStringSlice(path)
	antennas := parse(input)

	nodes := make(map[[2]int]bool)
	for i, antenna := range antennas {
		for _, second := range antennas[i+1:] {
			if antenna[2] == second[2] {
				for _, a := range getNodes(antenna[0], antenna[1], second[0], second[1]) {
					nodes[a] = true
				}
				for _, a := range getNodes(second[0], second[1], antenna[0], antenna[1]) {
					nodes[a] = true
				}
			}
		}
		nodes[[2]int{antenna[0], antenna[1]}] = true
	}

	score := 0
	for n, _ := range nodes {
		if n[0] >= 0 && n[0] < len(input) && n[1] >= 0 && n[1] < len(input[0]) {
			score += 1
		}
	}
	fmt.Printf("part2: %d\n", score)
}

func getNode(x1, y1, x2, y2 int) (int, int) {
	return x2 - (x2-x1)*2, y2 - (y2-y1)*2
}

func getNodes(x1, y1, x2, y2 int) [][2]int {
	var nodes [][2]int
	dx := (x2 - x1) * 2
	dy := (y2 - y1) * 2
	for {
		x1, x2 = x2-dx, x1
		y1, y2 = y2-dy, y1
		nodes = append(nodes, [2]int{x1, y1})
		if x1 < 0 || x1 > 1000 || y1 < 0 || y1 > 1000 {
			break
		}
	}
	return nodes
}

func parse(input []string) [][3]int {
	var data [][3]int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != '.' {
				data = append(data, [3]int{i, j, int(input[i][j])})
			}
		}
	}
	return data
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

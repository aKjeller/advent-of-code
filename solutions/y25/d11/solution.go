package main

import (
	"fmt"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "11"

func both(path string) {
	input := util.ToStringSlice(path)

	graph := make(map[string][]string)
	for _, line := range input {
		graph[line[:3]] = strings.Split(line[5:], " ")
	}

	part1 := dfs(graph, "you", "out", make(map[string]int))

	svr_dac := dfs(graph, "svr", "dac", make(map[string]int))
	dac_fft := dfs(graph, "dac", "fft", make(map[string]int))
	fft_out := dfs(graph, "fft", "out", make(map[string]int))

	svr_fft := dfs(graph, "svr", "fft", make(map[string]int))
	fft_dac := dfs(graph, "fft", "dac", make(map[string]int))
	dac_out := dfs(graph, "dac", "out", make(map[string]int))

	part2 := (svr_dac * dac_fft * fft_out) + (svr_fft * fft_dac * dac_out)

	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
}

func dfs(graph map[string][]string, current, end string, cache map[string]int) int {
	if v, ok := cache[current]; ok {
		return v
	}

	if current == end {
		return 1
	}

	sum := 0
	for _, edge := range graph[current] {
		sum += dfs(graph, edge, end, cache)
	}
	cache[current] = sum
	return sum
}

func main() {
	// both(util.ExamplePath(YEAR, DAY))
	both(util.InputPath(YEAR, DAY))
}

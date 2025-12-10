package main

import (
	"container/heap"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
	"github.com/aKjeller/advent-of-code/utilities/go/ds"
)

const YEAR = "25"
const DAY = "10"

var re1 = regexp.MustCompile(`\[(.*)\]`)
var re2 = regexp.MustCompile(`\((.*?)\)`)
var re3 = regexp.MustCompile(`{(.*)}`)

type button []int

func part2(path string) {
	input := util.ToStringSlice(path)

	part2 := 0
	tot := len(input)
	for i, line := range input {
		ans := solve(line)
		fmt.Println(fmt.Sprintf("%d/%d : %d", i+1, tot, ans))
		part2 += ans
	}

	fmt.Println("part2: ", part2)
}

func solve(line string) int {
	var goal []int
	for i := range strings.SplitSeq(re3.FindStringSubmatch(line)[1], ",") {
		goal = append(goal, util.ParseInt(i))
	}

	// skip large ones
	for _, g := range goal {
		if g > 60 {
			return -1
		}
	}

	buttons := getButtonIndexes(line)
	var niceButtons []button
	for _, button := range buttons {
		niceButton := make([]int, len(goal))
		for _, i := range button {
			niceButton[i] = 1
		}
		niceButtons = append(niceButtons, niceButton)
	}
	ans := brute(niceButtons, 0, make([]int, len(goal)), goal)
	return ans
}

func brute(buttons []button, index int, current, goal []int) int {

	if slices.Equal(current, goal) {
		return 0
	}
	if index == len(goal) {
		return 800000
	}

	// maybe it is faster to do the lowest target first, instead of 0..n
	target := goal[index] - current[index]
	if target == 0 {
		return brute(buttons, index+1, current, goal)
	}
	var buttonsForIndex []button
	for _, button := range buttons {
		ok := true
		for i := range current {
			if current[i] >= goal[i] && button[i] == 1 {
				ok = false
				break
			}
		}
		if ok && button[index] == 1 {
			buttonsForIndex = append(buttonsForIndex, button)
		}
	}

	if len(buttonsForIndex) == 0 {
		return 900000
	}

	// memory issues, this will store the whole search space in memory, get counts iterativly instead
	allValid := chooseCounts(len(buttonsForIndex), target)
	presses := math.MaxInt
	for _, valid := range allValid {
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		for i, times := range valid {
			for range times {
				newCurrent = AddSlices(newCurrent, buttonsForIndex[i])
			}
		}
		res := brute(buttons, index+1, newCurrent, goal)
		if res == -1 {
			continue
		}
		if target+res < presses {
			presses = target + res
		}

	}
	return presses
}

func AddSlices(a, b []int) []int {
	if len(a) != len(b) {
		panic("slices must have same length")
	}

	res := make([]int, len(a))
	for i := range a {
		res[i] = a[i] + b[i]
	}
	return res
}

func Sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func chooseCounts(n, target int) [][]int {
	var res [][]int
	counts := make([]int, n)

	var dfs func(pos, remaining int)
	dfs = func(pos, remaining int) {
		if pos == n-1 {
			counts[pos] = remaining
			cpy := make([]int, n)
			copy(cpy, counts)
			res = append(res, cpy)
			return
		}

		for k := 0; k <= remaining; k++ {
			counts[pos] = k
			dfs(pos+1, remaining-k)
		}
	}

	dfs(0, target)
	return res
}

func part1(path string) {
	input := util.ToStringSlice(path)

	part1 := 0

	for _, line := range input {
		end := vertex(re1.FindStringSubmatch(line)[1])
		start := vertex(strings.Repeat(".", len(end)))

		buttons := getButtonIndexes(line)

		graph := make(map[vertex][]edge)
		for _, vertex := range vertices(len(end)) {
			var edges []edge
			for _, button := range buttons {
				edge := edge{
					start: vertex,
					end:   flip(vertex, button),
					cost:  1,
				}
				edges = append(edges, edge)
			}
			graph[vertex] = edges
		}

		part1 += dijkstra(graph, start, end)
	}

	fmt.Println("part1: ", part1)
}

func getButtonIndexes(line string) []button {
	var buttonIndexes []button
	buttons := re2.FindAllStringSubmatch(line, -1)
	for _, b := range buttons {
		var button button
		for i := range strings.SplitSeq(b[1], ",") {
			button = append(button, util.ParseInt(i))
		}
		buttonIndexes = append(buttonIndexes, button)
	}
	return buttonIndexes
}

func flip(v vertex, indexes []int) vertex {
	b := []byte(v)
	for _, index := range indexes {
		if b[index] == '.' {
			b[index] = '#'
		} else {
			b[index] = '.'
		}
	}
	return vertex(b)
}

func vertices(n int) []vertex {
	if n == 0 {
		return []vertex{""}
	}

	prev := vertices(n - 1)

	var res []vertex
	for _, s := range prev {
		res = append(res, s+".")
		res = append(res, s+"#")
	}

	return res
}

type vertex string

type edge struct {
	start vertex
	end   vertex
	cost  int
}

func dijkstra(graph map[vertex][]edge, start, end vertex) int {
	dist := make(map[vertex]int)
	for v := range graph {
		dist[v] = math.MaxInt
	}
	dist[start] = 0

	unvisited := &ds.PriorityQueue[vertex]{}
	heap.Init(unvisited)
	heap.Push(unvisited, ds.Item[vertex]{Value: start, Priority: 0})

	visited := make(map[vertex]bool)

	for unvisited.Len() > 0 {
		current := heap.Pop(unvisited).(ds.Item[vertex]).Value

		if current == end {
			return dist[current]
		}

		if visited[current] {
			continue
		}

		visited[current] = true
		for _, e := range graph[current] {
			if visited[e.end] {
				continue
			}
			newDist := dist[current] + e.cost
			if newDist < dist[e.end] {
				dist[e.end] = newDist
				heap.Push(unvisited, ds.Item[vertex]{Value: e.end, Priority: newDist})
			}
		}
	}
	return -1
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	// Might have to write a ILP-solver
	// part2(util.InputPath(YEAR, DAY))
}

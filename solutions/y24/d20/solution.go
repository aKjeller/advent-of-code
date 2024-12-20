package main

import (
	"container/heap"
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"github.com/aKjeller/advent-of-code/utilities/go/ds"
	"math"
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
	m := util.ToGrid(inputPath)
	graph, start, end := createGraph(m)
	dist, path := dijkstra(graph, end, start)

	score := 0
	for _, p := range path {
		for x := -cheatTime; x <= cheatTime; x++ {
			for y := -cheatTime; y <= cheatTime; y++ {
				manhattan := util.Abs(x) + util.Abs(y)
				if manhattan > cheatTime {
					continue
				}
				newStart := vertex{x: p.x + x, y: p.y + y}
				if _, newStartExists := graph[newStart]; !newStartExists {
					continue
				}
				if newDist, ok := dist[newStart]; ok && newDist+manhattan <= dist[p]-100 {
					score += 1
				}
			}
		}
	}
	return score
}

func createGraph(m [][]uint8) (graph map[vertex][]edge, start, end vertex) {
	graph = make(map[vertex][]edge)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			s := m[i][j]
			if s != '#' {
				v := vertex{x: i, y: j}
				var edges []edge
				for _, d := range directions {
					if m[i+d[0]][j+d[1]] != '#' {
						edges = append(edges, edge{start: v, end: vertex{x: i + d[0], y: j + d[1]}})
					}

					if s == 'S' && d == EAST {
						start = v
					} else if s == 'E' {
						end = v
					}
				}
				graph[v] = edges
			}
		}
	}
	return graph, start, end
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

type edge struct {
	start vertex
	end   vertex
}

func dijkstra(graph map[vertex][]edge, start, end vertex) (map[vertex]int, []vertex) {
	dist := make(map[vertex]int)
	for v := range graph {
		dist[v] = math.MaxInt
	}
	dist[start] = 0

	unvisited := &ds.PriorityQueue[vertex]{}
	heap.Init(unvisited)
	heap.Push(unvisited, ds.Item[vertex]{Value: start, Priority: 0})

	visited := make(map[vertex]bool)
	previous := make(map[vertex]vertex)

	for unvisited.Len() > 0 {
		current := heap.Pop(unvisited).(ds.Item[vertex]).Value

		if current == end {
			return dist, getPath(end, previous)
		}

		if visited[current] {
			continue
		}

		visited[current] = true
		for _, e := range graph[current] {
			if visited[e.end] {
				continue
			}
			newDist := dist[current] + 1
			if newDist < dist[e.end] {
				dist[e.end] = newDist
				heap.Push(unvisited, ds.Item[vertex]{Value: e.end, Priority: newDist})
				previous[e.end] = current
			}
		}
	}
	return nil, nil
}

func getPath(current vertex, prev map[vertex]vertex) []vertex {
	var path []vertex
	path = append(path, current)
	if v, ok := prev[current]; ok {
		return append(path, getPath(v, prev)...)
	}

	return path
}

func main() {
	part1(util.InputPath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

package main

import (
	"container/heap"
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"github.com/aKjeller/advent-of-code/utilities/go/ds"
	"math"
)

const YEAR = "24"
const DAY = "18"

func part1(path string, size, length int) {
	graph := createGraph(util.ToStringSlice(path), size, length)
	distance := dijkstra(graph, vertex{0, 0}, vertex{size - 1, size - 1})
	fmt.Printf("part1: %d\n", distance)
}

func part2(path string, size int) {
	input := util.ToStringSlice(path)
	start, end := vertex{0, 0}, vertex{size - 1, size - 1}
	low, high := 0, len(input)-1
	for high-low > 1 {
		graph := createGraph(input, size, high-((high-low)/2))
		distance := dijkstra(graph, start, end)
		if distance < 0 {
			high = high - ((high - low) / 2)
		} else {
			low = low + ((high - low) / 2)
		}
	}
	fmt.Printf("part2: %s\n", input[high])
}

func createGraph(input []string, size, length int) (graph map[vertex][]edge) {
	m := make(map[[2]int]uint8)
	for r := range length + 1 {
		nums := util.GetIntsFromString(input[r])
		m[[2]int{nums[1], nums[0]}] = '#'
	}
	for i := range size {
		m[[2]int{i, -1}] = '#'
		m[[2]int{i, size}] = '#'
		m[[2]int{-1, i}] = '#'
		m[[2]int{size, i}] = '#'
	}

	graph = make(map[vertex][]edge)
	for i := -1; i < size+1; i++ {
		for j := -1; j < size+1; j++ {
			if m[[2]int{i, j}] != '#' {
				v := vertex{x: i, y: j}
				var edges []edge
				for _, d := range directions {
					if m[[2]int{i + d[0], j + d[1]}] != '#' {
						edges = append(edges, edge{start: v, end: vertex{x: i + d[0], y: j + d[1]}, cost: 1})
					}
				}
				graph[v] = edges
			}
		}
	}
	return graph
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

		if current.x == end.x && current.y == end.y {
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
	part1(util.ExamplePath(YEAR, DAY), 7, 12)
	part1(util.InputPath(YEAR, DAY), 71, 1024)
	part2(util.ExamplePath(YEAR, DAY), 7)
	part2(util.InputPath(YEAR, DAY), 71)
}

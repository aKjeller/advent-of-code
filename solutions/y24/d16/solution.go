package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "16"

func both(path string) {
	graph, start, end := createGraph(util.ToGrid(path))
	distance, unique := dijkstra(graph, start, end)
	fmt.Printf("part1: %d\n", distance)
	fmt.Printf("part2: %d\n", unique)
}

func createGraph(m [][]uint8) (graph map[vertex][]edge, start, end vertex) {
	graph = make(map[vertex][]edge)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			s := m[i][j]
			if s != '#' {
				for _, d := range directions {
					v := vertex{x: i, y: j, dir: d}
					var edges []edge
					if m[i+d[0]][j+d[1]] != '#' {
						edges = append(edges, edge{start: v, end: vertex{x: i + d[0], y: j + d[1], dir: d}, cost: 1})
					}
					edges = append(edges, edge{start: v, end: vertex{x: i, y: j, dir: [2]int{d[1], -d[0]}}, cost: 1000})
					edges = append(edges, edge{start: v, end: vertex{x: i, y: j, dir: [2]int{-d[1], d[0]}}, cost: 1000})
					graph[v] = edges

					if s == 'S' && d == EAST {
						start = v
					} else if s == 'E' {
						end = v
					}
				}
			}
		}
	}
	return graph, start, end
}

func main() {
	both(util.ExamplePath(YEAR, DAY))
	both(util.InputPath(YEAR, DAY))
}

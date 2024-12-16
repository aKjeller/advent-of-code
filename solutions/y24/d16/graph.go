package main

import (
	"math"
)

type direction [2]int

var (
	NORTH direction = [2]int{-1, 0}
	SOUTH direction = [2]int{1, 0}
	WEST  direction = [2]int{0, -1}
	EAST  direction = [2]int{0, 1}
)

var directions = [4]direction{NORTH, SOUTH, WEST, EAST}

type vertex struct {
	x   int
	y   int
	dir direction
}

type edge struct {
	start vertex
	end   vertex
	cost  int
}

func dijkstra(graph map[vertex][]edge, start, end vertex) (int, int) {
	dist := make(map[vertex]int)
	for v := range graph {
		dist[v] = math.MaxInt
	}

	unvisited := make(map[vertex]bool)
	visited := make(map[vertex]bool)
	previous := make(map[vertex][]vertex)

	dist[start] = 0
	unvisited[start] = true

	for len(unvisited) > 0 {
		current := getLowestCost(unvisited, dist)
		delete(unvisited, current)

		if current.x == end.x && current.y == end.y {
			unique := make(map[[2]int]bool)
			getUniqueVertices(current, previous, unique)
			return dist[current], len(unique)
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
				unvisited[e.end] = true
				previous[e.end] = []vertex{current}
			} else if newDist == dist[e.end] {
				previous[e.end] = append(previous[e.end], current)
			}
		}
	}
	return -1, -1
}

func getLowestCost(m map[vertex]bool, dist map[vertex]int) vertex {
	var lowest vertex
	lowestDistance := math.MaxInt
	for v := range m {
		if dist[v] < lowestDistance {
			lowest, lowestDistance = v, dist[v]
		}
	}
	return lowest
}

func getUniqueVertices(current vertex, prev map[vertex][]vertex, unique map[[2]int]bool) {
	unique[[2]int{current.x, current.y}] = true
	if prev[current] == nil || len(prev[current]) == 0 {
		return
	}

	for _, next := range prev[current] {
		getUniqueVertices(next, prev, unique)
	}
}

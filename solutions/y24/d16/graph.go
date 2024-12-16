package main

import (
	"container/heap"
	"github.com/aKjeller/advent-of-code/utilities/go/ds"
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
	dist[start] = 0

	unvisited := &ds.PriorityQueue[vertex]{}
	heap.Init(unvisited)
	heap.Push(unvisited, ds.Item[vertex]{Value: start, Priority: 0})

	visited := make(map[vertex]bool)
	previous := make(map[vertex][]vertex)

	for unvisited.Len() > 0 {
		current := heap.Pop(unvisited).(ds.Item[vertex]).Value

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
				heap.Push(unvisited, ds.Item[vertex]{Value: e.end, Priority: newDist})
				previous[e.end] = []vertex{current}
			} else if newDist == dist[e.end] {
				previous[e.end] = append(previous[e.end], current)
			}
		}
	}
	return -1, -1
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

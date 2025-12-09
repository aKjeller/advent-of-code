package main

import (
	"container/heap"
	"fmt"
	"maps"
	"slices"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
	"github.com/aKjeller/advent-of-code/utilities/go/ds"
)

const YEAR = "25"
const DAY = "08"

type box struct {
	x, y, z int
	index   int
}

func (b box) distance(other box) int {
	dx := b.x - other.x
	dy := b.y - other.y
	dz := b.z - other.z
	return dx*dx + dy*dy + dz*dz
}

type edge struct {
	src, dst box
}

func both(path string, n int) {
	dsu, boxes, edges := parse(path)

	var last edge
	var c1, c2, part1 int
	target := len(boxes) - 1

	for c2 < target {
		edge := heap.Pop(edges).(ds.Item[edge]).Value

		if dsu.Union(edge.src.index, edge.dst.index) {
			last = edge
			c2++
		}

		if c1 == n-1 {
			part1 = getPart1(dsu)
		}
		c1++
	}

	part2 := last.src.x * last.dst.x

	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
}

func getPart1(dsu *ds.Dsu) int {
	circuits := make(map[int]int)
	for _, union := range *dsu {
		circuits[dsu.Find(union.Parent)] += 1
	}

	values := slices.Sorted(maps.Values(circuits))
	slices.Reverse(values)

	return values[0] * values[1] * values[2]
}

func parse(path string) (*ds.Dsu, []box, *ds.PriorityQueue[edge]) {
	input := util.ToStringSlice(path)

	var dsu ds.Dsu
	var boxes []box
	for _, line := range input {
		parts := strings.Split(line, ",")
		index := dsu.Add()
		box := box{util.ParseInt(parts[0]), util.ParseInt(parts[1]), util.ParseInt(parts[2]), index}
		boxes = append(boxes, box)
	}

	edges := &ds.PriorityQueue[edge]{}
	heap.Init(edges)
	for i := range boxes {
		for j := i + 1; j < len(input); j++ {
			heap.Push(edges, ds.Item[edge]{Value: edge{boxes[i], boxes[j]}, Priority: boxes[i].distance(boxes[j])})
		}
	}

	return &dsu, boxes, edges
}

func main() {
	// both(util.ExamplePath(YEAR, DAY), 10)
	both(util.InputPath(YEAR, DAY), 1000)
}

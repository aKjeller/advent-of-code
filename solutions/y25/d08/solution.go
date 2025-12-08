package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "08"

type box struct {
	x, y, z int
	edges   []*box
}

func (b box) String() string {
	return fmt.Sprintf("[%d %d %d]", b.x, b.y, b.z)

}

func (b box) distance(other box) float64 {
	dx := float64(b.x - other.x)
	dy := float64(b.y - other.y)
	dz := float64(b.z - other.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type edge struct {
	a, b     *box
	distance float64
}

func part1(path string) {
	input := util.ToStringSlice(path)

	var boxes []*box
	var networks [][]*box
	for _, line := range input {
		parts := strings.Split(line, ",")
		b := box{util.ParseInt(parts[0]), util.ParseInt(parts[1]), util.ParseInt(parts[2]), []*box{}}
		boxes = append(boxes, &b)
		networks = append(networks, []*box{&b})
	}

	for range 1000 {
		var edges []edge
		for i := range boxes {
			for j := range boxes {
				boxA := boxes[i]
				boxB := boxes[j]
				if boxA == boxB {
					continue
				}
				if slices.Contains(boxA.edges, boxB) {
					continue
				}
				edges = append(edges, edge{boxA, boxB, boxA.distance(*boxB)})
			}
		}
		if len(edges) == 0 || len(networks) == 1 {
			break
		}
		edge := edges[0]
		for _, e := range edges {
			if e.distance < edge.distance {
				edge = e
			}
		}

		edge.a.edges = append(edge.a.edges, edge.b)
		edge.b.edges = append(edge.b.edges, edge.a)

		a, b := -1, -1
		for i, network := range networks {
			if slices.Contains(network, edge.a) {
				a = i
			}
			if slices.Contains(network, edge.b) {
				b = i
			}
		}
		if a == b {
			continue
		}
		if a > -1 && b > -1 {
			networks[a] = append(networks[a], networks[b]...)
			networks = util.RemoveElement(networks, b)
		}
	}

	sort.Slice(networks, func(i, j int) bool {
		return len(networks[i]) > len(networks[j])
	})

	res := 0
	if len(networks) >= 3 {
		res = len(networks[0]) * len(networks[1]) * len(networks[2])
	}
	fmt.Println("part1: ", res)
}

func part2(path string) {
	input := util.ToStringSlice(path)

	var boxes []*box
	var networks [][]*box
	for _, line := range input {
		parts := strings.Split(line, ",")
		b := box{util.ParseInt(parts[0]), util.ParseInt(parts[1]), util.ParseInt(parts[2]), []*box{}}
		boxes = append(boxes, &b)
		networks = append(networks, []*box{&b})
	}

	var last edge
	for {
		var edges []edge
		for i := range boxes {
			for j := range boxes {
				boxA := boxes[i]
				boxB := boxes[j]
				if boxA == boxB {
					continue
				}
				if slices.Contains(boxA.edges, boxB) {
					continue
				}
				edges = append(edges, edge{boxA, boxB, boxA.distance(*boxB)})
			}
		}
		if len(edges) == 0 || len(networks) == 1 {
			break
		}
		edge := edges[0]
		for _, e := range edges {
			if e.distance < edge.distance {
				edge = e
			}
		}

		last = edge

		edge.a.edges = append(edge.a.edges, edge.b)
		edge.b.edges = append(edge.b.edges, edge.a)

		a, b := -1, -1
		for i, network := range networks {
			if slices.Contains(network, edge.a) {
				a = i
			}
			if slices.Contains(network, edge.b) {
				b = i
			}
		}
		if a == b {
			continue
		}
		if a > -1 && b > -1 {
			networks[a] = append(networks[a], networks[b]...)
			networks = util.RemoveElement(networks, b)
		}
	}

	res := last.a.x * last.b.x
	fmt.Println("part2: ", res)
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

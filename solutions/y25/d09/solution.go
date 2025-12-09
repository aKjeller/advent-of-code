package main

import (
	"fmt"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "09"

type point struct {
	x, y int
}

type side struct {
	start, end point
}

func both(path string) {
	input := util.ToStringSlice(path)

	var points []point
	for i := range input {
		l1 := strings.Split(input[i], ",")
		p := point{util.ParseInt(l1[1]), util.ParseInt(l1[0])}
		points = append(points, p)
	}
	points = append(points, points[0])

	var sides []side
	for i := 1; i < len(points); i++ {
		sides = append(sides, side{points[i-1], points[i]})
	}

	var part1, part2 int
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := (util.Abs(points[j].x-points[i].x) + 1) * (util.Abs(points[j].y-points[i].y) + 1)
			if area > part2 && isValid(sides, points[i], points[j]) {
				part2 = area
			}
			if area > part1 {
				part1 = area
			}
		}
	}

	fmt.Println("part1: ", part1)
	fmt.Println("part2: ", part2)
}

func isValid(sides []side, p1, p2 point) bool {
	xMin, xMax, yMin, yMax := min(p1.x, p2.x), max(p1.x, p2.x), min(p1.y, p2.y), max(p1.y, p2.y)
	for _, side := range sides {
		if side.start.y == side.end.y {
			if side.start.y <= yMin || side.start.y >= yMax {
				continue
			}
			if max(side.start.x, side.end.x) > xMin && min(side.start.x, side.end.x) < xMax {
				return false
			}
		} else {
			if side.start.x <= xMin || side.start.x >= xMax {
				continue
			}
			if max(side.start.y, side.end.y) > yMin && min(side.start.y, side.end.y) < yMax {
				return false
			}
		}
	}
	return true
}

func main() {
	// both(util.ExamplePath(YEAR, DAY))
	both(util.InputPath(YEAR, DAY))
}

package main

import (
	"fmt"
	"sort"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "09"

func part1(path string) {
	input := util.ToStringSlice(path)

	area := 0
	for i := range input {
		l1 := strings.Split(input[i], ",")
		aX, aY := util.ParseInt(l1[0]), util.ParseInt(l1[1])
		for j := i + 1; j < len(input); j++ {
			l2 := strings.Split(input[j], ",")
			bX, bY := util.ParseInt(l2[0]), util.ParseInt(l2[1])

			newArea := (util.Abs(bX-aX) + 1) * (util.Abs(bY-aY) + 1)
			if newArea > area {
				area = newArea
			}

		}
	}

	fmt.Println("part1: ", area)
}

type point struct {
	x, y int
}

type side struct {
	start, end point
}

type rect struct {
	start, end point
	area       int
}

func part2(path string) {
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

	var rects []rect
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			area := (util.Abs(points[j].x-points[i].x) + 1) * (util.Abs(points[j].y-points[i].y) + 1)
			rects = append(rects, rect{points[i], points[j], area})
		}
	}

	sort.Slice(rects, func(i, j int) bool {
		return rects[i].area > rects[j].area
	})

	area := 0
	for _, rect := range rects {
		if isValid(sides, rect.start, rect.end) {
			area = rect.area
			break
		}
	}

	fmt.Println("part2: ", area)
}

func isValid(sides []side, p1, p2 point) bool {
	xMin, xMax, yMin, yMax := min(p1.x, p2.x), max(p1.x, p2.x), min(p1.y, p2.y), max(p1.y, p2.y)
	for _, side := range sides {
		if side.start.y == side.end.y {
			for x := min(side.start.x, side.end.x); x <= max(side.start.x, side.end.x); x++ {
				if x > xMin && x < xMax && side.start.y > yMin && side.start.y < yMax {
					return false
				}
			}
		} else {
			for y := min(side.start.y, side.end.y); y <= max(side.start.y, side.end.y); y++ {
				if y > yMin && y < yMax && side.start.x > xMin && side.start.x < xMax {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

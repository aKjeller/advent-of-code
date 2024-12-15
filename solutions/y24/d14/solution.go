package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"strconv"
)

const YEAR = "24"
const DAY = "14"

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func (r *robot) move(width, height int) {
	r.x += r.dx
	r.y += r.dy
	if r.x < 0 {
		r.x += width
	}
	if r.x >= width {
		r.x -= width
	}
	if r.y < 0 {
		r.y += height
	}
	if r.y >= height {
		r.y -= height
	}
}

func part1(path string, width, height int) {
	input := util.ToStringSlice(path)

	var robots []*robot
	for _, r := range input {
		nums := util.GetIntsFromStringWithNegative(r)
		robots = append(robots, &robot{x: nums[0], y: nums[1], dx: nums[2], dy: nums[3]})
	}

	for range 100 {
		for _, r := range robots {
			r.move(width, height)
		}
	}
	//printMap(robots, width, height)

	score := getSafetyFactor(makeMap(robots), width, height)
	fmt.Printf("part1: %d\n", score)
}

func part2(path string, width, height int) {
	input := util.ToStringSlice(path)

	var robots []*robot
	for _, r := range input {
		nums := util.GetIntsFromStringWithNegative(r)
		robots = append(robots, &robot{x: nums[0], y: nums[1], dx: nums[2], dy: nums[3]})
	}

	score := 0
	for {
		score += 1
		for _, r := range robots {
			r.move(width, height)
		}
		if isTree(robots, makeMap(robots)) {
			break
		}
	}
	printMap(robots, width, height)

	fmt.Printf("part2: %d\n", score)
}

// this is a julgran
// ..#..
// .###.
// #####
// ..#..
func isTree(robots []*robot, m map[[2]int]int) bool {
	for _, r := range robots {
		if m[[2]int{r.x - 1, r.y + 1}] == 0 {
			return false
		}
		if m[[2]int{r.x, r.y + 1}] == 0 {
			return false
		}
		if m[[2]int{r.x + 1, r.y + 1}] == 0 {
			return false
		}
		if m[[2]int{r.x - 1, r.y + 2}] == 0 {
			return false
		}
		if m[[2]int{r.x - 2, r.y + 2}] == 0 {
			return false
		}
		if m[[2]int{r.x - 1, r.y + 2}] == 0 {
			return false
		}
		if m[[2]int{r.x, r.y + 2}] == 0 {
			return false
		}
		if m[[2]int{r.x + 1, r.y + 2}] == 0 {
			return false
		}
		if m[[2]int{r.x + 2, r.y + 2}] == 0 {
			return false
		}
		if m[[2]int{r.x, r.y + 3}] == 0 {
			return false
		}
		return true
	}
	return false
}

func getSafetyFactor(m map[[2]int]int, width, height int) int {
	var a, b, c, d int
	for r, v := range m {
		if r[0] < width/2 && r[1] < height/2 {
			a += v
		}
		if r[0] > width/2 && r[1] < height/2 {
			b += v
		}
		if r[0] < width/2 && r[1] > height/2 {
			c += v
		}
		if r[0] > width/2 && r[1] > height/2 {
			d += v
		}
	}
	return a * b * c * d
}

func makeMap(robots []*robot) map[[2]int]int {
	m := make(map[[2]int]int)
	for _, r := range robots {
		m[[2]int{r.x, r.y}] += 1
	}
	return m
}

func printMap(robots []*robot, width, height int) {
	m := makeMap(robots)
	for i := 0; i < height; i++ {
		row := ""
		for j := 0; j < width; j++ {
			num := m[[2]int{j, i}]
			if num == 0 {
				row += "."
			} else {
				row += strconv.Itoa(num)
			}
		}
		fmt.Println(row)
	}
}

func main() {
	part1(util.ExamplePath(YEAR, DAY), 11, 7)
	part1(util.InputPath(YEAR, DAY), 101, 103)
	//part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY), 101, 103)
}

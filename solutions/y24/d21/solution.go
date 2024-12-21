package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"math"
	"strings"
)

const YEAR = "24"
const DAY = "21"

type point struct {
	x, y int
}

var keys = map[uint8]point{
	'7': {x: 0, y: 0},
	'8': {x: 0, y: 1},
	'9': {x: 0, y: 2},
	'4': {x: 1, y: 0},
	'5': {x: 1, y: 1},
	'6': {x: 1, y: 2},
	'1': {x: 2, y: 0},
	'2': {x: 2, y: 1},
	'3': {x: 2, y: 2},
	'0': {x: 3, y: 1},
	'A': {x: 3, y: 2},
	'^': {x: 4, y: 1},
	'a': {x: 4, y: 2},
	'<': {x: 5, y: 0},
	'v': {x: 5, y: 1},
	'>': {x: 5, y: 2},
}

func part1(path string) {
	fmt.Printf("part2: %d\n", both(path, 2))
}

func part2(path string) {
	fmt.Printf("part2: %d\n", both(path, 25))
}

func both(path string, maxLevel int) int {
	input := util.ToStringSlice(path)
	score := 0
	for _, r := range input {
		score += solve(r, 0, maxLevel, make(map[item]int)) * util.ParseInt(r[:len(r)-1])
	}
	return score
}

type item struct {
	code  string
	level int
}

func solve(code string, level, maxLevel int, cache map[item]int) int {
	if v, ok := cache[item{code, level}]; ok {
		return v
	}

	if level > maxLevel {
		return len(code)
	}

	var start point
	if level == 0 {
		start = keys['A']
	} else {
		start = keys['a']
	}

	result := 0
	for i, _ := range code {
		shortest := math.MaxInt
		dx, dy := start.x-keys[code[i]].x, start.y-keys[code[i]].y
		for _, p := range permutations(dx, dy) {
			if allowedPath(p, start) {
				shortest = min(shortest, solve(p+"a", level+1, maxLevel, cache))
			}
		}
		start = point{x: start.x - dx, y: start.y - dy}
		result += shortest
	}

	cache[item{code, level}] = result
	return result
}

func allowedPath(code string, start point) bool {
	for i, _ := range code {
		if code[i] == '^' {
			start.x -= 1
		}
		if code[i] == 'v' {
			start.x += 1
		}
		if code[i] == '<' {
			start.y -= 1
		}
		if code[i] == '>' {
			start.y += 1
		}
		exist := false
		for _, v := range keys {
			if v == start {
				exist = true
			}
		}
		if !exist {
			return false
		}
	}
	return true
}

func permutations(dx, dy int) []string {
	pattern := ""
	if dx < 0 {
		pattern += strings.Repeat("v", util.Abs(dx))
	} else {
		pattern += strings.Repeat("^", util.Abs(dx))
	}
	if dy < 0 {
		pattern += strings.Repeat(">", util.Abs(dy))
	} else {
		pattern += strings.Repeat("<", util.Abs(dy))
	}
	return util.Permutations(pattern)
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

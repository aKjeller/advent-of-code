package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
	"sync"
	"sync/atomic"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "10"

var re1 = regexp.MustCompile(`\[(.*)\]`)
var re2 = regexp.MustCompile(`\((.*?)\)`)
var re3 = regexp.MustCompile(`{(.*)}`)

type button []int

func part1(path string) {
	input := util.ToStringSlice(path)

	part1 := 0
	for _, line := range input {
		target := getTargetPart1(line)
		buttons := getButtons(line, len(target))
		combos := getCombos(buttons, target)
		best := math.MaxInt32
		for _, combo := range combos {
			best = min(best, len(combo))
		}
		part1 += best
	}

	fmt.Println("part1: ", part1)
}

func part2(path string) {
	input := util.ToStringSlice(path)

	var part2 atomic.Int64
	wg := sync.WaitGroup{}

	for _, line := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			target := getTargetPart2(line)
			buttons := getButtons(line, len(target))
			part2.Add(int64(solve(buttons, target, make(map[string]int))))
		}()
	}

	wg.Wait()
	fmt.Println("part2: ", part2.Load())
}

func solve(buttons []button, target []int, cache map[string]int) int {
	key := fmt.Sprint(target)
	if v, ok := cache[key]; ok {
		return v
	}

	if slices.Equal(target, make([]int, len(target))) {
		return 0
	}

	pattern := make([]int, len(target))
	for i := range target {
		if target[i]%2 == 1 {
			pattern[i] = 1
		}
	}

	combos := getCombos(buttons, pattern)

	best := math.MaxInt32

outer:
	for _, combo := range combos {
		current := slices.Clone(target)
		for _, v := range combo {
			current = sub(current, buttons[v])
		}

		for _, v := range current {
			if v < 0 {
				continue outer
			}
		}

		current = divide(current, 2)

		best = min(best, 2*solve(buttons, current, cache)+len(combo))
	}

	cache[key] = best
	return best
}

func sub(a, b []int) []int {
	for i := range a {
		a[i] = a[i] - b[i]
	}
	return a
}

func divide(a []int, divisor int) []int {
	for i := range a {
		a[i] = a[i] / divisor
	}
	return a
}

func getCombos(buttons []button, target []int) [][]int {
	n := len(target)
	current := make([]int, n)
	var combos [][]int

	var dfs func(idx int, chosen []int)
	dfs = func(i int, chosen []int) {
		if i == len(buttons) {
			if slices.Equal(current, target) {
				combo := slices.Clone(chosen)
				combos = append(combos, combo)
			}
			return
		}

		// skip
		dfs(i+1, chosen)

		// take
		row := buttons[i]
		for i := range n {
			current[i] ^= row[i]
		}
		dfs(i+1, append(chosen, i))
		for i := range n {
			current[i] ^= row[i]
		}
	}

	dfs(0, nil)
	return combos
}

func getTargetPart1(line string) []int {
	end := re1.FindStringSubmatch(line)[1]
	target := make([]int, len(end))
	for i := range end {
		if end[i] == '#' {
			target[i] = 1
		} else {
			target[i] = 0
		}
	}
	return target
}

func getTargetPart2(line string) []int {
	var target []int
	for i := range strings.SplitSeq(re3.FindStringSubmatch(line)[1], ",") {
		target = append(target, util.ParseInt(i))
	}
	return target
}

func getButtons(line string, n int) []button {
	var buttons []button
	input := re2.FindAllStringSubmatch(line, -1)
	for _, b := range input {
		button := make(button, n)
		for i := range strings.SplitSeq(b[1], ",") {
			button[util.ParseInt(i)] = 1
		}
		buttons = append(buttons, button)
	}
	return buttons
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

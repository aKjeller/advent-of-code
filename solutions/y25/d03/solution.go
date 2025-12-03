package main

import (
	"fmt"
	"strconv"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "03"

func part1(path string) {
	banks := util.ToStringSlice(path)

	sum := 0
	for _, bank := range banks {
		i, a := findMaxJoltageIndex(bank, 0, len(bank)-1)
		_, b := findMaxJoltageIndex(bank, i+1, len(bank))
		sum += util.Concatenate(a, b)
	}

	fmt.Println("part1: ", sum)
}

func part2(path string) {
	banks := util.ToStringSlice(path)

	sum := 0
	for _, bank := range banks {
		ans := ""
		var joltageIndex, joltage int
		for c := range 12 {
			joltageIndex, joltage = findMaxJoltageIndex(bank, joltageIndex, len(bank)-11+c)
			joltageIndex += 1
			ans += strconv.Itoa(joltage)
		}
		sum += util.ParseInt(ans)
	}

	fmt.Println("part2: ", sum)
}

func findMaxJoltageIndex(bank string, start, end int) (maxIndex int, maxJoltage int) {
	for i := start; i < end; i++ {
		v := util.ParseInt(string(bank[i]))
		if v > maxJoltage {
			maxIndex, maxJoltage = i, v
		}
	}
	return
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

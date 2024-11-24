package main

import (
	"fmt"
	"log"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "02"

func part1(path string) {
	input := util.ToStringSlice(path)
	score := 0
	for _, s := range input {
		a, b := parse(s)
		score += getBonus(b)
		score += getScore(a, b)
	}
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToStringSlice(path)
	score := 0
	for _, s := range input {
		a, x := parse(s)
		var b string
		switch string(x) {
		case "A":
			b = findResult(0, a)
		case "B":
			b = findResult(3, a)
		case "C":
			b = findResult(6, a)
		default:
			log.Fatalf("error, x is: %s", x)
		}
		score += getBonus(b)
		score += getScore(a, b)
	}
	fmt.Printf("part2: %d\n", score)
}

func findResult(res int, hand string) string {
	hands := [3]string{"A", "B", "C"}
	for _, h := range hands {
		score := getScore(hand, h)
		if score == res {
			return h
		}
	}
	log.Fatal("res not found")
	return ""
}

func getBonus(c string) int {
	if c == "A" {
		return 1
	}
	if c == "B" {
		return 2
	}
	return 3
}

func getScore(a, b string) int {
	if a == b {
		return 3
	}
	if a == "A" && b == "B" || a == "B" && b == "C" || a == "C" && b == "A" {
		return 6
	}
	return 0
}

func parse(s string) (string, string) {
	a := string(s[0])
	switch s[2] {
	case 'X':
		return a, "A"
	case 'Y':
		return a, "B"
	case 'Z':
		return a, "C"
	}
	log.Fatal("could not parse")
	return " ", " "
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

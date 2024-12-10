package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "24"
const DAY = "04"

func part1(path string) {
	data := util.ToGrid(path)

	score := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			score += findWord("XMAS", data, i, j)
		}
	}

	fmt.Printf("part1: %d\n", score)
}

func findWord(word string, input [][]uint8, i, j int) int {
	score := 0
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i-n, j-n) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i-n, j) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i-n, j+n) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i, j-n) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i, j+n) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i+n, j-n) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i+n, j) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	for n := 0; n < len(word); n++ {
		if !valid(input, word[n], i+n, j+n) {
			break
		} else if n == len(word)-1 {
			score += 1
		}
	}
	return score
}

func valid(input [][]uint8, c uint8, i, j int) bool {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	return c == input[i][j]
}

func part2(path string) {
	data := util.ToGrid(path)

	score := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			s := ""
			s += string(getChar(data, i, j))
			s += string(getChar(data, i+1, j+1))
			s += string(getChar(data, i+2, j+2))
			s += string(getChar(data, i, j+2))
			s += string(getChar(data, i+1, j+1))
			s += string(getChar(data, i+2, j))
			if s == "MASMAS" {
				score += 1
			}
			if s == "MASSAM" {
				score += 1
			}
			if s == "SAMMAS" {
				score += 1
			}
			if s == "SAMSAM" {
				score += 1
			}
		}
	}
	fmt.Printf("part2: %d\n", score)
}

func getChar(data [][]uint8, i, j int) uint8 {
	defer func() {
		recover()
	}()
	return data[i][j]
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

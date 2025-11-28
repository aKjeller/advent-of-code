package main

import (
	"fmt"
	"slices"

	util "github.com/aKjeller/advent-of-code/utilities/go"
	"github.com/aKjeller/advent-of-code/utilities/go/ds"
)

const YEAR = "22"
const DAY = "05"

func part1(path string) {
	input := util.ToStringSlice(path)
	stacks, moves := getStacksAndMoves(input)

	for _, move := range moves {
		m := util.GetIntsFromString(move)
		for range m[0] {
			item, err := stacks[m[1]].Pop()
			if err != nil {
				panic(err)
			}
			stacks[m[2]].Push(item)
		}
	}

	var res []byte
	for i := range len(stacks) {
		item, err := stacks[i+1].Peek()
		if err != nil {
			panic(err)
		}
		res = append(res, item)
	}

	fmt.Printf("part1: %s\n", res)
}

func part2(path string) {
	input := util.ToStringSlice(path)
	stacks, moves := getStacksAndMoves(input)

	for _, move := range moves {
		m := util.GetIntsFromString(move)
		var boxes []byte
		for range m[0] {
			item, err := stacks[m[1]].Pop()
			if err != nil {
				panic(err)
			}
			boxes = append(boxes, item)
		}
		slices.Reverse(boxes)
		for _, box := range boxes {
			stacks[m[2]].Push(box)
		}
	}

	var res []byte
	for i := range len(stacks) {
		item, err := stacks[i+1].Peek()
		if err != nil {
			panic(err)
		}
		res = append(res, item)
	}

	fmt.Printf("part1: %s\n", res)
}

func getStacksAndMoves(input []string) (map[int]*ds.Stack[byte], []string) {
	stacks := make(map[int]*ds.Stack[byte])

	cut := 0
	for i, line := range input {
		if line == "" {
			cut = i
			break
		}
	}

	for index := cut - 2; index >= 0; index-- {
		boxes := input[index]
		box := 1
		for i := 1; i < len(boxes); i += 4 {
			if _, ok := stacks[box]; !ok {
				stacks[box] = &ds.Stack[byte]{}
			}
			if boxes[i] != ' ' {
				stacks[box].Push(boxes[i])
			}
			box++
		}
	}
	return stacks, input[cut+1:]
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

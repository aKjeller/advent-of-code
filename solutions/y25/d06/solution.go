package main

import (
	"fmt"
	"regexp"
	"unicode"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "25"
const DAY = "06"

var reop = regexp.MustCompile(`\*|\+`)

func part1(path string) {
	input := util.ToStringSlice(path)

	var nums [][]int
	for i := range len(input) - 1 {
		nums = append(nums, util.GetIntsFromString(input[i]))
	}

	sum := 0
	ops := reop.FindAllString(input[len(input)-1], -1)
	for c, op := range ops {
		total := nums[0][c]
		for r := 1; r < len(nums); r++ {
			if op == "*" {
				total *= nums[r][c]
			} else {
				total += nums[r][c]
			}
		}
		sum += total
	}

	fmt.Println("part1: ", sum)
}

type op struct {
	op         string
	start, end int
}

func part2(path string) {
	input := util.ToStringSlice(path)
	for i := range input {
		input[i] += " "
	}

	var ops []op
	opRow := input[len(input)-1]
	end := len(opRow) - 1
	for i := end; i >= 0; i-- {
		v := opRow[i]
		if v != ' ' {
			ops = append(ops, op{
				op:    string(v),
				start: i,
				end:   end,
			})
			end = i - 1
		}
	}

	sum := 0
	for _, op := range ops {
		total := 0
		for col := op.start; col < op.end; col++ {
			var num []byte
			for row := range len(input) - 1 {
				digit := input[row][col]
				if unicode.IsDigit(rune(digit)) {
					num = append(num, digit)
				}
			}
			number := util.ParseInt(string(num))
			if op.op == "*" {
				if total == 0 {
					total = number
				} else {
					total *= number
				}
			} else {
				total += number
			}
		}
		sum += total
	}

	fmt.Println("part2: ", sum)
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

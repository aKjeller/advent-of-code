package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"math"
	"strconv"
	"strings"
)

const YEAR = "24"
const DAY = "17"

func part1(path string) {
	nums := util.GetIntsFromString(util.ToString(path))
	c := newComputer(nums[3:], nums[0], nums[1], nums[2])
	c.run()
	fmt.Printf("part1: %s\n", outputToString(c.output))
}

func part2(path string) {
	nums := util.GetIntsFromString(util.ToString(path))
	c := newComputer(nums[3:], 0, nums[1], nums[2])
	a := 0
	for i := len(c.program) - 1; i >= 0; i-- {
		a = a * 8
		for {
			c.setRegister(a, nums[1], nums[2])
			c.run()
			if util.CompareSlices(c.output, c.program[i:]) {
				break
			}
			a++
		}
	}

	fmt.Printf("part2: %d\n", a)
}

type computer struct {
	program  []int
	register map[uint8]int
	output   []int
}

func newComputer(program []int, a, b, c int) computer {
	register := make(map[uint8]int)
	comp := computer{program: program, register: register}
	comp.setRegister(a, b, c)
	return comp
}

func (comp *computer) setRegister(a, b, c int) {
	comp.register['A'] = a
	comp.register['B'] = b
	comp.register['C'] = c
}

func (comp *computer) run() {
	comp.output = []int{}
	counter := 0
	for counter < len(comp.program)-1 {
		operand := comp.program[counter+1]
		switch comp.program[counter] {
		case 0:
			adv(comp.register, operand)
		case 1:
			bxl(comp.register, operand)
		case 2:
			bst(comp.register, operand)
		case 3:
			counter = jnz(comp.register, operand, counter)
			continue
		case 4:
			bxc(comp.register)
		case 5:
			comp.output = append(comp.output, out(comp.register, operand))
		case 6:
			bdv(comp.register, operand)
		case 7:
			cdv(comp.register, operand)
		}
		counter += 2
	}
}

func getComboOperand(register map[uint8]int, o int) int {
	switch {
	case o >= 0 && o <= 3:
		return o
	case o == 4:
		return register['A']
	case o == 5:
		return register['B']
	case o == 6:
		return register['C']
	default:
		panic(fmt.Sprintf("Invalid combo operand: %d", o))
	}
}

func adv(register map[uint8]int, operand int) {
	operand = getComboOperand(register, operand)
	register['A'] = register['A'] / int(math.Pow(2, float64(operand)))
}

func bxl(register map[uint8]int, operand int) {
	register['B'] = register['B'] ^ operand
}

func bst(register map[uint8]int, operand int) {
	operand = getComboOperand(register, operand)
	register['B'] = operand % 8
}

func jnz(register map[uint8]int, operand, counter int) int {
	if register['A'] > 0 {
		return operand
	}
	return counter + 2
}

func bxc(register map[uint8]int) {
	register['B'] = register['B'] ^ register['C']
}

func out(register map[uint8]int, operand int) int {
	operand = getComboOperand(register, operand)
	return operand % 8
}

func bdv(register map[uint8]int, operand int) {
	operand = getComboOperand(register, operand)
	register['B'] = register['A'] / int(math.Pow(2, float64(operand)))
}

func cdv(register map[uint8]int, operand int) {
	operand = getComboOperand(register, operand)
	register['C'] = register['A'] / int(math.Pow(2, float64(operand)))
}

func outputToString(output []int) string {
	var tmp []string
	for _, o := range output {
		tmp = append(tmp, strconv.Itoa(o))
	}
	return strings.Join(tmp, ",")
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

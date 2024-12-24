package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"sort"
	"strings"
)

const YEAR = "24"
const DAY = "24"

func part1(path string) {
	wires, gates, zs := parse(path)

	for !done(wires, zs) {
		update(wires, gates)
	}

	number := 0
	for _, bit := range zs {
		number = (number << 1) | wires[bit]
	}

	fmt.Printf("part1: %d\n", number)
}

func part2(path string) {
	_, gates, _ := parse(path)

	var faulty []string
	for _, g := range gates {
		// must be an XOR before z
		if strings.HasPrefix(g.out, "z") && g.opId != XOR && g.out != "z45" {
			faulty = append(faulty, g.out)
		}

		// XOR must end with Z or start with xy
		if !strings.HasPrefix(g.out, "z") {
			if !strings.Contains("xy", g.in1[:1]) && !strings.Contains("xy", g.in2[:1]) && g.opId == XOR {
				faulty = append(faulty, g.out)
			}
		}

		// XOR starting with xy must out to another XOR
		if strings.Contains("xy", g.in1[:1]) && strings.Contains("xy", g.in2[:1]) && g.opId == XOR {
			correct := false
			for _, g2 := range gates {
				if g2.in1 == g.out || g2.in2 == g.out && g2.opId == XOR {
					correct = true
				}
			}
			if !correct && !strings.Contains(g.in1, "00") && !strings.Contains(g.in2, "00") {
				faulty = append(faulty, g.out)
			}
		}

		// AND must have OR as out
		if g.opId == AND {
			correct := false
			for _, g2 := range gates {
				if g2.opId == OR && (g2.in1 == g.out || g2.in2 == g.out) {
					correct = true
				}
			}
			if !correct && !strings.Contains(g.in1, "00") && !strings.Contains(g.in2, "00") && !strings.HasPrefix(g.out, "z") {
				faulty = append(faulty, g.out)
			}
		}
	}

	sort.Strings(faulty)
	fmt.Printf("part2: %s\n", strings.Join(faulty, ","))
}

func update(wires map[string]int, gates []gate) {
	newWire := make(map[string]int)
	for _, g := range gates {
		if g.hasInput(wires) {
			newWire[g.out] = g.op(wires[g.in1], wires[g.in2])
		}
	}
	for v, k := range newWire {
		wires[v] = k
	}
}

func done(wires map[string]int, zs []string) bool {
	for _, z := range zs {
		if _, ok := wires[z]; !ok {
			return false
		}
	}
	return true
}

type opId string

var (
	AND opId = "AND"
	OR  opId = "OR"
	XOR opId = "XOR"
)

type gate struct {
	in1, in2 string
	op       func(in1, in2 int) int
	opId     opId
	out      string
}

func (g gate) hasInput(wires map[string]int) bool {
	if _, ok := wires[g.in1]; !ok {
		return false
	}
	if _, ok := wires[g.in2]; !ok {
		return false
	}
	return true
}

func and(in1, in2 int) int {
	if in1 == 1 && in2 == 1 {
		return 1
	}
	return 0
}

func or(in1, in2 int) int {
	if in1 == 1 || in2 == 1 {
		return 1
	}
	return 0
}

func xor(in1, in2 int) int {
	if in1 != in2 {
		return 1
	}
	return 0
}

func parse(path string) (map[string]int, []gate, []string) {
	wires := make(map[string]int)
	var gates []gate
	var zs []string

	firstHalf := true
	for _, r := range util.ToStringSlice(path) {
		if r == "" {
			firstHalf = false
		} else if firstHalf {
			split := strings.Split(r, ": ")
			wires[split[0]] = util.ParseInt(split[1])
		} else {
			split := strings.Split(r, " ")
			if strings.HasPrefix(split[4], "z") {
				zs = append(zs, split[4])
			}
			switch split[1] {
			case "AND":
				gates = append(gates, gate{split[0], split[2], and, AND, split[4]})
			case "OR":
				gates = append(gates, gate{split[0], split[2], or, OR, split[4]})
			case "XOR":
				gates = append(gates, gate{split[0], split[2], xor, XOR, split[4]})
			}
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(zs)))

	return wires, gates, zs
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

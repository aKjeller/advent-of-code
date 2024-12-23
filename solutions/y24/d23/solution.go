package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"maps"
	"slices"
	"sort"
	"strings"
)

const YEAR = "24"
const DAY = "23"

func part1(path string) {
	computers := makeComputers(path)

	unique := make(map[string]bool)
	for c1, _ := range computers {
		if strings.HasPrefix(c1, "t") {
			for _, c2 := range computers[c1] {
				for _, c3 := range computers[c2] {
					for _, c4 := range computers[c3] {
						if c1 == c4 {
							conn := []string{c1, c2, c3}
							sort.Strings(conn)
							unique[strings.Join(conn, "")] = true
						}
					}
				}
			}
		}
	}

	fmt.Printf("part1: %d\n", len(unique))
}

func part2(path string) {
	computers := makeComputers(path)

	var largestClique []string
	for c, _ := range computers {
		clique := getLargestClique(c, c, computers, make(map[string]bool))
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}

	fmt.Printf("part2: %s\n", strings.Join(largestClique, ","))
}

func makeComputers(path string) map[string][]string {
	computers := make(map[string][]string)
	for _, r := range util.ToStringSlice(path) {
		c := strings.Split(r, "-")
		computers[c[0]] = append(computers[c[0]], c[1])
		computers[c[1]] = append(computers[c[1]], c[0])
	}
	return computers
}

func getLargestClique(start, current string, computers map[string][]string, currentClique map[string]bool) []string {
	if start == current && currentClique[start] {
		return slices.Sorted(maps.Keys(currentClique))
	}

	for v, _ := range currentClique {
		if !slices.Contains(computers[v], current) {
			return []string{}
		}
	}

	currentClique[current] = true

	var largestClique []string
	for _, c := range computers[current] {
		clique := getLargestClique(start, c, computers, currentClique)
		if len(clique) > len(largestClique) {
			largestClique = clique
		}
	}
	return largestClique
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

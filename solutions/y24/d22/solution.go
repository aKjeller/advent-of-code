package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"math"
)

const YEAR = "24"
const DAY = "22"

func part1(path string) {
	score := 0
	for _, r := range util.ToIntSlice(path) {
		secret := *r
		for range 2000 {
			secret = nextSecret(secret)
		}
		score += secret
	}
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	allPrices := make(map[sequence]int)
	for _, r := range util.ToIntSlice(path) {
		secret := *r
		prices := make(map[sequence]int)

		seq := sequence{}
		price, delta := secret%10, 0
		for i := range 2000 {
			secret = nextSecret(secret)
			newPrice := secret % 10
			delta = newPrice - price
			seq.push(delta)
			if _, ok := prices[seq]; !ok && i >= 3 {
				prices[seq] = newPrice
			}
			price = newPrice
		}

		for k, v := range prices {
			allPrices[k] += v
		}
	}

	bananas := math.MinInt
	for _, b := range allPrices {
		bananas = max(bananas, b)
	}

	fmt.Printf("part2: %d\n", bananas)
}

type sequence [4]int

func (s *sequence) push(x int) {
	s[0] = s[1]
	s[1] = s[2]
	s[2] = s[3]
	s[3] = x
}

func nextSecret(secret int) int {
	secret = (secret ^ (secret * 64)) % 16777216
	secret = (secret ^ (secret / 32)) % 16777216
	secret = (secret ^ (secret * 2048)) % 16777216
	return secret
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

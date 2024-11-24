package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func InputPath(year, day string) string {
	return fmt.Sprintf("solutions/y%s/d%s/input.txt", year, day)
}

func ExamplePath(year, day string) string {
	return fmt.Sprintf("solutions/y%s/d%s/example.txt", year, day)
}

func ToIntSlice(path string) []*int {
	lines := ToStringSlice(path)

	var nums []*int
	for _, s := range lines {
		if s == "" {
			nums = append(nums, nil)
		} else {
			i, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			nums = append(nums, &i)
		}
	}
	return nums
}

func ToStringSlice(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var s []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return s
}

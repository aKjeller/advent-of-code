package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func InputPath(year, day string) string {
	return fmt.Sprintf("solutions/y%s/d%s/input.txt", year, day)
}

func ExamplePath(year, day string) string {
	return fmt.Sprintf("solutions/y%s/d%s/example.txt", year, day)
}

func GetIntsFromString(s string) []int {
	var nums []int

	var tmp string
	for _, c := range s {
		if unicode.IsDigit(c) {
			tmp += string(c)
		} else if tmp != "" {
			num := ParseInt(tmp)
			nums = append(nums, num)
			tmp = ""
		}
	}

	if tmp != "" {
		num := ParseInt(tmp)
		nums = append(nums, num)
	}

	return nums
}

func ToIntSlice(path string) []*int {
	lines := ToStringSlice(path)

	var nums []*int
	for _, s := range lines {
		if s == "" {
			nums = append(nums, nil)
		} else {
			i := ParseInt(s)
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

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

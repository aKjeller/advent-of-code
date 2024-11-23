package util

import (
	"bufio"
	"fmt"
	"os"
)

func InputPath(year, day string) string {
	return fmt.Sprintf("solutions/y%s/d%s/input.txt", year, day)
}

func ExamplePath(year, day string) string {
	return fmt.Sprintf("solutions/y%s/d%s/example.txt", year, day)
}

func ToStringSlice(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var s []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	return s
}

func ToString(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(b)
}

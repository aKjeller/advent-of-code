package util

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func GetIntsFromStringWithNegative(s string) []int {
	r := regexp.MustCompile(`-?\d+`)
	var nums []int
	for _, match := range r.FindAllString(s, -1) {
		nums = append(nums, ParseInt(match))
	}
	return nums
}

func GetFloatsFromString(s string) []float64 {
	var floats []float64
	ints := GetIntsFromString(s)
	for _, i := range ints {
		floats = append(floats, float64(i))
	}
	return floats
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

func ToGrid(path string) (m [][]uint8) {
	input := ToStringSlice(path)
	for i := 0; i < len(input); i++ {
		var row []uint8
		for j := 0; j < len(input[i]); j++ {
			row = append(row, input[i][j])
		}
		m = append(m, row)
	}
	return m
}

func ToString(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(b), "\r\n", "\n", -1)
}

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return int(i)
}

// Abs returns the absolute value of i.
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// RemoveElement returns a new slice without the element at index i.
func RemoveElement[T any](src []T, i int) []T {
	output := make([]T, 0)
	output = append(output, src[:i]...)
	return append(output, src[i+1:]...)
}

// DeepCopy returns a deep copy of a 2d-slice
func DeepCopy[T any](src [][]T) [][]T {
	dst := make([][]T, len(src))
	for i := range src {
		dst[i] = make([]T, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

// Concatenate concatenates two integers, 1337 + 1337 = 13371337
func Concatenate(a, b int) int {
	mul := 1
	for tmp := b; tmp > 0; tmp /= 10 {
		mul *= 10
	}
	return a*mul + b
}

// Gcd calculates gcd
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// CompareSlices returns true if the slices are equal
func CompareSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

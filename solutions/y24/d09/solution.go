package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"slices"
)

const YEAR = "24"
const DAY = "09"

const EMPTY = -1

func part1(path string) {
	data := parse1(util.ToString(path))
	j := 0
	for i := len(data) - 1; i-1 > j; i-- {
		for data[j] != EMPTY {
			j++
		}
		data[j], data[i] = data[i], EMPTY
	}
	fmt.Printf("part1: %d\n", checksum1(data))
}

func parse1(data string) []int {
	var arr []int
	id := 0
	isFile := true
	for i, _ := range data {
		x := EMPTY
		if isFile {
			x = id
			id++
		}
		for range data[i] - '0' {
			arr = append(arr, x)
		}
		isFile = !isFile
	}
	return arr
}

func checksum1(data []int) int {
	sum := 0
	for i := 0; i < len(data); i++ {
		if data[i] != EMPTY {
			sum += i * data[i]
		}
	}
	return sum
}

type file struct {
	id   int
	size int
}

func part2(path string) {
	input := util.ToString(path)
	files := parse2(input)
	for i := len(files) - 1; i > -1; i-- {
		if files[i].id != EMPTY {
			moveFileLeft(files, i)
		}
	}

	fmt.Printf("part2: %d\n", checksum2(files))
}

func moveFileLeft(files []file, i int) {
	for j := 0; j < i; j++ {
		if files[j].id == EMPTY && files[j].size >= files[i].size {
			files[j], files[i] = files[i], files[j]
			if files[i].size > files[j].size {
				newFile := file{id: EMPTY, size: files[i].size - files[j].size}
				files[i].size = files[j].size
				slices.Insert(files, j+1, newFile)
			}
			return
		}
	}
}

func parse2(data string) []file {
	var files []file
	id, isFile := 0, true
	for i, _ := range data {
		f := file{id: EMPTY, size: int(data[i] - '0')}
		if isFile {
			f.id = id
			id++
		}
		files = append(files, f)
		isFile = !isFile
	}
	return files
}

func checksum2(files []file) int {
	sum, pos := 0, 0
	for _, f := range files {
		for range f.size {
			if f.id != EMPTY {
				sum += f.id * pos
			}
			pos++
		}
	}
	return sum
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

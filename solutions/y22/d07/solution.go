package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/aKjeller/advent-of-code/utilities/go"
)

const YEAR = "22"
const DAY = "07"

type file struct {
	name   string
	parent *file

	size     int
	subFiles map[string]*file
}

func part1(path string) {
	input := util.ToStringSlice(path)

	_, dirs := getFs(input)
	score := 0
	for _, dir := range dirs {
		size := dirSize(dir)
		if size <= 100000 {
			score += size
		}

	}
	fmt.Printf("part1: %d\n", score)
}

func part2(path string) {
	input := util.ToStringSlice(path)

	root, dirs := getFs(input)
	used := dirSize(root)

	minSize := used - (70000000 - 30000000)

	best := int(^uint(0) >> 1)
	for _, dir := range dirs {
		size := dirSize(dir)
		if size >= minSize {
			best = min(best, size)
		}

	}
	fmt.Printf("part2: %d\n", best)
}

func getFs(input []string) (*file, []*file) {
	root := &file{
		name:     "/",
		size:     -1,
		subFiles: map[string]*file{},
	}

	dirs := []*file{}

	ptr := root
	for _, line := range input {
		// handle ls
		if strings.HasPrefix(line, "$ ls") {
			continue
		}

		// handle cd
		if strings.HasPrefix(line, "$") {
			cdTarget, _ := strings.CutPrefix(line, "$ cd ")
			switch cdTarget {
			case "/":
				ptr = root
			case "..":
				ptr = ptr.parent
			default:
				ptr = ptr.subFiles[cdTarget]
			}
			continue
		}

		// handle ls output
		if strings.HasPrefix(line, "dir") {
			folder, _ := strings.CutPrefix(line, "dir ")
			if _, ok := ptr.subFiles[folder]; !ok {
				dir := &file{
					name:     folder,
					parent:   ptr,
					size:     -1,
					subFiles: map[string]*file{},
				}
				ptr.subFiles[folder] = dir
				dirs = append(dirs, dir)
			}
			continue
		}

		parts := strings.Split(line, " ")
		name := parts[1]
		if _, ok := ptr.subFiles[name]; !ok {
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			ptr.subFiles[name] = &file{
				name:   name,
				parent: ptr,
				size:   size,
			}
		}
	}

	return root, dirs
}

func dirSize(dir *file) int {
	size := 0
	for _, v := range dir.subFiles {
		if v.size > 0 {
			size += v.size
		} else {
			size += dirSize(v)
		}
	}
	return size
}

func main() {
	// part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	// part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

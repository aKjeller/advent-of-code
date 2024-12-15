package main

import (
	"fmt"
	util "github.com/aKjeller/advent-of-code/utilities/go"
	"strings"
)

const YEAR = "24"
const DAY = "15"

func part1(path string) {
	input := strings.Split(util.ToString(path), "\n\n")

	inputMap := strings.Split(input[0], "\n")
	var robot object
	var m [][]mover
	for i := 0; i < len(inputMap); i++ {
		var row []mover
		for j := 0; j < len(inputMap[i]); j++ {
			if inputMap[i][j] == '#' {
				row = append(row, &object{x: i, y: j, movable: false, t: "#"})
			} else if inputMap[i][j] == 'O' {
				row = append(row, &object{x: i, y: j, movable: true, t: "0"})
			} else if inputMap[i][j] == '@' {
				robot = object{x: i, y: j, movable: true, t: "@"}
				row = append(row, &robot)
			} else {
				row = append(row, nil)
			}
		}
		m = append(m, row)
	}

	moveRobot(m, input[1], &robot)

	fmt.Printf("part1: %d\n", getScore(m))
}

func part2(path string) {
	input := strings.Split(util.ToString(path), "\n\n")

	inputMap := strings.Split(input[0], "\n")
	var robot object
	var m [][]mover
	for i := 0; i < len(inputMap); i++ {
		var row []mover
		for j := 0; j < len(inputMap[i]); j++ {
			if inputMap[i][j] == '#' {
				row = append(row, &object{x: i, y: j * 2, movable: false, t: "#"})
				row = append(row, &object{x: i, y: j*2 + 1, movable: false, t: "#"})
			} else if inputMap[i][j] == 'O' {
				row = append(row, &wideObject{object{x: i, y: j * 2, movable: true, t: "["}})
				row = append(row, &wideObject{object{x: i, y: j*2 + 1, movable: true, t: "]"}})
			} else if inputMap[i][j] == '@' {
				robot = object{x: i, y: j * 2, movable: true, t: "@"}
				row = append(row, &robot)
				row = append(row, nil)
			} else {
				row = append(row, nil)
				row = append(row, nil)
			}
		}
		m = append(m, row)
	}

	moveRobot(m, input[1], &robot)

	fmt.Printf("part2: %d\n", getScore(m))
}

func moveRobot(m [][]mover, dirs string, robot *object) {
	for _, d := range dirs {
		if d == '^' {
			moveIfPossible(robot, m, UP)
		}
		if d == 'v' {
			moveIfPossible(robot, m, DOWN)
		}
		if d == '<' {
			moveIfPossible(robot, m, LEFT)
		}
		if d == '>' {
			moveIfPossible(robot, m, RIGHT)
		}
	}
}

func moveIfPossible(robot *object, m [][]mover, d dir) {
	if robot.canMove(m, d) {
		robot.move(m, d)
	}
}

func getScore(m [][]mover) int {
	score := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != nil && (m[i][j].getType() == "[" || m[i][j].getType() == "0") {
				score += 100*i + j
			}
		}
	}
	return score
}

func main() {
	part1(util.ExamplePath(YEAR, DAY))
	part1(util.InputPath(YEAR, DAY))
	part2(util.ExamplePath(YEAR, DAY))
	part2(util.InputPath(YEAR, DAY))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Vec struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	grid := utils.ExtractGrid(lines)
	robot := Vec{}
	boxes := []*Vec{}
	walls := []*Vec{}

	for y, _ := range grid {
		for x, char := range grid[y] {
			if char == '@' { // Robot
				robot = Vec{x, y}
			} else if char == 'O' { // Box
				box := Vec{x, y}
				boxes = append(boxes, &box)
			} else if char == '#' { // Wall
				wall := Vec{x, y}
				walls = append(walls, &wall)
			}
		}
	}

	moves := []Vec{}

	for _, line := range utils.ReadLines(scanner) {
		for _, char := range line {
			if char == '^' { // Up
				moves = append(moves, Vec{0, -1})
			} else if char == '>' { // Right
				moves = append(moves, Vec{1, 0})
			} else if char == 'v' { // Down
				moves = append(moves, Vec{0, 1})
			} else if char == '<' { // Left
				moves = append(moves, Vec{-1, 0})
			}
		}
	}

	for _, move := range moves {
		walk(&grid, &robot, &move, &boxes, &walls)
	}

	for _, box := range boxes {
		res += box.x + 100*box.y
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func walk(grid *[][]rune, pos, move *Vec, boxes, walls *[]*Vec) bool {
	if hasObject(Vec{pos.x + move.x, pos.y + move.y}, walls) != nil {
		return false
	}
	if box := hasObject(Vec{pos.x + move.x, pos.y + move.y}, boxes); box != nil {
		if walk(grid, box, move, boxes, walls) {
			pos.x += move.x
			pos.y += move.y
			return true
		}
		return false
	} else {
		pos.x += move.x
		pos.y += move.y
	}
	return true
}

func hasObject(vec Vec, objects *[]*Vec) *Vec {
	for _, object := range *objects {
		if vec == *object {
			return object
		}
	}
	return nil
}

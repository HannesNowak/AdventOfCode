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

type Box struct {
	l, r Vec
}

func (l Vec) add(move Vec) Vec {
	nextVec := Vec{l.x, l.y}
	nextVec.x += move.x
	nextVec.y += move.y
	return nextVec
}

type Maze struct {
	robot Vec
	parts map[Vec]Box
	boxes map[Box]struct{}
	walls map[Vec]struct{}
}

func canPush(maze *Maze, pos, move Vec) bool {
	canMove := true
	box := maze.parts[pos]
	nextL := box.l.add(move)
	nextR := box.r.add(move)

	_, wallL := maze.walls[nextL]
	_, wallR := maze.walls[nextR]
	if wallL || wallR {
		return false
	}

	if move.x < 0 { // Left
		_, left := maze.parts[nextL]
		if left {
			canMove = canPush(maze, nextL, move)
		}
	} else if move.x > 0 { // Right
		_, right := maze.parts[nextR]
		if right {
			canMove = canPush(maze, nextR, move)
		}
	} else { // Up or Down
		boxL, left := maze.parts[nextL]
		boxR, right := maze.parts[nextR]

		if left {
			canMove = canMove && canPush(maze, nextL, move)
		}
		if right && boxL != boxR {
			canMove = canMove && canPush(maze, nextR, move)
		}
	}
	return canMove
}

func push(maze *Maze, pos, move Vec) {
	box := maze.parts[pos]
	boxL, boxR := box.l, box.r
	nextL, nextR := boxL.add(move), boxR.add(move)

	if move.x < 0 {
		_, left := maze.parts[nextL]
		if left {
			push(maze, nextL, move)
		}

		delete(maze.boxes, box)
		delete(maze.parts, boxL)
		delete(maze.parts, boxR)

		box.r = box.l
		box.l = nextL
		maze.boxes[box] = struct{}{}
		maze.parts[boxL] = box
		maze.parts[nextL] = box
	} else if move.x > 0 {
		_, right := maze.parts[nextR]
		if right {
			push(maze, nextR, move)
		}
		delete(maze.boxes, box)
		delete(maze.parts, boxL)
		delete(maze.parts, boxR)

		box.l = box.r
		box.r = nextR

		maze.boxes[box] = struct{}{}
		maze.parts[boxR] = box
		maze.parts[nextR] = box
	} else {
		bbL, left := maze.parts[nextL]
		bbR, right := maze.parts[nextR]

		if left {
			push(maze, nextL, move)
		}
		if right && bbL != bbR {
			push(maze, nextR, move)
		}
		delete(maze.boxes, box)
		delete(maze.parts, boxL)
		delete(maze.parts, boxR)

		box.l = nextL
		box.r = nextR

		maze.boxes[box] = struct{}{}
		maze.parts[nextL] = box
		maze.parts[nextR] = box
	}
}

func walk(maze *Maze, move Vec) {
	nextPos := maze.robot.add(move)

	_, wall := maze.walls[nextPos]
	if wall {
		return
	}

	_, part := maze.parts[nextPos]
	if part {
		if canPush(maze, nextPos, move) {
			push(maze, nextPos, move)
		} else {
			return
		}
	}
	maze.robot = nextPos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	initGrid := utils.ExtractGrid(lines)
	grid := make([][]rune, len(initGrid))
	for y := range initGrid {
		for x := range initGrid[y] {
			grid[y] = append(grid[y], initGrid[y][x], '.')
		}
	}

	maze := Maze{}
	maze.robot = Vec{}
	maze.parts = map[Vec]Box{}
	maze.boxes = map[Box]struct{}{}
	maze.walls = map[Vec]struct{}{}

	for y := range grid {
		for x, char := range grid[y] {
			pos := Vec{x, y}
			nextPos := pos.add(Vec{1, 0})
			if char == '@' { // Robot
				maze.robot = pos
			} else if char == 'O' { // Box
				box := Box{pos, nextPos}
				maze.boxes[box] = struct{}{}
				maze.parts[pos] = box
				maze.parts[nextPos] = box
			} else if char == '#' { // wall
				maze.walls[pos] = struct{}{}
				maze.walls[nextPos] = struct{}{}
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
		walk(&maze, move)
	}
	utils.PrintGrid(updateGrid(grid, maze))

	for box := range maze.boxes {
		res += box.l.x + 100*box.l.y
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func updateGrid(grid [][]rune, maze Maze) [][]rune {
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = '.'
		}
	}
	for wall := range maze.walls {
		grid[wall.y][wall.x] = '#'
	}
	for box := range maze.boxes {
		grid[box.l.y][box.l.x] = '['
		grid[box.r.y][box.r.x] = ']'
	}
	grid[maze.robot.y][maze.robot.x] = '@'
	return grid
}

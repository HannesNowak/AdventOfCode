package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Guard struct {
	x    int
	y    int
	dirX int
	dirY int
}

func (g *Guard) walk() (int, int) {
	return g.x + g.dirX, g.y + g.dirY
}

func (g *Guard) turn() {
	g.dirX, g.dirY = -g.dirY, g.dirX
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	grid := utils.ExtractGrid(lines)

	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				guard := Guard{x, y, 0, -1}
				res = walk(grid, guard)
				break
			}
		}
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
	fmt.Println(cnt)
}

func outOfBounds(grid [][]rune, x, y int) bool {
	return y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y])
}

func walk(grid [][]rune, guard Guard) int {
	var loops int
	unobsGuard := guard
	obstructions := [][2]int{}

	for !outOfBounds(grid, guard.x, guard.y) {
		nextX, nextY := unobsGuard.walk()
		if outOfBounds(grid, nextX, nextY) {
			break
		}
		if grid[nextY][nextX] == '#' {
			unobsGuard.turn()
			continue
		} else if !slices.Contains(obstructions, [2]int{nextX, nextY}) {
			grid[nextY][nextX] = '#'
			if checkLoop(grid, guard) {
				loops++
				obstructions = append(obstructions, [2]int{nextX, nextY})
			}
			grid[nextY][nextX] = '.'
		}
		unobsGuard.x, unobsGuard.y = nextX, nextY
	}
	return loops
}

func checkLoop(grid [][]rune, guard Guard) bool {
	walls := [][4]int{}
	loopGuard := guard

	for !outOfBounds(grid, loopGuard.x, loopGuard.y) {
		nextX, nextY := loopGuard.walk()
		if outOfBounds(grid, nextX, nextY) {
			break
		}
		if grid[nextY][nextX] == '#' {
			wall := [4]int{nextX, nextY, loopGuard.dirX, loopGuard.dirY}
			if slices.Contains(walls, wall) {
				return true
			}
			walls = append(walls, wall)
			loopGuard.turn()
		} else {
			loopGuard.x, loopGuard.y = nextX, nextY
		}
	}
	return false
}

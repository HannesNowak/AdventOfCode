package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	grid := utils.ExtractGrid(lines)

	for y, row := range grid {
		for x, cell := range row {
			if cell == '^' {
				res = walk(grid, x, y-1, 0, -1, 1)
				break
			}
		}
	}
	utils.PrintGrid(grid)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func walk(grid [][]rune, x, y, dirX, dirY, fields int) int {
	x += dirX
	y += dirY
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
		return fields
	}
	if grid[y][x] == '#' {
		x -= dirX
		y -= dirY
		dirX, dirY = -dirY, dirX
	} else {
		if grid[y][x] != 'X' {
			fields++
			grid[y][x] = 'X'
		}
	}
	return walk(grid, x, y, dirX, dirY, fields)
}

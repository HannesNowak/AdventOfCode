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

	search := []rune("MAS")

	grid := utils.ExtractGrid(lines)

	directions := [][2]int{
		{1, 1},   // down-right
		{-1, 1},  // down-left
		{-1, -1}, // up-left
		{1, -1},  // up-right
	}

	for y, _ := range grid {
		for x, _ := range grid[y] {
			for _, dir := range directions {
				if searchWord(grid, search, x, y, dir) {
					if searchWord(grid, search, x, y+(len(search)-1)*dir[0], [2]int{-dir[0], dir[1]}) ||
						searchWord(grid, search, x+(len(search)-1)*dir[1], y, [2]int{dir[0], -dir[1]}) {
						res++
					}
				}
			}
		}
	}

	res = res / 2

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func searchWord(grid [][]rune, search []rune, row int, col int, direction [2]int) bool {
	for idx := 0; idx < len(search); idx++ {
		x := col + idx*direction[0]
		y := row + idx*direction[1]
		if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
			return false
		}
		if search[idx] != grid[y][x] {
			return false
		}
	}
	return true
}

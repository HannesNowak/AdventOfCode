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
	satellites := map[rune][][2]int{}

	for y, _ := range grid {
		for x, _ := range grid[y] {
			char := grid[y][x]
			if char != '.' {
				satellites[char] = append(satellites[char], [2]int{x, y})
			}
		}
	}

	for _, freq := range satellites {
		res += frequencyAntinodes(grid, freq)
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func frequencyAntinodes(grid [][]rune, satellites [][2]int) int {
	antinodes := 0
	for i, sat1 := range satellites {
		for j, sat2 := range satellites {
			if i < j {
				antinodes += satsAntinode(grid, sat1, sat2)
			}
		}
	}
	return antinodes
}

func satsAntinode(grid [][]rune, sat1 [2]int, sat2 [2]int) int {
	count := 0
	dx, dy := sat2[0]-sat1[0], sat2[1]-sat1[1]
	pos1X, pos1Y := sat1[0]-dx, sat1[1]-dy
	pos2X, pos2Y := sat2[0]+dx, sat2[1]+dy

	if !utils.OutOfGrid(grid, pos1X, pos1Y) && grid[pos1Y][pos1X] != '#' {
		grid[pos1Y][pos1X] = '#'
		count++
	}

	if !utils.OutOfGrid(grid, pos2X, pos2Y) && grid[pos2Y][pos2X] != '#' {
		grid[pos2Y][pos2X] = '#'
		count++
	}

	return count
}

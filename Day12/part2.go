package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Loc struct {
	x, y int
}

type Plot struct {
	plant                 rune
	area, fences, corners int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	grid := map[Loc]rune{}

	for y, _ := range utils.ExtractGrid(lines) {
		for x, plant := range utils.ExtractGrid(lines)[y] {
			grid[Loc{x, y}] = plant
		}
	}

	plots := []Plot{}
	visited := map[Loc]bool{}

	for field := range grid {
		if visited[field] {
			continue
		}
		visited[field] = true

		var plot Plot
		plot = visitField(grid, &visited, field)
		plots = append(plots, plot)
	}

	for _, plot := range plots {
		res += plot.area * plot.corners
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func visitField(grid map[Loc]rune, visited *map[Loc]bool, field Loc) Plot {
	// create a new plot starting from the field
	plot := Plot{plant: grid[field], area: 1, fences: 0, corners: 0}
	queue := []Loc{field}

	for len(queue) > 0 {
		curr := queue[0]

		// checks the neighbors of the current field
		for j := -1; j <= 1; j++ {
			for i := -1; i <= 1; i++ {
				if i == j || i+j%2 == 0 {
					// skips diagonal directions
					continue
				}

				neighbor := Loc{curr.x + i, curr.y + j}
				if grid[neighbor] != grid[curr] {
					plot.fences++ // neighbor has a different plant
					turn := Loc{curr.x - j, curr.y + i}
					diag := Loc{turn.x + i, turn.y + j}
					if grid[turn] != grid[curr] || grid[diag] == grid[curr] {
						// next neighbor (turn) always has a different plant
						// or the diagonal field (diag) has the same plant
						plot.corners++
					}
				} else if !(*visited)[neighbor] {
					(*visited)[neighbor] = true
					queue = append(queue, neighbor)
					plot.area++
				}
			}
		}
		// dequeue the current field
		queue = queue[1:]
	}
	return plot
}

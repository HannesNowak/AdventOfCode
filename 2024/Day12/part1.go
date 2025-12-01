package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Loc struct {
	x, y int
}

type Plot struct {
	plant  rune
	fences int
	area   []Loc
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	grid := utils.ExtractGrid(lines)

	plots := []Plot{}

	for y, _ := range grid {
		for x, _ := range grid[y] {
			plots = evalPlot(grid, y, x, plots)
		}
	}

	for i := 0; i < 2; i++ {
		plots = combinePlots(plots)
	}

	checkMap := map[rune]int{}

	for _, plot := range plots {
		res += len(plot.area) * plot.fences
		checkMap[plot.plant] += len(plot.area)
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalPlot(grid [][]rune, x, y int, plots []Plot) []Plot {
	plant := grid[y][x]
	idx := -1
	for i, p := range plots {
		if p.plant == plant && slices.Contains(p.area, Loc{x, y}) {
			idx = i
			break
		}
	}
	if idx < 0 {
		plot := Plot{}
		plot.plant = plant
		plot.area = append(plot.area, Loc{x, y})
		idx = len(plots)
		plots = append(plots, plot)
	}
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i != j && i+j%2 != 0 {
				nextX, nextY := x+i, y+j
				if utils.OutOfGrid(grid, nextX, nextY) || grid[nextY][nextX] != plant {
					plots[idx].fences++
				} else {
					plots[idx].area = utils.AppendUnique(plots[idx].area, Loc{nextX, nextY})
				}
			}
		}
	}
	return plots
}

func combinePlots(plots []Plot) []Plot {
	for i := 0; i < len(plots); i++ {
		for j := i + 1; j < len(plots); j++ {
			if plots[i].plant == plots[j].plant {
				intersect := false
				for _, loc := range plots[i].area {
					if slices.Contains(plots[j].area, loc) {
						intersect = true
						break
					}
				}
				if !intersect {
					continue
				}
				plots[i].area = utils.AppendUnique(plots[i].area, plots[j].area...)
				plots[i].fences += plots[j].fences
				plots = utils.Remove(plots, j)
				j--
			}
		}
	}
	return plots
}

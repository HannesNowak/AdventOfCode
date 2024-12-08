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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	grid := utils.ExtractGrid(lines)
	satellites := map[rune][]Loc{}

	for y, _ := range grid {
		for x, _ := range grid[y] {
			freq := grid[y][x]
			if freq != '.' {
				satellites[freq] = append(satellites[freq], Loc{x, y})
			}
		}
	}

	antinodes := map[rune][]Loc{}

	for freq, sats := range satellites {
		for i, sat1 := range sats {
			for j, sat2 := range sats {
				if i >= j || sat1 == sat2 {
					continue
				}
				antinodes[freq] = appendUnique(antinodes[freq], satsAntinode(grid, sats, sat1, sat2)...)
			}
		}
	}

	uniqueAntinodes := []Loc{}
	for _, antis := range antinodes {
		uniqueAntinodes = appendUnique(uniqueAntinodes, antis...)
	}

	res = len(uniqueAntinodes)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func satsAntinode(grid [][]rune, sats []Loc, sat1 Loc, sat2 Loc) []Loc {
	diff := Loc{sat1.x - sat2.x, sat1.y - sat2.y}

	addSats := false
	for i := 1; ; i++ {
		subcount := 0

		diff1 := Loc{sat1.x + diff.x*i, sat1.y + diff.y*i}
		if !utils.OutOfGrid(grid, diff1.x, diff1.y) {
			newAntinodes := appendUnique(sats, diff1)
			if len(newAntinodes) > len(sats) {
				sats = newAntinodes
				subcount++
			}
		}

		diff2 := Loc{sat2.x - diff.x*i, sat2.y - diff.y*i}
		if !utils.OutOfGrid(grid, diff2.x, diff2.y) {
			newAntinodes := appendUnique(sats, diff2)
			if len(newAntinodes) > len(sats) {
				sats = newAntinodes
				subcount++
			}
		}

		if subcount == 0 {
			break
		}
	}

	if addSats {
		sats = appendUnique(sats, sat1, sat2)
	}

	return sats
}

func appendUnique[T comparable](slice []T, elem ...T) []T {
	res := slice
	for _, el := range elem {
		if slices.Contains(res, el) {
			continue
		}
		res = append(res, el)
	}
	return res
}

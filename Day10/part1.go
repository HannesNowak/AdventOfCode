package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Trailhead struct {
	x, y int
	tops [][2]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	trailmap := initMap(lines)

	trailheads := evalTrailheads(trailmap)

	res = countTrails(trailheads)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func initMap(lines []string) [][]int {
	trailmap := [][]int{}
	for y, line := range lines {
		trailmap = append(trailmap, make([]int, len(line)))
		for x, char := range line {
			num := int(char - '0')
			trailmap[y][x] = num
		}
	}
	return trailmap
}

func evalTrailheads(trailmap [][]int) []Trailhead {
	trailheads := []Trailhead{}
	for y, row := range trailmap {
		for x, num := range row {
			if num == 0 {
				trailhead := Trailhead{x, y, [][2]int{}}
				trailheads = append(trailheads, walk(trailmap, trailhead, x, y))
			}
		}
	}
	return trailheads
}

func walk(trailmap [][]int, head Trailhead, x, y int) Trailhead {
	num := trailmap[y][x]
	if num == 9 {
		if !slices.Contains(head.tops, [2]int{x, y}) {
			head.tops = append(head.tops, [2]int{x, y})
		}
		return head
	}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != j && i+j%2 != 0 && !utils.OutOfGrid(trailmap, x+i, y+j) && trailmap[y+j][x+i] == num+1 {
				head = walk(trailmap, head, x+i, y+j)
			}
		}
	}
	return head
}

func countTrails(trailheads []Trailhead) int {
	count := 0
	for _, head := range trailheads {
		count += len(head.tops)
	}
	return count
}

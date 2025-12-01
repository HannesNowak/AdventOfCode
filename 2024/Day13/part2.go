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

	machines := [][2][3]int{}

	for lines != nil {
		m := [2][3]int{}

		// Button A
		nums := utils.Filter(utils.AllInts(lines[0]), func(n int) bool {
			return n > 0
		})
		m[0][0] = nums[0]
		m[1][0] = nums[1]

		// Button B
		nums = utils.Filter(utils.AllInts(lines[1]), func(n int) bool {
			return n > 0
		})
		m[0][1] = nums[0]
		m[1][1] = nums[1]

		// Prize
		nums = utils.Filter(utils.AllInts(lines[2]), func(n int) bool {
			return n > 0
		})
		m[0][2] = nums[0] + 1e13
		m[1][2] = nums[1] + 1e13

		machines = append(machines, m)
		lines = utils.ReadLines(scanner)
	}

	for _, m := range machines {
		a, b := solveMat(m)
		res += 3*a + b
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func solveMat(m [2][3]int) (int, int) {
	lcm := utils.LCM(m[0][0], m[1][0])

	var precise [2][3]float64
	for y, _ := range m {
		multiplier := float64(lcm) / float64(m[y][0])
		for x, _ := range m[y] {
			precise[y][x] = multiplier * float64(m[y][x])
		}
	}

	b := (precise[1][2] - precise[0][2]) / (precise[1][1] - precise[0][1])
	a := (precise[0][2] - b*precise[0][1]) / precise[0][0]

	if int(a)*m[0][0]+int(b)*m[0][1] == m[0][2] {
		return int(a), int(b)
	}
	return 0, 0
}

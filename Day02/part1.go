package main

import (
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	lines, err := utils.ReadLines(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	var res int
	startTime := time.Now()

	for _, line := range lines {
		report := getLevels(line)
		if evalReport(report) {
			fmt.Println(report, "is safe")
			res++
		}
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func getLevels(line string) []int {
	var levels []int
	var idx, level int
	for idx < len(line) {
		level, idx = utils.NextInt(line, idx, len(line))
		levels = append(levels, level)
	}
	return levels
}

func evalReport(report []int) bool {
	var inc bool
	safe := true
	if len(report) < 2 {
		return false
	}
	for idx := 0; idx < len(report); idx++ {
		if idx == 0 {
			inc = report[idx] < report[idx+1]
			continue
		}
		diff := report[idx] - report[idx-1]
		if inc && diff < 0 || !inc && diff > 0 {
			safe = false
			break
		}

		diff = utils.Abs(diff)
		if diff < 1 || diff > 3 {
			safe = false
			break
		}
	}
	return safe
}

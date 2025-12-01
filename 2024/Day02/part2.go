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

	for _, line := range lines {
		report := utils.AllInts(line)
		if evalReport(report) {
			fmt.Println(report, "is safe")
			res++
		}
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func checkReport(report []int) bool {
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

func evalReport(report []int) bool {
	fmt.Println(report)
	if checkReport(report) {
		return true
	}
	for idx := 0; idx < len(report); idx++ {
		// TODO move to utils/array.go
		temp := utils.Remove(report, idx)
		fmt.Println(report, idx, "->", temp)
		if checkReport(temp) {
			return true
		}
	}
	return false
}

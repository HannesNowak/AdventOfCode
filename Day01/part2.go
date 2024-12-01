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
	startTime := time.Now()

	firstNums, lastNums := getNumLists(lines)

	firstNums = utils.Sort(firstNums)
	lastNums = utils.Sort(lastNums)

	dist := getSimilarity(firstNums, lastNums)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(dist)
}

func getNumLists(lines []string) ([]int, []int) {
	var firstNums, lastNums []int
	var firstNum, lastNum, idx int
	for _, line := range lines {
		firstNum, idx = utils.NextInt(line, 0, len(line))
		lastNum, _ = utils.NextInt(line, idx, len(line))
		firstNums = append(firstNums, firstNum)
		lastNums = append(lastNums, lastNum)
	}
	return firstNums, lastNums
}

func getSimilarity(firstNums []int, lastNums []int) int {
	var sum int
	for _, firstNum := range firstNums {
		cnt := 0
		for _, lastNum := range lastNums {
			if firstNum == lastNum {
				cnt++
			}
		}
		sum += cnt * firstNum
	}
	return sum
}

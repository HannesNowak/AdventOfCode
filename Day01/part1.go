package main

import (
	"fmt"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	lines, err := utils.ReadLines("./input")
	if err != nil {
		fmt.Println(err)
	}

	firstNums, lastNums := getNumLists(lines)

	firstNums = utils.Sort(firstNums)
	lastNums = utils.Sort(lastNums)

	dist := getDistance(firstNums, lastNums)

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

func getDistance(firstNums []int, lastNums []int) int {
	var sum int
	for i, firstNum := range firstNums {
		lastNum := lastNums[i]
		sum += utils.Abs(lastNum - firstNum)
	}
	return sum
}

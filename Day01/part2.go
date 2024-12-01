package main

import (
	"fmt"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	lines, err := utils.ReadLines("./input")
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

func getSimilarity(firstNums []int, lastNums []int) int {
	var sum, startIdx int
	for _, firstNum := range firstNums {
		cnt := 0
	inner:
		for idx, lastNum := range lastNums[startIdx:] {
			if firstNum == lastNum {
				cnt++
			} else if firstNum < lastNum {
				startIdx = idx
				break inner
			}
		}
		sum += cnt * firstNum
	}
	return sum
}

// 18805872

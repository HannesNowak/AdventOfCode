package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	startTime := time.Now()

	firstNums, lastNums := getNumLists(lines)

	slices.Sort(firstNums)
	slices.Sort(lastNums)

	dist := getSimilarity(firstNums, lastNums)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(dist)
}

func getNumLists(lines []string) ([]int, []int) {
	var firstNums, lastNums []int
	for _, line := range lines {
		nums := utils.AllInts(line)
		firstNums = append(firstNums, nums[0])
		lastNums = append(lastNums, nums[1])
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

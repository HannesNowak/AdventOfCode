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

	dist := getDistance(firstNums, lastNums)

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

func getDistance(firstNums []int, lastNums []int) int {
	var sum int
	for i, firstNum := range firstNums {
		lastNum := lastNums[i]
		sum += utils.Abs(lastNum - firstNum)
	}
	return sum
}

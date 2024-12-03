package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
		res += evalLine(line)
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalLine(line string) int {
	var sum int
	regexr := regexp.MustCompile(`mul\(\d\d?\d?,\d\d?\d?\)`)
	matches := regexr.FindAllString(string(line), -1)
	for _, match := range matches {
		regexr := regexp.MustCompile(`\d*`)
		numStrings := regexr.FindAllString(match, -1)
		nums := []int{}
		for _, num := range numStrings {
			conv, _ := strconv.Atoi(num)
			if conv <= 0 {
				continue
			}
			nums = append(nums, conv)
		}
		sum += nums[0] * nums[1]
	}
	return sum
}

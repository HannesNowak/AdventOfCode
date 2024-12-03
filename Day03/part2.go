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

	allowed := true
	for _, line := range lines {
		var sum int
		sum, allowed = evalLine(line, allowed)
		res += sum
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalLine(line string, allowed bool) (int, bool) {
	var sum int
	allowances := map[int]bool{}

	doMatches := getDos(line)
	dontMatches := getDonts(line)
	for _, do := range doMatches {
		allowances[do[0]] = true
	}
	for _, dont := range dontMatches {
		allowances[dont[0]] = false
	}
	numMatches := getMultiplies(line)

	for idx := 0; idx < len(line); idx++ {
		if _, ok := allowances[idx]; ok {
			allowed = allowances[idx]
			delete(allowances, idx)
		}
		if !allowed {
			continue
		}
		for _, match := range numMatches {
			if idx != match[0] {
				continue
			}
			nums := convertMultiplies(line, match[0], match[1])
			sum += nums[0] * nums[1]
		}
	}
	return sum, allowed
}

func getMultiplies(line string) [][]int {
	regexr := regexp.MustCompile(`mul\(\d\d?\d?,\d\d?\d?\)`)
	matches := regexr.FindAllStringSubmatchIndex(string(line), -1)
	return matches
}

func getDos(line string) [][]int {
	regexr := regexp.MustCompile(`do\(\)`)
	matches := regexr.FindAllStringSubmatchIndex(string(line), -1)
	return matches
}

func getDonts(line string) [][]int {
	regexr := regexp.MustCompile(`don't\(\)`)
	matches := regexr.FindAllStringSubmatchIndex(string(line), -1)
	return matches
}

func convertMultiplies(line string, start int, end int) []int {
	nums := []int{}
	regexr := regexp.MustCompile(`\d*`)
	numStrings := regexr.FindAllString(line[start:end], -1)
	for _, num := range numStrings {
		conv, _ := strconv.Atoi(num)
		if conv <= 0 {
			continue
		}
		nums = append(nums, conv)
	}
	return nums
}

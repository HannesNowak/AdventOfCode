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

	doMatches := utils.RegexIndex(line, `do\(\)`)
	dontMatches := utils.RegexIndex(line, `don't\(\)`)
	for _, do := range doMatches {
		allowances[do[0]] = true
	}
	for _, dont := range dontMatches {
		allowances[dont[0]] = false
	}
	numMatches := utils.RegexIndex(line, `mul\(\d\d?\d?,\d\d?\d?\)`)

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
			nums := utils.AllInts(line[match[0]:match[1]])
			sum += nums[0] * nums[1]
		}
	}
	return sum, allowed
}

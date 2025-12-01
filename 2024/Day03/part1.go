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
		res += evalLine(line)
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalLine(line string) int {
	var sum int
	matches := utils.RegexString(line, `mul\(\d\d?\d?,\d\d?\d?\)`)
	for _, match := range matches {
		nums := utils.AllInts(match)
		sum += nums[0] * nums[1]
	}
	return sum
}

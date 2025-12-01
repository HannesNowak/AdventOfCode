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

	locks := [][5]int{}
	keys := [][5]int{}

	// parse all patterns and assign them to the correct list (locks or keys)
	for len(lines) > 0 {
		pattern, isLock := parsePattern(lines)
		if isLock {
			locks = append(locks, pattern)
		} else {
			keys = append(keys, pattern)
		}
		lines = utils.ReadLines(scanner)
	}

	// test all combinations of locks and keys
	for _, lock := range locks {
		for _, key := range keys {
			valid := true
			for i := 0; i < len(key); i++ {
				if lock[i]+key[i] > len(key) {
					valid = false
					break
				}
			}
			if valid {
				res++
			}
		}
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func parsePattern(lines []string) ([5]int, bool) {
	isLock := lines[0][0] == '#'
	pattern := [5]int{}
	for _, line := range lines[1 : len(lines)-1] {
		for idx, char := range line {
			if char == '#' {
				pattern[idx] += 1
			}
		}
	}
	return pattern, isLock
}

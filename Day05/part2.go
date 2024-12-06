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
	var res int
	startTime := time.Now()

	rules := map[int][]int{}
	for _, line := range lines {
		nums := utils.AllInts(line)
		rules[nums[0]] = append(rules[nums[0]], nums[1])
	}

	updates := utils.ReadLines(scanner)

	for _, update := range updates {
		order := utils.AllInts(update)
		res += swapUpdate(order, rules)
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func swapUpdate(update []int, rules map[int][]int) int {
	reordered := false
	for idx, page := range update {
		for i := 0; i < idx; i++ {
			if slices.Contains(rules[page], update[i]) {
				reordered = true
				update[idx] = update[idx] ^ update[i]
				update[i] = update[i] ^ update[idx]
				update[idx] = update[idx] ^ update[i]
			}
		}
	}
	if reordered {
		return update[len(update)/2]
	}
	return 0
}

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

	equations := map[int][]int{}

	for _, line := range lines {
		nums := utils.AllInts(line)
		equations[nums[0]] = nums[1:]
	}

	for sum, nums := range equations {
		if evalEquation(nums, 0, 0, sum) {
			res += sum
		}
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalEquation(nums []int, idx, subRes, endRes int) bool {
	if idx == len(nums) {
		return subRes == endRes
	}

	if evalEquation(nums, idx+1, subRes+nums[idx], endRes) {
		return true
	}

	if evalEquation(nums, idx+1, subRes*nums[idx], endRes) {
		return true
	}
	return false
}

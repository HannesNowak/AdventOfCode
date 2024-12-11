package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	stones := utils.AllInts(lines[0])
	stoneMap := convMap(stones)

	for i := 0; i < 75; i++ {
		stoneMap = evalStoneMap(stoneMap)
	}

	for _, count := range stoneMap {
		res += count
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func convMap(stones []int) map[int]int {
	stoneMap := map[int]int{}
	for _, stone := range stones {
		stoneMap[stone]++
	}
	return stoneMap
}

func evalStoneMap(stoneMap map[int]int) map[int]int {
	countMap := map[int]int{}
	for stone, count := range stoneMap {
		if stone == 0 {
			countMap[1] += count
		} else if str := strconv.Itoa(stone); len(str)%2 == 0 {
			front, _ := strconv.Atoi(str[:len(str)/2])
			back, _ := strconv.Atoi(str[len(str)/2:])
			countMap[front] += count
			countMap[back] += count
		} else {
			countMap[stone*2024] += count
		}
	}
	return countMap
}

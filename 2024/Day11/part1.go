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

	for i := 0; i < 25; i++ {
		stones = evalStones(stones)
	}

	res = len(stones)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalStones(stones []int) []int {
	generated := []int{}
	for i := 0; i < len(stones); i++ {
		if stones[i] == 0 {
			stones[i] = 1
		} else if str := strconv.Itoa(stones[i]); len(str)%2 == 0 {
			front, _ := strconv.Atoi(str[:len(str)/2])
			back, _ := strconv.Atoi(str[len(str)/2:])
			stones[i] = front
			generated = append(generated, back)
		} else {
			stones[i] *= 2024
		}
	}
	stones = append(stones, generated...)
	return stones
}

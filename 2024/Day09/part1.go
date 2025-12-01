package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	var res int
	startTime := time.Now()

	blocks := initBlocks(input)

	blocks = rearrangeBlocks(blocks)

	res = calcChecksum(blocks)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func initBlocks(input string) []int {
	blocks := []int{}
	for i, char := range input {
		num := int(char - '0')
		for j := 0; j < num; j++ {
			id := -1
			if i%2 == 0 {
				id = i / 2
			}
			blocks = append(blocks, id)
		}
	}
	return blocks
}

func rearrangeBlocks(blocks []int) []int {
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != -1 {
			continue
		}
		for j := len(blocks) - 1; j > i; j-- {
			if blocks[j] == -1 {
				continue
			}
			blocks[i] = blocks[j]
			blocks = blocks[:j]
			break
		}
	}
	return blocks
}

func calcChecksum(blocks []int) int {
	checksum := 0
	for i, num := range blocks {
		if num < 0 {
			continue
		}
		checksum += i * num
	}
	return checksum
}

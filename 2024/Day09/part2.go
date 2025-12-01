package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Block struct {
	id, size int
	start    bool
}

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

func initBlocks(input string) []Block {
	blocks := []Block{}
	for i, char := range input {
		num := int(char - '0')
		for j := 0; j < num; j++ {
			id := -1
			if i%2 == 0 {
				id = i / 2
			}
			blocks = append(blocks, Block{id, num, j == 0})
		}
	}
	return blocks
}

func rearrangeBlocks(blocks []Block) []Block {
	for j := len(blocks) - 1; j > 0; j-- {
		if !blocks[j].start || blocks[j].id == -1 {
			continue
		}
		for i := 0; i < j; i++ {
			if !blocks[i].start || blocks[i].id != -1 || blocks[i].size < blocks[j].size {
				continue
			}
			// move block to the left-most position
			for k := 0; k < blocks[j].size; k++ {
				blocks[i+k] = blocks[j+k]
				blocks[i+k].start = k == 0
				blocks[j+k].id = -1
			}
			// clean up partial blocks
			if blocks[i+blocks[j].size].id == -1 {
				idx := i + blocks[j].size
				for k := 0; k < blocks[idx].size-blocks[j].size; k++ {
					blocks[idx+k].size -= blocks[j].size
					blocks[idx+k].start = k == 0
				}
			}
			break
		}
	}
	return blocks
}

func calcChecksum(blocks []Block) int {
	checksum := 0
	for i, block := range blocks {
		if block.id < 0 {
			continue
		}
		checksum += i * block.id
	}
	return checksum
}

func printBlocks(blocks []Block) {
	for _, block := range blocks {
		out := fmt.Sprintf("%d", block.id)
		if block.id < 0 {
			out = "."
		}
		fmt.Print(out)
	}
	fmt.Println()
}

package utils

import (
	"bufio"
	"regexp"
)

func ReadLines(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			return lines
		}
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func ExtractGrid(imput []string) [][]rune {
	var grid [][]rune
	for _, line := range imput {
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	return grid
}

func OutOfGrid(grid [][]rune, x, y int) bool {
	return y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y])
}

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			print(string(cell))
		}
		println()
	}
}

func NextInt(line string, start int, amount int) (int, int) {
	var n, idx int
	for idx = start; idx < start+amount; idx++ {
		if idx >= len(line) {
			break
		}
		r := rune(line[idx])
		if r >= '0' && r <= '9' {
			n = n*10 + int(r-'0')
		} else if n > 0 {
			break
		}
	}
	return n, idx
}

func AllInts(line string) []int {
	var ints []int
	for idx := 0; idx < len(line); {
		n, next := NextInt(line, idx, len(line)-idx)
		if n > 0 {
			ints = append(ints, n)
		}
		idx = next
	}
	return ints
}

func RegexIndex(line string, re string) [][]int {
	regex := regexp.MustCompile(re)
	matches := regex.FindAllStringSubmatchIndex(line, -1)
	return matches
}

func RegexString(line string, re string) []string {
	regex := regexp.MustCompile(re)
	matches := regex.FindAllString(line, -1)
	return matches
}

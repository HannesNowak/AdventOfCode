package utils

import (
	"bufio"
	"os"
)

func ReadLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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

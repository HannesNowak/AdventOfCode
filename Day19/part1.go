package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	parts := strings.ReplaceAll(lines[0], ", ", "|")
	regex := "^(" + parts + ")+$"

	patterns := utils.ReadLines(scanner)
	for _, pattern := range patterns {
		if len(utils.RegexString(pattern, regex)) > 0 {
			res++
		}
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

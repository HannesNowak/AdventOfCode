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

	secrets := []int{}

	for _, line := range lines {
		num, _ := utils.NextInt(line, 0, len(line))
		secrets = append(secrets, num)
	}

	for idx, _ := range secrets {
		for i := 0; i < 2000; i++ {
			secrets[idx] = eval(secrets[idx])
		}
	}

	for _, secret := range secrets {
		res += secret
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func mix(result, secret int) int {
	return result ^ secret
}

func prune(result int) int {
	return result % 16777216
}

func eval(secret int) int {
	secret = prune(mix(secret*64, secret))
	secret = prune(mix(secret/32, secret))
	secret = prune(mix(secret*2048, secret))
	return secret
}

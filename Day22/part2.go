package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Seq [4]int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	secrets := []int{}
	seqAmounts := map[Seq]int{}

	for _, line := range lines {
		num, _ := utils.NextInt(line, 0, len(line))
		secrets = append(secrets, num)
	}

	for idx := range secrets {
		secretSeq := map[Seq]int{}
		secretNums := []int{secrets[idx]}
		amounts := []int{secrets[idx] % 10}
		diffs := []int{10} // 10 is never reachable
		for i := 1; i < 2001; i++ {
			secretNums = append(secretNums, eval(secretNums[i-1]))
			amounts = append(amounts, secretNums[i]%10)
			diffs = append(diffs, amounts[i]-amounts[i-1])
			if len(diffs) < 4 {
				continue
			}
			// last sequence - difference of last 4 numbers
			tempSeq := Seq{diffs[len(diffs)-4], diffs[len(diffs)-3], diffs[len(diffs)-2], diffs[len(diffs)-1]}

			// if sequence is not in map, add it
			if _, ok := secretSeq[tempSeq]; !ok {
				secretSeq[tempSeq] = amounts[i]
				// increase amount of sequence
				seqAmounts[tempSeq] += secretSeq[tempSeq]
			}
		}
	}

	for _, amount := range seqAmounts {
		res = max(res, amount)
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

// 1799 too low

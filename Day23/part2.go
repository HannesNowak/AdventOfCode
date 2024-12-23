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

	conns := map[string][]string{}

	for _, line := range lines {
		pc0 := line[:2]
		pc1 := line[3:5]
		conns[pc0] = append(conns[pc0], pc1)
		conns[pc1] = append(conns[pc1], pc0)
	}

	nodes := []string{}
	for node := range conns {
		nodes = append(nodes, node)
	}

	maxCliques := [][]string{}
	utils.MaximumCliques([]string{}, nodes, []string{}, conns, &maxCliques)

	password := ""
	for _, clique := range maxCliques {
		if len(clique) > res {
			res = len(clique)
			password = strings.Join(clique, ",")
		}
	}

	fmt.Println("Password:", password)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
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
	trios := []string{}

	for _, line := range lines {
		pc0 := line[:2]
		pc1 := line[3:5]
		conns[pc0] = append(conns[pc0], pc1)
		conns[pc1] = append(conns[pc1], pc0)
	}

	for pc, pcs := range conns {
		for i := 0; i < len(pcs); i++ {
			for j := i + 1; j < len(pcs); j++ {
				if pc[0] != 't' && pcs[i][0] != 't' && pcs[j][0] != 't' {
					continue
				}
				if !slices.Contains(conns[pcs[i]], pcs[j]) {
					continue
				}
				names := []string{pc, pcs[i], pcs[j]}
				sort.Strings(names)
				trio := strings.Join(names, ",")
				trios = utils.AppendUnique(trios, trio)
			}
		}
	}

	res = len(trios)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

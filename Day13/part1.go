package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Loc struct {
	x, y int
}

type Machine struct {
	a, b, prize Loc
	counter     [2]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	machines := []Machine{}

	for lines != nil {
		m := Machine{}
		nums := utils.Filter(utils.AllInts(lines[0]), func(n int) bool {
			return n > 0
		})
		m.a = Loc{nums[0], nums[1]}
		nums = utils.Filter(utils.AllInts(lines[1]), func(n int) bool {
			return n > 0
		})
		m.b = Loc{nums[0], nums[1]}
		nums = utils.Filter(utils.AllInts(lines[2]), func(n int) bool {
			return n > 0
		})
		m.prize = Loc{nums[0], nums[1]}
		machines = append(machines, m)
		lines = utils.ReadLines(scanner)
	}

	for i, m := range machines {
		evalButtons(&m)
		machines[i] = m
	}

	for _, m := range machines {
		res += 3*m.counter[0] + m.counter[1]
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalButtons(m *Machine) {
	if (m.a.x+m.b.x)*100 < m.prize.x || (m.a.y+m.b.y)*100 < m.prize.y {
		return
	}
	cost := 10000000
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			if 3*i+j > cost || m.a.x*i+m.b.x*j != m.prize.x || m.a.y*i+m.b.y*j != m.prize.y {
				continue
			}
			(*m).counter[0] = i
			(*m).counter[1] = j
			cost = 3*i + j
		}
	}
}

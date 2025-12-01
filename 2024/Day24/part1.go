package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Gate struct {
	inputs [2]string
	output string
	op     string
}

func (g *Gate) eval(wires map[string]int) bool {
	for _, input := range g.inputs {
		if _, ok := wires[input]; !ok {
			return false
		}
	}
	switch g.op {
	case "AND":
		wires[g.output] = wires[g.inputs[0]] & wires[g.inputs[1]]
	case "OR":
		wires[g.output] = wires[g.inputs[0]] | wires[g.inputs[1]]
	case "XOR":
		wires[g.output] = wires[g.inputs[0]] ^ wires[g.inputs[1]]
	default:
		panic("Unknown operation")
	}
	return true
}

func (g *Gate) String() string {
	return fmt.Sprintf("%v %v %v -> %v", g.inputs[0], g.op, g.inputs[1], g.output)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	wires := map[string]int{}
	for _, line := range lines {
		wires[line[:3]], _ = utils.NextInt(line, 5, len(line))
	}

	lines = utils.ReadLines(scanner)

	gates := []Gate{}
	for _, line := range lines {
		words := strings.Split(line, " ")
		gate := Gate{inputs: [2]string{words[0], words[2]}, output: words[4], op: words[1]}
		gates = append(gates, gate)
	}

	done := false
	for !done {
		done = true
		for _, gate := range gates {
			if !gate.eval(wires) {
				done = false
			}
		}
	}

	for wire, value := range wires {
		if wire[0] != 'z' {
			continue
		}
		pos, _ := utils.NextInt(wire, 1, len(wire))
		res += value << pos
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

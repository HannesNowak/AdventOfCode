package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

var regs map[rune]int
var output string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	regs = map[rune]int{}
	output = ""

	for _, line := range lines {
		regs[rune(line[9])], _ = utils.NextInt(line, 12, len(line))
	}

	lines = utils.ReadLines(scanner)

	operations := utils.RegexString(lines[0], `[0-7]+`)
	program := []int{}
	for _, p := range operations {
		num, _ := strconv.Atoi(p)
		program = append(program, num)
	}

	for i := 0; i < len(program)-1; i += 2 {
		opcode := program[i]
		operand := program[i+1]

		switch opcode {
		case 0:
			adv(operand)
		case 1:
			bxl(operand)
		case 2:
			bst(operand)
		case 3:
			if idx := jnz(operand); idx != -1 {
				i = idx - 2
			}
		case 4:
			bxc(operand)
		case 5:
			out(operand)
		case 6:
			bdv(operand)
		case 7:
			cdv(operand)
		default:
			fmt.Println("Invalid instruction")
		}
	}

	res, _ = strconv.Atoi(output)

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func comboOp(op int) int {
	switch op {
	case 4:
		return regs['A']
	case 5:
		return regs['B']
	case 6:
		return regs['C']
	default:
		return op
	}
}

func adv(op int) { // opcode 0
	regs['A'] >>= comboOp(op)
}

func bxl(op int) { // opcode 1
	regs['B'] ^= op
}

func bst(op int) { // opcode 2
	regs['B'] = comboOp(op) % 8
}

func jnz(op int) int { // opcode 3
	if regs['A'] != 0 {
		return op
	}
	return -1
}

func bxc(op int) { // opcode 4
	regs['B'] ^= regs['C']
}

func out(op int) { // opcode 5
	output += strconv.Itoa(comboOp(op) % 8)
}

func bdv(op int) { // opcode 6
	regs['B'] = regs['A'] >> comboOp(op)
}

func cdv(op int) { // opcode 7
	regs['C'] = regs['A'] >> comboOp(op)
}

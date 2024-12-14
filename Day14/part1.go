package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Robot struct {
	x, y       int
	velX, velY int
}

func (r *Robot) Move(times, width, height int) {
	r.x += r.velX * times
	r.x %= width
	if r.x < 0 {
		r.x += width
	}

	r.y += r.velY * times
	r.y %= height
	if r.y < 0 {
		r.y += height
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	width, height := 101, 103

	robots := []Robot{}

	for _, line := range lines {
		var x, y, velX, velY int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &velX, &velY)

		robot := Robot{x, y, velX, velY}
		robot.Move(100, width, height)
		robots = append(robots, robot)
	}

	quads := evalQuadrants(robots, width, height)

	res = 1
	for _, quad := range quads {
		res *= quad
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func evalQuadrants(robots []Robot, width, height int) [4]int {
	quads := [4]int{}
	midX, midY := width/2, height/2
	fmt.Println(midX, midY)
	for _, r := range robots {
		if r.x < midX && r.y < midY {
			quads[0]++
		} else if r.x > midX && r.y < midY {
			quads[1]++
		} else if r.x < midX && r.y > midY {
			quads[2]++
		} else if r.x > midX && r.y > midY {
			quads[3]++
		}
	}
	return quads
}

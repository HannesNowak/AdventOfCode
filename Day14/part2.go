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
		robots = append(robots, robot)
	}

seconds:
	for res = 1; ; res++ {
		tempRobots := []Robot{}
		single := map[int]bool{}
		for _, r := range robots {
			r.Move(res, width, height)
			tempRobots = append(tempRobots, r)
			if single[r.x*width+r.y] {
				continue seconds
			}
			single[r.x*width+r.y] = true
		}
		break
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

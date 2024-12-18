package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	utils "github.com/HannesNowak/AdventOfCode/utils"
)

type Dir struct {
	x, y int
}

type Node struct {
	x, y int
	cost int
	dir  *Dir
	prev *Node
}

func (n *Node) add(dir Dir) {
	n.x += dir.x
	n.y += dir.y
	n.cost += 1
}

func (n *Node) moveRight() Node {
	n.dir.x, n.dir.y = -n.dir.y, n.dir.x
	n.cost += 1000
	n.add(*n.dir)
	return *n
}

func (n *Node) moveLeft() Node {
	n.dir.x, n.dir.y = n.dir.y, -n.dir.x
	n.cost += 1000
	n.add(*n.dir)
	return *n
}

func (n *Node) copy() *Node {
	newNode := *n
	if n.dir != nil {
		dirCopy := *n.dir
		newNode.dir = &dirCopy
	}
	if n.prev != nil {
		prevCopy := *n.prev
		newNode.prev = &prevCopy
	}
	return &newNode
}

func (n *Node) neighbors() []Node {
	neighbors := []Node{}
	neighbors = append(neighbors, n.copy().moveRight(), n.copy().moveLeft())
	return neighbors
}

func (n *Node) heuristic(end Node) int {
	if n.x != end.x || n.y != end.y || n.dir.x != end.dir.x || n.dir.y != end.dir.y {
		return 1000 + n.manhattan(end)
	}
	return n.manhattan(end)
}

func (n *Node) manhattan(end Node) int {
	return utils.Abs(n.x-end.x) + utils.Abs(n.y-end.y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := utils.ReadLines(scanner)
	var res int
	startTime := time.Now()

	var start, end Node
	dir := Dir{1, 0}

	walls := []Node{}

	for y, _ := range lines {
		for x, char := range lines[y] {
			switch char {
			case 'S':
				start = Node{x, y, 0, &dir, nil}
			case 'E':
				end = Node{x, y, 0, nil, nil}
			case '#':
				walls = append(walls, Node{x, y, 0, nil, nil})
			}
		}
	}

	astar := astar(&start, &end, &walls)

	var minNode *Node
	for _, node := range astar {
		if minNode == nil || node.cost < minNode.cost {
			minNode = node
			fmt.Println(minNode.cost)
		}
	}
	res = minNode.cost

	for _, node := range astar {
		grid := utils.ExtractGrid(lines)
		turns := 0
		moves := 0
		fmt.Println(node.x, node.y, node.cost, node.dir.x, node.dir.y)
		for node != nil {
			if node.dir.x == 1 && node.dir.y == 0 {
				grid[node.y][node.x] = '>'
			} else if node.dir.x == -1 && node.dir.y == 0 {
				grid[node.y][node.x] = '<'
			} else if node.dir.x == 0 && node.dir.y == 1 {
				grid[node.y][node.x] = 'v'
			} else if node.dir.x == 0 && node.dir.y == -1 {
				grid[node.y][node.x] = '^'
			}
			if node.prev != nil && (node.dir.x != node.prev.dir.x || node.dir.y != node.prev.dir.y) {
				turns++
			}
			node = node.prev
			moves++
		}

		utils.PrintGrid(grid)
		fmt.Println("Turns:", turns, "Moves:", moves)
	}

	fmt.Println("Execution time:", time.Since(startTime))
	fmt.Println(res)
}

func astar(start, end *Node, walls *[]Node) []*Node {
	var closedSet []*Node
	var openSet = []*Node{start}
	start.cost = 0

	paths := []*Node{}
	for len(openSet) > 0 {
		var current = openSet[0]
		if current.x == end.x && current.y == end.y {
			paths = append(paths, current)
		}
		openSet = openSet[1:]
		closedSet = append(closedSet, current)
		neighbors := []Node{}

		forDir := current.copy()
		forDir.add(*forDir.dir)
		forDir.prev = current

		leftDir := current.copy()
		leftDir.moveLeft()
		leftDir.prev = current

		rightDir := current.copy()
		rightDir.moveRight()
		rightDir.prev = current

		neighbors = append(neighbors, *forDir, *leftDir, *rightDir)

		for _, neighbor := range neighbors {
			if isWall(neighbor, walls) {
				continue
			}

			inClosedSet := false
			for _, closedNode := range closedSet {
				if closedNode.x == neighbor.x && closedNode.y == neighbor.y && closedNode.cost <= neighbor.cost {
					inClosedSet = true
					break
				}
			}
			if inClosedSet {
				continue
			}

			inOpenSet := false
			for _, openNode := range openSet {
				if openNode.x == neighbor.x && openNode.y == neighbor.y {
					inOpenSet = true
					break
				}
			}
			if !inOpenSet {
				openSet = append(openSet, &neighbor)
			}
		}
	}
	return paths
}

func isWall(loc Node, walls *[]Node) bool {
	for _, wall := range *walls {
		if wall.x == loc.x && wall.y == loc.y {
			return true
		}
	}
	return false
}

func removeLoc(slice []Node, loc Node) []Node {
	for i, v := range slice {
		if v == loc {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// TODO 74428 too high

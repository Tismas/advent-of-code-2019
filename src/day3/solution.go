package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}

func getNewPosition(pos point, move string) point {
	dir := move[0]
	amount, e := strconv.Atoi(move[1:])
	if e != nil {
		panic(e)
	}
	switch dir {
	case 'R':
		pos.x += amount
	case 'L':
		pos.x -= amount
	case 'U':
		pos.y += amount
	case 'D':
		pos.y -= amount
	}
	return pos
}

func sign(x int) int {
	switch {
	case x < 0:
		return -1
	case x > 0:
		return 1
	default:
		return 0
	}
}

func encodePos(pos point) string {
	return strconv.Itoa(pos.x) + "," + strconv.Itoa(pos.y)
}

func getMinNotZero(current, new int) int {
	if new == 0 {
		return current
	}
	if new < current || current == 0 {
		return new
	}
	return current
}
func updateGrid(grid map[string]string, pos point, wire string) []point {
	var collitionPoints []point
	encodedPos := encodePos(pos)
	if grid[encodedPos] != wire && grid[encodedPos] != "" {
		collitionPoints = append(collitionPoints, pos)
	}
	grid[encodedPos] = wire
	return collitionPoints
}
func connect(grid map[string]string, pos point, newPos point, wire string) []point {
	var collitionPoints []point
	for ; pos.x != newPos.x; pos.x += sign(newPos.x - pos.x) {
		points := updateGrid(grid, pos, wire)
		collitionPoints = append(collitionPoints, points...)
	}
	for ; pos.y != newPos.y; pos.y += sign(newPos.y - pos.y) {
		points := updateGrid(grid, pos, wire)
		collitionPoints = append(collitionPoints, points...)
	}
	return collitionPoints
}
func getManhattanDistance(p1, p2 point) int {
	return int(math.Abs(float64(p1.x - p2.x)) + math.Abs(float64(p1.y - p2.y)))
}
func getClosest(collisionPoints []point) int {
	closest := 0
	center := point{}
	for _, pos := range collisionPoints {
		closest = getMinNotZero(closest, getManhattanDistance(pos, center))
	}
	return closest
}

func isBetween(p1, p2, goal point) bool {
	dxc := goal.x - p1.x;
	dyc := goal.y - p1.y;

	dxl := p2.x - p1.x;
	dyl := p2.y - p1.y;

	cross := dxc * dyl - dyc * dxl;
	if cross != 0 {
		return false
	}
	if (math.Abs(float64(dxl)) >= math.Abs(float64(dyl))) {
		if dxl > 0 {
			return p1.x <= goal.x && goal.x <= p2.x
		} else {
			return p2.x <= goal.x && goal.x <= p1.x
		}
	} else {
		if dyl > 0 {
			return p1.y <= goal.y && goal.y <= p2.y
		} else {
			return p2.y <= goal.y && goal.y <= p1.y
		}
	}
}
func getSteps(path []string, goal point) int {
	pos := point{}
	steps := 0
	for _, move := range path {
		newPos := getNewPosition(pos, move)
		if goal.x == pos.x && goal.y == pos.y {
			return steps
		} else if isBetween(pos, newPos, goal) {
			steps += getManhattanDistance(pos, goal)
			return steps
		} else {
			steps += getManhattanDistance(pos, newPos)
		}
		pos = newPos
	}
	return steps
}

func getLeastSteps(collisionPoints []point, path1 []string, path2 []string) int {
	var path1Distances []int
	var path2Distances []int

	for _, point := range collisionPoints {
		path1Distances = append(path1Distances, getSteps(path1, point))
		path2Distances = append(path2Distances, getSteps(path2, point))
	}

	min := path1Distances[0] + path2Distances[0]
	for i := 1; i < len(path1Distances); i++ {
		min = getMinNotZero(min, path1Distances[i]+path2Distances[i])
	}
	return min
}

func main() {
	file, e := os.Open("./input.txt")
	lines := helpers.ReadAndSplitStrings(file, e, ",")
	path1 := lines[0]
	path2 := lines[1]
	grid := make(map[string]string)
	var collisionPoints []point

	pos := point{}
	for _, move := range path1 {
		newPos := getNewPosition(pos, move)
		connect(grid, pos, newPos, "wire1")
		pos = newPos
	}

	pos = point{}
	for _, move := range path2 {
		newPos := getNewPosition(pos, move)
		points := connect(grid, pos, newPos, "wire2")
		collisionPoints = append(collisionPoints, points...)
		pos = newPos
	}

	fmt.Println("Part 1: ", getClosest(collisionPoints))
	fmt.Println("Part 2: ", getLeastSteps(collisionPoints, path1, path2))
}

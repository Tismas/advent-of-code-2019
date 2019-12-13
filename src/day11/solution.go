package main

import (
	"fmt"
	"helpers"
	"os"
	"strconv"
	"strings"
)

const up = 0
const right = 1
const down = 2
const left = 3

func getPosString(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
func getPosFromKey(key string) (int, int) {
	splitted := strings.Split(key, ",")
	x, _ := strconv.Atoi(splitted[0])
	y, _ := strconv.Atoi(splitted[1])
	return x, y
}

func rotate(direction int, color int) int {
	if color == 0 {
		direction--
	} else {
		direction++
	}
	if direction < up {
		direction = left
	}
	if direction > left {
		direction = up
	}
	return direction
}

func move(x, y, direction int) (int, int) {
	switch direction {
	case up:
		y--
	case right:
		x++
	case down:
		y++
	case left:
		x--
	}
	return x, y
}

func paint(initialMemory []int, initialColor int, totalMoves int) map[string]int {
	outputChan := make(chan int)
	inputChan := make(chan int)
	halts := make(chan bool)
	go helpers.Interprete(initialMemory, inputChan, outputChan, halts, false)

	grid := make(map[string]int)
	x, y := 0, 0
	grid[getPosString(x, y)] = initialColor
	direction := up

	for i := 0; i < totalMoves; i++ {
		inputChan <- grid[getPosString(x, y)]
		grid[getPosString(x, y)] = <-outputChan
		direction = rotate(direction, <-outputChan)
		x, y = move(x, y, direction)
	}

	return grid
}

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	memory := make([]int, 2000)
	for i, v := range initialMemory {
		memory[i] = v
	}

	fmt.Println(len(paint(memory, 0, 10751)))

	grid := paint(memory, 1, 249)
	minX, maxX := 0, 0
	minY, maxY := 0, 0
	for key := range grid {
		x, y := getPosFromKey(key)
		if x < minX {
			minX = x
		}
		if y < minY {
			minY = y
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			color := grid[getPosString(x, y)]
			if color == 1 {
				fmt.Print("#")
			}
			if color == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

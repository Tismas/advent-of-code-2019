package main

import (
	"fmt"
	"helpers"
	"os"
	"strconv"
)

func getPosString(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	memory := make([]int, 10000)
	for i, v := range initialMemory {
		memory[i] = v
	}

	outputChan := make(chan int)
	inputChan := make(chan int)
	halts := make(chan bool)
	grid := make(map[string]int)
	blockTiles := 0
	go helpers.Interprete(memory, inputChan, outputChan, halts, false)

	for {
		x, ok := <-outputChan
		if !ok {
			break
		}
		y := <-outputChan
		tile := <-outputChan

		if tile == 2 {
			blockTiles++
		}

		grid[getPosString(x, y)] = tile
	}

	fmt.Println(blockTiles)
}

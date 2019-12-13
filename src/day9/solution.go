package main

import (
	"fmt"
	"helpers"
	"os"
)

func runProgram(initialMemory []int, input int) int {
	outputChan := make(chan int)
	inputChan := make(chan int)
	go helpers.Interprete(initialMemory, inputChan, outputChan, make(chan bool), false)
	inputChan <- input
	return <-outputChan
}

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	memory := make([]int, 2000)
	for i, v := range initialMemory {
		memory[i] = v
	}
	fmt.Println("Part 1: ", runProgram(memory, 1))
	fmt.Println("Part 2: ", runProgram(memory, 2))
}

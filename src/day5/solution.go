package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	outputChan := make(chan int)
	inputChan := make(chan int)
	go helpers.Interprete(initialMemory, inputChan, outputChan, false)
	inputChan <- 1
	for output := range outputChan {
		if output != 0 {
			fmt.Println("Part 1: ", output)
		}
	}

	outputChan = make(chan int)
	inputChan = make(chan int)
	go helpers.Interprete(initialMemory, inputChan, outputChan, false)
	inputChan <- 5
	for output := range outputChan {
		if output != 0 {
			fmt.Println("Part 2: ", output)
		}
	}
}

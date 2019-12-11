package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	initialMemory[1] = 12
	initialMemory[2] = 2
	inputChan := make(chan int)
	outputChan := make(chan int)
	go helpers.Interprete(initialMemory, inputChan, outputChan, make(chan bool), true)
	fmt.Println("Part 1: ", <-outputChan)

	expectedOutput := 19690720
	noun := 0
	verb := 0
	for verb < 100 {
		inputChan = make(chan int)
		outputChan = make(chan int)
		initialMemory[1] = noun
		initialMemory[2] = verb
		go helpers.Interprete(initialMemory, inputChan, outputChan, make(chan bool), true)
		if <-outputChan == expectedOutput {
			fmt.Println("Part 2: ", 100*noun+verb)
			break
		}

		noun++
		if noun == 100 {
			verb++
			noun = 0
		}
	}
}

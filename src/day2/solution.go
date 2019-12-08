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
	output := helpers.Interprete(initialMemory, 0)
	fmt.Println("Part 1: ", output)

	expectedOutput := 19690720
	noun := 0
	verb := 0
	for verb < 100 {
		initialMemory[1] = noun
		initialMemory[2] = verb
		output := helpers.Interprete(initialMemory, 0)
		if output == expectedOutput {
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

package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	var initialMemory = helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	initialMemory[1] = 12
	initialMemory[2] = 2
	var output = helpers.Interprete(initialMemory)
	fmt.Println("Part 1: ", output)

	var expectedOutput = 19690720
	var noun = 0
	var verb = 0
	for verb < 100 {
		initialMemory[1] = noun
		initialMemory[2] = verb
		var output = helpers.Interprete(initialMemory)
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

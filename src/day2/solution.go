package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	var initialMemory = helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	var instructionPointer = 0
	var memory = initialMemory[:]
	memory[1] = 12
	memory[2] = 2
	for memory[instructionPointer] != 99 {
		var opcode = memory[instructionPointer]
		if opcode == 1 {
			var args = memory[instructionPointer+1 : instructionPointer+4]
			memory[args[2]] = memory[args[0]] + memory[args[1]]
			instructionPointer += 4
		} else if opcode == 2 {
			var args = memory[instructionPointer+1 : instructionPointer+4]
			memory[args[2]] = memory[args[0]] * memory[args[1]]
			instructionPointer += 4
		} else {
			panic("Invalid state")
		}
	}
	fmt.Println("Part 1: ", memory[0])

	// var expectedOutput = 19690720
}

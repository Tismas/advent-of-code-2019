package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	fmt.Println("Part 1")
	helpers.Interprete(initialMemory, 1)
	fmt.Println("Part 2")
	helpers.Interprete(initialMemory, 5)
}

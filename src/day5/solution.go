package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	_, output := helpers.Interprete(initialMemory, []int{1})
	fmt.Println("Part 1: ", output)
	_, output = helpers.Interprete(initialMemory, []int{5})
	fmt.Println("Part 2: ", output)
}

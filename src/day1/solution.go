package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	modules := helpers.ReadNumbers(os.Open("./input.txt"))
	fuel := 0
	for _, value := range modules {
		fuel += value/3 - 2
	}
	fmt.Println("Part 1: ", fuel)

	fuel2 := 0
	for _, value := range modules {
		value = value/3 - 2
		for value > 0 {
			fuel2 += value
			value = value/3 - 2
		}
	}
	fmt.Println("Part 2: ", fuel2)
}

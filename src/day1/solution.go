package main

import (
	"fmt"
	"helpers"
	"os"
)

func main() {
	var modules = helpers.ReadNumbers(os.Open("./input.txt"))
	fuel := 0
	for _, value := range modules {
		fuel += value/3 - 2
	}
	fmt.Println("Part 1: ", fuel)

	additionalFuel := 0
	fuelAcc := fuel/3 - 2
	for fuelAcc > 0 {
		additionalFuel += fuelAcc
		fuelAcc = fuelAcc/3 - 2
	}
	fmt.Println("Part 2: ", fuel+additionalFuel)
}

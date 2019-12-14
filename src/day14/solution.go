package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
)

func findRule(reactions map[helpers.Chemical][]helpers.Chemical, name string) (helpers.Chemical, []helpers.Chemical) {
	for product, requirements := range reactions {
		if product.Name == name {
			return product, requirements
		}
	}
	return helpers.Chemical{}, nil
}

func produce(reactions map[helpers.Chemical][]helpers.Chemical, bank map[string]int, name string, needed int, oresLimit int) int {
	oresNeeded := 0
	nextFeedbackPoint := oresLimit / 100
	product, requirements := findRule(reactions, name)
	for bank[name] < needed || (needed == 0 && oresNeeded < oresLimit) {
		if oresLimit != 0 && oresNeeded > nextFeedbackPoint {
			fmt.Println(math.Floor(float64(nextFeedbackPoint)/float64(oresLimit)*100.0), "%")
			nextFeedbackPoint += oresLimit / 100
		}
		for _, req := range requirements {
			if req.Name == "ORE" {
				oresNeeded += req.Quantity
			} else if bank[req.Name] < req.Quantity {
				oresNeeded += produce(reactions, bank, req.Name, req.Quantity, oresLimit)
			}
			bank[req.Name] -= req.Quantity
		}
		bank[product.Name] += product.Quantity
	}
	return oresNeeded
}

func main() {
	reactions := helpers.ReadReactions(os.Open("./input.txt"))
	bank := map[string]int{}
	fmt.Println("Part 1:", produce(reactions, bank, "FUEL", 1, 0))

	oresLimit := 1000000000000
	bank = map[string]int{}
	produce(reactions, bank, "FUEL", 0, oresLimit)
	fmt.Println("Part 2:", bank["FUEL"]-1)
}

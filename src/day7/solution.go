package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
)

func getThrusterOutput(initialMemory []int, phaseSetting [5]int) int {
	inputA := make(chan int, 1)
	inputB := make(chan int)
	inputC := make(chan int)
	inputD := make(chan int)
	inputE := make(chan int)
	halts := make(chan bool)

	go helpers.Interprete(initialMemory, inputA, inputB, halts, false)
	go helpers.Interprete(initialMemory, inputB, inputC, halts, false)
	go helpers.Interprete(initialMemory, inputC, inputD, halts, false)
	go helpers.Interprete(initialMemory, inputD, inputE, halts, false)
	go helpers.Interprete(initialMemory, inputE, inputA, halts, false)

	inputA <- phaseSetting[0]
	inputB <- phaseSetting[1]
	inputC <- phaseSetting[2]
	inputD <- phaseSetting[3]
	inputE <- phaseSetting[4]

	inputA <- 0

	for i := 0; i < 5; i++ {
		<-halts
	}

	return <-inputA
}

func perm(a [5]int, f func([5]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func getPhasePermutationsPart1() [][5]int {
	var permutations [][5]int

	perm([5]int{0, 1, 2, 3, 4}, func(permuation [5]int) {
		permutations = append(permutations, permuation)
	}, 0)

	return permutations
}

func getPhasePermutationsPart2() [][5]int {
	var permutations [][5]int

	perm([5]int{5, 6, 7, 8, 9}, func(permuation [5]int) {
		permutations = append(permutations, permuation)
	}, 0)

	return permutations
}

func main() {
	initialMemory := helpers.ReadNumbersSingleLine(os.Open("./input.txt"))
	maxOutput := 0
	permutations := getPhasePermutationsPart1()
	for i := 0; i < len(permutations); i++ {
		outputThruster := getThrusterOutput(initialMemory, permutations[i])
		maxOutput = int(math.Max(float64(maxOutput), float64(outputThruster)))
	}

	fmt.Println("Part 1: ", maxOutput)

	maxOutput = 0
	permutations = getPhasePermutationsPart2()
	for i := 0; i < len(permutations); i++ {
		outputThruster := getThrusterOutput(initialMemory, permutations[i])
		maxOutput = int(math.Max(float64(maxOutput), float64(outputThruster)))
	}

	fmt.Println("Part 2: ", maxOutput)
}

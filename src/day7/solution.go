package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
)

func getThrusterOutput(initialMemory []int, phaseSetting [5]int) int {
	inputChanA := make(chan int)
	outputChanA := make(chan int)
	inputChanB := make(chan int)
	outputChanB := make(chan int)
	inputChanC := make(chan int)
	outputChanC := make(chan int)
	inputChanD := make(chan int)
	outputChanD := make(chan int)
	inputChanE := make(chan int)
	outputChanE := make(chan int)

	go helpers.Interprete(initialMemory, inputChanA, outputChanA, false)
	go helpers.Interprete(initialMemory, inputChanB, outputChanB, false)
	go helpers.Interprete(initialMemory, inputChanC, outputChanC, false)
	go helpers.Interprete(initialMemory, inputChanD, outputChanD, false)
	go helpers.Interprete(initialMemory, inputChanE, outputChanE, false)

	inputChanA <- phaseSetting[0]
	inputChanB <- phaseSetting[1]
	inputChanC <- phaseSetting[2]
	inputChanD <- phaseSetting[3]
	inputChanE <- phaseSetting[4]

	inputChanA <- 0
	inputChanB <- (<-outputChanA)
	inputChanC <- (<-outputChanB)
	inputChanD <- (<-outputChanC)
	inputChanE <- (<-outputChanD)

	// var output int
	// for output = range outputChanE {
	// 	inputChanA <- output
	// 	inputChanB <- (<-outputChanA)
	// 	inputChanC <- (<-outputChanB)
	// 	inputChanD <- (<-outputChanC)
	// 	inputChanE <- (<-outputChanD)
	// }

	return <-outputChanE
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

	// maxOutput = 0
	// permutations = getPhasePermutationsPart2()
	// for i := 0; i < len(permutations); i++ {
	// 	outputThruster := getThrusterOutput(initialMemory, permutations[i])
	// 	maxOutput = int(math.Max(float64(maxOutput), float64(outputThruster)))
	// }

	// fmt.Println("Part 2: ", maxOutput)
}

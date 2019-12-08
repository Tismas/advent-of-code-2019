package main

import (
	"fmt"
	"helpers"
	"math"
	"os"
)

func getThrusterOutput(initialMemory []int, phaseSetting [5]int) int {
	_, outputA := helpers.Interprete(initialMemory, []int{phaseSetting[0], 0})
	_, outputB := helpers.Interprete(initialMemory, []int{phaseSetting[1], outputA})
	_, outputC := helpers.Interprete(initialMemory, []int{phaseSetting[2], outputB})
	_, outputD := helpers.Interprete(initialMemory, []int{phaseSetting[3], outputC})
	_, outputThruster := helpers.Interprete(initialMemory, []int{phaseSetting[4], outputD})
	return outputThruster
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

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isCorrectPassword(num int) bool {
	digits := strings.Split(strconv.Itoa(num), "")
	minAllowedDigit, _ := strconv.Atoi(digits[0])
	areTwoAdjecentSame := false
	for i := 1; i < len(digits); i++ {
		digit, _ := strconv.Atoi(digits[i])
		if digit < minAllowedDigit {
			return false
		}
		if minAllowedDigit == digit {
			areTwoAdjecentSame = true
		}
		minAllowedDigit = digit
	}
	return areTwoAdjecentSame
}

func isCorrectPart2(num int) bool {
	digits := strings.Split(strconv.Itoa(num), "")
	streak := 1
	for i := 1; i < len(digits); i++ {
		if digits[i] == digits[i-1] {
			streak++
		} else {
			if streak == 2 {
				return true
			}
			streak = 1
		}
	}
	return streak == 2
}

func main() {
	min := 402328
	max := 864247
	var possiblePasswordsPart1 []int
	var possiblePasswordsPart2 []int

	for i := min; i < max; i++ {
		if isCorrectPassword(i) {
			possiblePasswordsPart1 = append(possiblePasswordsPart1, i)
			if isCorrectPart2(i) {
				possiblePasswordsPart2 = append(possiblePasswordsPart2, i)
			}
		}
	}
	fmt.Println("Part 1: ", len(possiblePasswordsPart1))
	fmt.Println("Part 2: ", len(possiblePasswordsPart2))
}

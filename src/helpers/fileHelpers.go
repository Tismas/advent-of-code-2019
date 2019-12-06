package helpers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadStrings returns array of strings(lines) in a file
func ReadStrings(file *os.File, e error) []string {
	if e != nil {
		panic(e)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// ReadAndSplitStrings reads lines and split them by delimeter
func ReadAndSplitStrings(file *os.File, e error, delimeter string) [][]string {
	lines := ReadStrings(file, e)
	splitted := make([][]string, len(lines))
	for i, line := range lines {
		splitted[i] = strings.Split(line, delimeter)
	}
	return splitted
}

// ReadNumbers returns array of integers(each in new line) in a file
func ReadNumbers(file *os.File, e error) []int {
	var numbers []int
	lines := ReadStrings(file, e)
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

// ReadNumbersSingleLine returns array of integers in the first line of the file
func ReadNumbersSingleLine(file *os.File, e error) []int {
	var numbers []int
	line := ReadAndSplitStrings(file, e, ",")[0]
	for _, numberStr := range line {
		num, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

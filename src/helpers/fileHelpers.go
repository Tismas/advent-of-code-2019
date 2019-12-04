package helpers

import (
	"bufio"
	"os"
	"strconv"
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

// ReadNumbers returns array of integers(each in new line) in a file
func ReadNumbers(file *os.File, e error) []int {
	var lines = ReadStrings(file, e)
	var numbers []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

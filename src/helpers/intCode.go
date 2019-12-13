package helpers

import (
	"strconv"
	"strings"
)

var argCounts = map[int]int{
	1: 3,
	2: 3,
	3: 1,
	4: 1,
	5: 2,
	6: 2,
	7: 3,
	8: 3,
	9: 1,
}

var relativeBase = 0

func getValueUsignMode(memory []int, arg, mode int) int {
	if mode == 0 {
		return memory[arg]
	}
	if mode == 2 {
		return memory[relativeBase+arg]
	}
	return arg
}
func getAddressUsingMode(arg, mode int) int {
	if mode == 2 {
		return relativeBase + arg
	}
	return arg
}

func add(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	v2 := getValueUsignMode(memory, args[1], modes[1])
	address := getAddressUsingMode(args[2], modes[2])
	memory[address] = v1 + v2
	return pointer + 4
}
func multiply(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	v2 := getValueUsignMode(memory, args[1], modes[1])
	address := getAddressUsingMode(args[2], modes[2])
	memory[address] = v1 * v2
	return pointer + 4
}
func set(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	address := getAddressUsingMode(args[0], modes[0])
	memory[address] = <-inputChan
	return pointer + 2
}
func get(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	outputChan <- v1
	return pointer + 2
}
func jumpIfTrue(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	v2 := getValueUsignMode(memory, args[1], modes[1])
	if v1 != 0 {
		return v2
	}
	return pointer + 3
}
func jumpIfFalse(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	v2 := getValueUsignMode(memory, args[1], modes[1])
	if v1 == 0 {
		return v2
	}
	return pointer + 3
}
func lessThan(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	v2 := getValueUsignMode(memory, args[1], modes[1])
	address := getAddressUsingMode(args[2], modes[2])
	if v1 < v2 {
		memory[address] = 1
	} else {
		memory[address] = 0
	}
	return pointer + 4
}
func equals(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	v1 := getValueUsignMode(memory, args[0], modes[0])
	v2 := getValueUsignMode(memory, args[1], modes[1])
	address := getAddressUsingMode(args[2], modes[2])
	if v1 == v2 {
		memory[address] = 1
	} else {
		memory[address] = 0
	}
	return pointer + 4
}
func setRelativeBase(memory []int, pointer int, modes [3]int, inputChan chan int, outputChan chan int, args ...int) int {
	relativeBase += getValueUsignMode(memory, args[0], modes[0])
	return pointer + 2
}

var handlers = map[int]func([]int, int, [3]int, chan int, chan int, ...int) int{
	1: add,
	2: multiply,
	3: set,
	4: get,
	5: jumpIfTrue,
	6: jumpIfFalse,
	7: lessThan,
	8: equals,
	9: setRelativeBase,
}

func reverse(arr []string) []string {
	var reversed []string
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func interpreteInstruction(instruction int) (int, [3]int) {
	instructionString := strconv.Itoa(instruction)
	instructionArr := reverse(strings.Split(instructionString, ""))
	opcode, mode0, mode1, mode2 := 0, 0, 0, 0
	for i := 0; i < len(instructionArr); i++ {
		value, _ := strconv.Atoi(instructionArr[i])
		switch i {
		case 0:
			opcode += value
		case 1:
			opcode += value * 10
		case 2:
			mode0 = value
		case 3:
			mode1 = value
		case 4:
			mode2 = value
		}
	}
	modes := [3]int{mode0, mode1, mode2}
	return opcode, modes
}

// Interprete intcode and return end memory
func Interprete(initialMemory []int, inputChan chan int, outputChan chan int, halts chan bool, outputFromMemory bool) {
	memory := append([]int(nil), initialMemory...)
	instructionPointer := 0
	for memory[instructionPointer] != 99 {
		opcode, modes := interpreteInstruction(memory[instructionPointer])
		if handler, ok := handlers[opcode]; ok {
			jump := argCounts[opcode] + 1
			args := append([]int(nil), memory[instructionPointer+1:instructionPointer+jump]...)
			instructionPointer = handler(memory, instructionPointer, modes, inputChan, outputChan, args...)
		} else {
			panic("Something went wrong in Intcode")
		}
	}
	if outputFromMemory {
		outputChan <- memory[0]
	}
	close(outputChan)
	relativeBase = 0
	halts <- true
}

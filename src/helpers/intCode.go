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
}

func add(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	v2 := args[1]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	if modes[1] == 0 {
		v2 = memory[v2]
	}
	address := args[2]
	memory[address] = v1 + v2
	return pointer + 4
}
func multiply(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	v2 := args[1]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	if modes[1] == 0 {
		v2 = memory[v2]
	}
	address := args[2]
	memory[address] = v1 * v2
	return pointer + 4
}
func set(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	address := args[0]
	memory[address] = inputValues[0]
	return pointer + 2
}
func print(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	return pointer + 2
}
func jumpIfTrue(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	v2 := args[1]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	if modes[1] == 0 {
		v2 = memory[v2]
	}
	if v1 != 0 {
		return v2
	}
	return pointer + 3
}
func jumpIfFalse(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	v2 := args[1]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	if modes[1] == 0 {
		v2 = memory[v2]
	}
	if v1 == 0 {
		return v2
	}
	return pointer + 3
}
func lessThan(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	v2 := args[1]
	v3 := args[2]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	if modes[1] == 0 {
		v2 = memory[v2]
	}
	if v1 < v2 {
		memory[v3] = 1
	} else {
		memory[v3] = 0
	}
	return pointer + 4
}
func equals(memory []int, pointer int, modes [3]int, inputValues []int, args ...int) int {
	v1 := args[0]
	v2 := args[1]
	v3 := args[2]
	if modes[0] == 0 {
		v1 = memory[v1]
	}
	if modes[1] == 0 {
		v2 = memory[v2]
	}
	if v1 == v2 {
		memory[v3] = 1
	} else {
		memory[v3] = 0
	}
	return pointer + 4
}

var handlers = map[int]func([]int, int, [3]int, []int, ...int) int{
	1: add,
	2: multiply,
	3: set,
	4: print,
	5: jumpIfTrue,
	6: jumpIfFalse,
	7: lessThan,
	8: equals,
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

// Interprete intcode and return output(memory[0])
func Interprete(initialMemory []int, inputValues []int) (int, int) {
	memory := append([]int(nil), initialMemory...)
	instructionPointer := 0
	output := 0
	for memory[instructionPointer] != 99 {
		opcode, modes := interpreteInstruction(memory[instructionPointer])
		if handler, ok := handlers[opcode]; ok {
			jump := argCounts[opcode] + 1
			args := append([]int(nil), memory[instructionPointer+1:instructionPointer+jump]...)
			instructionPointer = handler(memory, instructionPointer, modes, inputValues, args...)
			if opcode == 3 {
				inputValues = inputValues[1:]
			}
			if opcode == 4 {
				v1 := args[0]
				if modes[0] == 0 {
					v1 = memory[v1]
				}
				output = v1
			}
		} else {
			panic("Something went wrong in Intcode")
		}
	}
	return memory[0], output
}

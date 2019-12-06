package helpers

func add(memory []int, addressA int, addressB int, dest int) {
	memory[dest] = memory[addressA] + memory[addressB]
}

func multiply(memory []int, addressA int, addressB int, dest int) {
	memory[dest] = memory[addressA] * memory[addressB]
}

var handlers = map[int]func([]int, int, int, int){
	1: add,
	2: multiply,
}

// Interprete intcode and return output(memory[0])
func Interprete(initialMemory []int) int {
	memory := make([]int, len(initialMemory))
	copy(memory, initialMemory)
	instructionPointer := 0
	for memory[instructionPointer] != 99 {
		opcode := memory[instructionPointer]
		if handler, ok := handlers[opcode]; ok {
			args := memory[instructionPointer+1 : instructionPointer+4]
			handler(memory, args[0], args[1], args[2])
			instructionPointer += 4
		} else {
			panic("Something went wrong in Intcode")
		}
	}
	return memory[0]
}

package intcode

type Interpreter struct {
	tape           []int
	programCounter int
}

func (inter *Interpreter) load(data []int) {
	inter.tape = data
	inter.programCounter = 0
}

func (inter *Interpreter) dump() []int {
	return inter.tape
}

func (inter *Interpreter) step() bool {
	op := inter.tape[inter.programCounter]

	switch op {
	case 1:
		in0 := inter.tape[inter.programCounter+1]
		in1 := inter.tape[inter.programCounter+2]
		out := inter.tape[inter.programCounter+3]
		inter.tape[out] = inter.tape[in0] + inter.tape[in1]
		inter.programCounter += 4
		return true
	case 2:
		in0 := inter.tape[inter.programCounter+1]
		in1 := inter.tape[inter.programCounter+2]
		out := inter.tape[inter.programCounter+3]
		inter.tape[out] = inter.tape[in0] * inter.tape[in1]
		inter.programCounter += 4
		return true
	case 99:
		return false
	default:
		return false
	}
}

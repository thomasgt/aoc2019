package intcode

type Computer struct {
	memory         []int
	programCounter int
}

func (c *Computer) Load(data []int) {
	c.memory = data
	c.programCounter = 0
}

func (c *Computer) Dump() []int {
	return c.memory
}

func (c *Computer) Step() bool {
	op := c.memory[c.programCounter]

	switch op {
	case 1:
		in0 := c.memory[c.programCounter+1]
		in1 := c.memory[c.programCounter+2]
		out := c.memory[c.programCounter+3]
		c.memory[out] = c.memory[in0] + c.memory[in1]
		c.programCounter += 4
		return true
	case 2:
		in0 := c.memory[c.programCounter+1]
		in1 := c.memory[c.programCounter+2]
		out := c.memory[c.programCounter+3]
		c.memory[out] = c.memory[in0] * c.memory[in1]
		c.programCounter += 4
		return true
	case 99:
		return false
	default:
		return false
	}
}

func (c *Computer) Execute(memory []int, noun int, verb int) int {
	memory[1] = noun
	memory[2] = verb
	c.Load(memory)
	for c.Step() {}
	return c.memory[0]
}

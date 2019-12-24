package intcode

import "testing"

func runProgram(t *testing.T, input []int, output []int) {
	c := Computer{}
	c.Load(input)
	for c.Step() {}

	dump := c.Dump()

	if len(dump) != len(output) {
		t.Errorf("output size does not match expected size (%d vs %d)", len(dump), len(output))
	}

	for i := 0; i < len(dump); i++ {
		if dump[i] != output[i] {
			t.Errorf("output program state at %d does not match expected (%d vs %d)", i, dump[i], output[i])
		}
	}
}

func TestComputer(t *testing.T) {
	t.Run("1+1=2", func(t *testing.T) {
		runProgram(t, []int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99})
	})

	t.Run("3x2=6", func(t *testing.T) {
		runProgram(t, []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99})
	})

	t.Run("99x99=9801", func(t *testing.T) {
		runProgram(t, []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801})
	})

	t.Run("1+1=2,5x6=30", func(t *testing.T) {
		runProgram(t, []int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99})
	})
}

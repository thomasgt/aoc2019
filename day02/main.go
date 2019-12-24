package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/thomasgt/aoc2019/day02/intcode"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: input state0 state1")
	}

	file, fileErr := os.Open(os.Args[1])
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()

	var state [2]int

	_, state0Err := fmt.Sscan(os.Args[2], &state[0])
	if state0Err != nil {
		log.Fatal(state0Err)
	}

	_, state1Err := fmt.Sscan(os.Args[3], &state[1])
	if state1Err != nil {
		log.Fatal(state1Err)
	}

	var program []int
	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	})

	for scanner.Scan() {
		var code int
		_, codeErr := fmt.Sscan(scanner.Text(), &code)
		if codeErr != nil {
			log.Fatal(codeErr)
		}
		program = append(program, code)
	}

	program[1] = state[0]
	program[2] = state[1]

	computer := intcode.Computer{}
	result := computer.Execute(program, state[0], state[1])

	log.Printf("Result[0]: %d", result)
}

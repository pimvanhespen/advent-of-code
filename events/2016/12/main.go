package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

type Input = [][]string

func main() {
	event := aoc.New(2016, 12, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) ([]string, error) {
		return strings.Fields(s), nil
	})
}

func part1(input Input) string {
	regs := map[string]int{}
	compute(input, regs)
	return aoc.Result(regs["a"])
}

func part2(input Input) string {
	regs := map[string]int{"c": 1}

	compute(input, regs)

	return aoc.Result(regs["a"])
}

func compute(input Input, regs map[string]int) {

	valueOf := func(s string) int {
		if v, err := strconv.Atoi(s); err == nil {
			return v
		}
		return regs[s]
	}

	var ptr int
	for ptr < len(input) {
		instr := input[ptr]

		switch instr[0] {
		case "cpy":
			regs[instr[2]] = valueOf(instr[1])
			ptr++
		case "inc":
			regs[instr[1]]++
			ptr++
		case "dec":
			regs[instr[1]]--
			ptr++
		case "jnz":
			if valueOf(instr[1]) != 0 {
				ptr += valueOf(instr[2])
			} else {
				ptr++
			}
		}
	}
}

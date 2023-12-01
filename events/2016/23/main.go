package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input = [][]string

func main() {
	event := aoc.New(2016, 23, parse)
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
	return "n/a"
}

func isRegister(s string) bool {
	return len(s) == 1 && strings.Contains("abcd", s)
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

		log.Println(regs)
		log.Println(instr)

		time.Sleep(100 * time.Millisecond)

		switch instr[0] {
		case "cpy":
			if isRegister(instr[2]) {
				regs[instr[2]] = valueOf(instr[1])
			}
			ptr++
		case "inc":
			if isRegister(instr[1]) {
				regs[instr[1]]++
			}
			ptr++
		case "dec":
			if isRegister(instr[1]) {
				regs[instr[1]]--
			}
			ptr++
		case "jnz":
			if v := valueOf(instr[1]); v != 0 {
				if v < 0 {
					ptr -= valueOf(instr[2])
				} else {
					ptr += valueOf(instr[2])
				}
			} else {
				ptr++
			}
		case "tgl":

			n := valueOf(instr[1])
			targetPtr := ptr + n

			if targetPtr < 0 || targetPtr >= len(input) {
				ptr++
				continue
			}

			targetInstr := input[targetPtr]

			switch len(targetInstr) {
			case 2:
				switch targetInstr[0] {
				case "inc":
					targetInstr[0] = "dec"
				default:
					targetInstr[0] = "inc"
				}
			case 3:
				switch targetInstr[0] {
				case "jnz":
					targetInstr[0] = "cpy"
				default:
					targetInstr[0] = "jnz"
				}
			}

			ptr++
		}
	}
}

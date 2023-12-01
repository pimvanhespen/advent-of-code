package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []int

func main() {
	event := aoc.New(2017, 5, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, strconv.Atoi)
}

func part1(input Input) string {
	jumps := make([]int, len(input))
	copy(jumps, input)

	steps := 0

	var ptr int

	for ptr >= 0 && ptr < len(jumps) {
		var jump int
		jump, jumps[ptr] = jumps[ptr], jumps[ptr]+1

		ptr += jump
		steps++
	}

	return strconv.Itoa(steps)
}

func part2(input Input) string {
	jumps := make([]int, len(input))
	copy(jumps, input)

	steps := 0

	var ptr int

	for ptr >= 0 && ptr < len(jumps) {
		var jump int
		if jumps[ptr] >= 3 {
			jump, jumps[ptr] = jumps[ptr], jumps[ptr]-1
		} else {
			jump, jumps[ptr] = jumps[ptr], jumps[ptr]+1
		}

		ptr += jump
		steps++
	}

	return strconv.Itoa(steps)
}

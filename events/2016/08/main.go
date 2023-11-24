package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []Instruction

func main() {
	event := aoc.New(2016, 8, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, parseInstruction)
}

func parseInstruction(s string) (Instruction, error) {
	fields := strings.Fields(s)

	if len(fields) == 5 {
		if fields[1] == "row" {
			var y, n int
			_, err := fmt.Sscanf(s, "rotate row y=%d by %d", &y, &n)
			if err != nil {
				return nil, err
			}
			return ShiftRow{y, n}, nil
		} else if fields[1] == "column" {
			var x, n int
			_, err := fmt.Sscanf(s, "rotate column x=%d by %d", &x, &n)
			if err != nil {
				return nil, err
			}
			return ShiftColumn{x, n}, nil
		}
	}

	for len(fields) == 2 {
		return RectFromString(s)
	}

	return nil, fmt.Errorf("unknown instruction: %s", s)
}

func part1(input Input) string {

	d := NewDisplay(50, 6)
	for _, i := range input {
		i.Apply(d)
	}
	return fmt.Sprint(d.Count())
}

func part2(input Input) string {
	d := NewDisplay(50, 6)
	for _, i := range input {
		i.Apply(d)
	}
	return fmt.Sprintf("\n%s\n", d.String())
}

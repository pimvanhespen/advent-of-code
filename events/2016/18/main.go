package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Seed []byte
	Rows int
}

func main() {
	event := aoc.New(2016, 18, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := aoc.ReadAll(r)
	if err != nil {
		return Input{}, err
	}

	return Input{
		Seed: b,
		Rows: 40,
	}, nil
}

func nextRow(row []byte) []byte {
	next := make([]byte, len(row))
	for i := 0; i < len(next); i++ {
		var prev []byte

		switch i {
		case 0:
			prev = []byte{'.', row[i], row[i+1]}
		case len(next) - 1:
			prev = []byte{row[i-1], row[i], '.'}
		default:
			prev = row[i-1 : i+2]
		}

		switch string(prev) {
		case "^^.":
			next[i] = '^'
		case ".^^":
			next[i] = '^'
		case "^..":
			next[i] = '^'
		case "..^":
			next[i] = '^'
		default:
			next[i] = '.'
		}
	}

	return next
}

func solve(row []byte, cycles int) int {
	safe := bytes.Count(row, []byte("."))
	for i := 1; i < cycles; i++ {
		row = nextRow(row)
		safe += bytes.Count(row, []byte("."))
	}

	return safe
}

func part1(input Input) string {
	return fmt.Sprint(solve(input.Seed, input.Rows))
}

func part2(input Input) string {
	return fmt.Sprint(solve(input.Seed, 400_000))
}

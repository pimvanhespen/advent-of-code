package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input = []byte

func main() {
	event := aoc.New(2016, 9, io.ReadAll)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func part1(input Input) string {

	count, err := walk(input, sizeOf)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(count)
}

func part2(input Input) string {

	count, err := walk(input, valueOf)
	if err != nil {
		panic(err)
	}
	return strconv.Itoa(count)
}

func walk(input Input, scoreFn func([]byte) (int, error)) (int, error) {
	input = bytes.TrimSpace(input)

	var count int

	for len(input) > 0 {
		switch input[0] {
		case '(': // Parse marker
			var m marker
			n, err := extract(input, &m)
			if err != nil {
				panic(err)
			}

			value, err := scoreFn(input[n : n+m.width])
			if err != nil {
				return 0, err
			}

			input = input[n+m.width:]
			count += m.times * value

		case '\n', '\r', ' ', '\t': // Ignore whitespace
			input = input[1:]

		default: // Count regular characters
			input = input[1:]
			count++
		}
	}

	return count, nil
}

func sizeOf(data []byte) (int, error) {
	return len(data), nil
}

func valueOf(data []byte) (int, error) {
	return walk(data, valueOf)
}

type marker struct {
	width int
	times int
}

func extract(src []byte, m *marker) (int, error) {
	if src[0] != '(' {
		return 0, fmt.Errorf("not a marker")
	}

	sep := bytes.IndexByte(src, 'x')
	width, err := strconv.Atoi(string(src[1:sep]))
	if err != nil {
		return 0, fmt.Errorf("width: %w", err)
	}

	end := bytes.IndexByte(src, ')')
	times, err := strconv.Atoi(string(src[sep+1 : end]))
	if err != nil {
		return 0, fmt.Errorf("times: %w", err)
	}

	m.width = width
	m.times = times

	return end + 1, nil
}

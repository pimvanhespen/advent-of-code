package main

import (
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input = []byte

func main() {
	event := aoc.New(2017, 1, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ReadAll(r)
}

func part1(input Input) string {
	var sum int
	for i := 0; i < len(input); i++ {
		next := (i + 1) % len(input)
		if input[i] == input[next] {
			sum += int(input[i] - '0')
		}
	}
	return fmt.Sprint(sum)
}

func part2(input Input) string {
	var sum int
	for i := 0; i < len(input); i++ {
		next := (i + len(input)/2) % len(input)
		if input[i] == input[next] {
			sum += int(input[i] - '0')
		}
	}
	return fmt.Sprint(sum)
}

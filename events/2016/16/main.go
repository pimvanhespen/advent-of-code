package main

import (
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct{}

func main() {
	event := aoc.New(2016, 16, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	panic("implement me")
}

func part1(input Input) string {
	return "n/a"
}

func part2(input Input) string {
	return "n/a"
}

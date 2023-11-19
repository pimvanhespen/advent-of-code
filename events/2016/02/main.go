package main

import (
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct{}

func main() {
	event := aoc.New(2016, 2, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func part1(i Input) string {
	panic("implement me")
}

func part2(i Input) string {
	panic("implement me")
}

func parse(r io.Reader) (Input, error) {
	panic("implement me")
}

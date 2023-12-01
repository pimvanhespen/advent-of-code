package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input = int

func main() {
	event := aoc.New(2015, 20, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := aoc.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(b))
}

func part1(input Input) string {

	houses := make([]int, input/10)

	for i := 1; i < input/10; i++ {
		for j := i; j < input/10; j += i {
			houses[j] += i * 10
		}
	}

	for i, v := range houses {
		if v >= input {
			return strconv.Itoa(i)
		}
	}

	return "n/a"
}

func part2(input Input) string {
	houses := make([]int, input/10)

	for i := 1; i < input/10; i++ {
		for j := i; j <= i*50 && j < len(houses); j += i {
			houses[j] += i * 11
		}
	}

	for i, v := range houses {
		if v >= input {
			return strconv.Itoa(i)
		}
	}

	return "n/a"
}

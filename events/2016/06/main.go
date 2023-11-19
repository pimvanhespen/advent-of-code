package main

import (
	"fmt"
	"io"
	"math"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input [][]byte

type Alphabet [26]uint16

func main() {
	event := aoc.New(2016, 6, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) ([]byte, error) {
		return []byte(s), nil
	})
}

func part1(input Input) string {
	return solve(input, high)
}

func solve(input Input, selector func(alphabet Alphabet) byte) string {
	n := len(input[0])

	counts := make([]Alphabet, n)
	for _, word := range input {
		for i, c := range word {
			counts[i][c-'a']++
		}
	}

	result := make([]byte, n)
	for i, a := range counts {
		result[i] = selector(a)
	}

	return string(result)
}

func high(a Alphabet) byte {
	var c byte
	var n uint16

	for i, v := range a {
		if v > n {
			n = v
			c = byte('a' + i)
		}
	}

	return c
}

func low(a Alphabet) byte {
	var c byte
	var n uint16 = math.MaxUint16

	for i, v := range a {
		if v < n {
			n = v
			c = byte('a' + i)
		}
	}

	return c
}

func part2(input Input) string {
	return solve(input, low)
}

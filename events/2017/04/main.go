package main

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input [][]string

func main() {
	event := aoc.New(2017, 4, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(line string) ([]string, error) {
		return strings.Fields(line), nil
	})
}

func part1(input Input) string {
	var sum int

	for _, line := range input {
		if valid(line) {
			sum++
		}
	}

	return fmt.Sprint(sum)
}

func valid(line []string) bool {
	seen := make(map[string]bool)

	for _, word := range line {
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

func part2(input Input) string {
	var sum int

	for _, line := range input {
		if valid2(line) {
			sum++
		}
	}

	return fmt.Sprint(sum)
}

func valid2(line []string) bool {
	seen := make(map[string]bool)

	for _, word := range line {
		b := []byte(word)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		word = string(b)
		if seen[word] {
			return false
		}
		seen[word] = true
	}

	return true
}

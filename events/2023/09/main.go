package main

import (
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input [][]int

func main() {
	event := aoc.New(2023, 9, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func part1(input Input) string {
	var sum int
	for _, line := range input {
		sum += extrapolateForward(line)
	}
	return aoc.Result(sum)
}

func part2(input Input) string {
	var sum int
	for _, line := range input {
		sum += extrapolateBackward(line)
	}
	return aoc.Result(sum)
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(line string) ([]int, error) {
		return aoc.Ints(line)
	})
}

func extrapolateForward(line []int) int {
	levels := buildLevels(line)

	var addition int
	for i := len(levels) - 1; i >= 0; i-- {
		addition = levels[i][len(levels[i])-1] + addition
	}
	return addition
}

func extrapolateBackward(line []int) int {
	levels := buildLevels(line)

	var addition int
	for i := len(levels) - 1; i >= 0; i-- {
		addition = levels[i][0] - addition
	}
	return addition
}

func buildLevels(line []int) [][]int {
	var lists [][]int
	lists = append(lists, line)

	for {
		n := derivate(lists[len(lists)-1])
		lists = append(lists, n)
		if len(n) <= 1 || allEqual(n) {
			break
		}
	}
	return lists
}

func derivate(line []int) []int {
	result := make([]int, len(line)-1)
	for pi, c := range line[1:] {
		result[pi] = c - line[pi]
	}
	return result
}

func allEqual(line []int) bool {
	for _, c := range line[1:] {
		if c != line[0] {
			return false
		}
	}
	return true
}

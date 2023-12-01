package main

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input [][]int

func main() {
	event := aoc.New(2017, 2, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) ([]int, error) {
		fields := strings.Fields(s)
		numbers := make([]int, len(fields))
		for i, field := range fields {
			numbers[i] = aoc.Must(strconv.Atoi(field))
		}
		return numbers, nil
	})
}

func part1(input Input) string {
	var sum int
	for _, row := range input {
		low, high := math.MaxInt, math.MinInt
		for _, n := range row {
			low = min(low, n)
			high = max(high, n)
		}
		sum += high - low
	}
	return strconv.Itoa(sum)
}

func part2(input Input) string {
	var sum int
	for _, row := range input {
		for i, n := range row {
			for j, m := range row {
				if i == j {
					continue
				}
				if n%m == 0 {
					sum += n / m
				}
			}
		}
	}
	return strconv.Itoa(sum)
}

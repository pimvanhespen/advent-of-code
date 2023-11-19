package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
)

type Input struct {
	Weights []int
}

func main() {

	aoc.New[Input](24, os.Stdout,
		aoc.NewParser(parse),
		aoc.NewSolver(part1),
		aoc.NewSolver(part2),
	)
}

func parse(reader io.Reader) (Input, error) {
	weights, err := aoc.ParseLines(reader, strconv.Atoi)
	if err != nil {
		return Input{}, err
	}

	return Input{Weights: weights}, nil
}

func quantumEntanglement(weights []int) int {
	qe := 1
	for _, w := range weights {
		qe *= w
	}
	return qe
}

func smallest(weights []int, size int) []int {

	options := permute(weights, size)

	shortest := math.MaxInt
	for _, o := range options {
		if shortest > len(o) {
			shortest = len(o)
		}
	}

	for i := len(options) - 1; i >= 0; i-- {
		if len(options[i]) > shortest {
			copy(options[i:], options[i+1:])
			options = options[:len(options)-1]
		}
	}

	slices.SortFunc(options, func(a, b []int) int {
		return quantumEntanglement(a) - quantumEntanglement(b)
	})

	return options[0]
}

func permute(weights []int, target int) [][]int {
	var configurations [][]int
	for i, weight := range weights {
		if weight == target {
			configurations = append(configurations, []int{weight})
		} else if weight < target {
			for _, perm := range permute(weights[i+1:], target-weight) {
				configurations = append(configurations, append([]int{weight}, perm...))
			}
		}
	}
	return configurations
}

func part1(i Input) string {

	sum := aoc.Sum(i.Weights...)
	if sum%3 != 0 {
		panic(fmt.Sprintf("sum %d is not divisible by 3", sum))
	}

	size := sum / 3

	small := smallest(i.Weights, size)

	qe := quantumEntanglement(small)

	return fmt.Sprintf("%d", qe)
}

func part2(i Input) string {

	sum := aoc.Sum(i.Weights...)
	if sum%4 != 0 {
		panic(fmt.Sprintf("sum %d is not divisible by 4", sum))
	}

	size := sum / 4

	small := smallest(i.Weights, size)

	qe := quantumEntanglement(small)

	return fmt.Sprintf("%d", qe)
}

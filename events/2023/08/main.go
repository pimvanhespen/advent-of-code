package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/arithmatic"
	"io"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Route string

	Branches map[string][2]string
}

func (i Input) String() string {
	var sb strings.Builder

	_, _ = fmt.Fprintln(&sb, i.Route)
	_, _ = fmt.Fprintln(&sb)
	for k, v := range i.Branches {
		_, _ = fmt.Fprintf(&sb, "%s = (%s, %s)\n", k, v[0], v[1])
	}

	return sb.String()
}

func main() {
	event := aoc.New(2023, 8, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	lines, err := aoc.ReadLines(r)
	if err != nil {
		return Input{}, err
	}

	route := lines[0]

	branches := make(map[string][2]string)
	for _, line := range lines[2:] {
		var key, left, right string
		_, err = fmt.Sscanf(line, "%3s = (%3s, %3s)", &key, &left, &right)
		if err != nil {
			return Input{}, fmt.Errorf("failed to parse line %q: %w", line, err)
		}
		branches[key] = [2]string{left, right}
	}

	return Input{
		Route:    route,
		Branches: branches,
	}, nil
}

func part1(input Input) string {
	var count int

	current := "AAA"
	for current != "ZZZ" {
		if 'L' == input.Route[count%len(input.Route)] {
			current = input.Branches[current][0]
		} else {
			current = input.Branches[current][1]
		}
		count++
	}

	return fmt.Sprint(count)
}

func part2(input Input) string {

	var aas []string

	for k := range input.Branches {
		if k[2] != 'A' {
			continue
		}
		aas = append(aas, k)
	}

	steps := make([]int, len(aas))

	for i, aa := range aas {
		current := aa
		for current[2] != 'Z' {
			if 'L' == input.Route[steps[i]%len(input.Route)] {
				current = input.Branches[current][0]
			} else {
				current = input.Branches[current][1]
			}
			steps[i]++
		}
	}

	// calc LCM - least common multiple of all steps
	lcm := arithmatic.LCM(steps...)
	return fmt.Sprint(lcm)
}

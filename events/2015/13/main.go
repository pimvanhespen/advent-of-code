package main

import (
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

type Person = string

type DirectedScores map[Person]map[Person]int

func main() {

	r, err := aoc.Get(13)
	if err != nil {
		panic(err)
	}

	input, err := parseInput(r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(input))

	const me Person = "me"

	input[me] = make(map[Person]int)
	for p := range input {
		input[p][me] = 0
		input[me][p] = 0
	}

	fmt.Println("Part 2:", part1(input))
}

func part1(in DirectedScores) int {

	people := make([]Person, 0, len(in))
	for p := range in {
		people = append(people, p)
	}

	perms := permutations(people)

	max := 0

	for _, p := range perms {
		h := happiness(p, in)
		if h > max {
			max = h
		}
	}

	return max
}

func happiness(order []Person, in DirectedScores) int {
	var h int

	for i := 0; i < len(order); i++ {
		var left, right Person

		left = order[i%len(order)]
		right = order[(i+1)%len(order)]

		h += in[left][right]
		h += in[right][left]
	}

	return h
}

func permutations(in []Person) [][]Person {
	if len(in) == 1 {
		return [][]Person{in}
	}

	var out [][]Person

	for i, p := range in {
		rest := make([]Person, len(in)-1)
		copy(rest, in[:i])
		copy(rest[i:], in[i+1:])

		for _, perm := range permutations(rest) {
			out = append(out, append([]Person{p}, perm...))
		}
	}

	return out
}

func parseInput(input io.Reader) (DirectedScores, error) {

	b, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	lines := strings.Split(string(b), "\n")

	m := make(DirectedScores)

	for _, l := range lines {
		if l == "" {
			continue
		}

		l = strings.Trim(l, ".")
		parts := strings.Split(l, " ")

		subject := parts[0]
		verb := parts[2]
		amount := parts[3]
		neighbour := parts[10]

		//fmt.Println(subject, verb, amount, neighbour)/**/

		n, err := strconv.Atoi(amount)
		if err != nil {
			return nil, fmt.Errorf("could not parse amount: %w", err)
		}

		if verb == "lose" {
			n = -n
		}

		if _, ok := m[subject]; !ok {
			m[subject] = make(map[Person]int)
		}

		m[subject][neighbour] = n
	}

	return m, nil
}

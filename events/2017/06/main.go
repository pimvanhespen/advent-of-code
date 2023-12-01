package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []int

func main() {
	event := aoc.New(2017, 6, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := aoc.ReadAll(r)
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(b))

	input := make(Input, len(fields))
	for i, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		input[i] = n
	}

	return input, nil
}

func cycle(data []byte) {
	var mx byte
	var idx int

	for i, mem := range data {
		if mem > mx {
			mx = mem
			idx = i
		}
	}

	data[idx] = 0

	units, remainder := mx/byte(len(data)), mx%byte(len(data))

	for i := 0; i < len(data); i++ {

		x := (idx + i + 1) % len(data)

		if byte(i) < remainder {
			data[x] += units + 1
		} else {
			data[x] += units
		}
	}
}

type Breaker func(s string) bool

func solve(input Input, fn Breaker) {
	mems := make([]byte, len(input))
	for i, n := range input {
		mems[i] = byte(n)
	}

	for !fn(string(mems)) {
		cycle(mems)
	}
}

func part1(input Input) string {

	seen := make(map[string]struct{})

	cycles := 0

	solve(input, func(s string) bool {
		cycles++
		_, ok := seen[s]
		if ok {
			return true
		}
		seen[s] = struct{}{}
		return false
	})

	return fmt.Sprint(cycles - 1) // -1 because the first lookup is not a cycle
}

func part2(input Input) string {

	seen := make(map[string]int)
	cycs := -1
	var dist int
	solve(input, func(s string) bool {
		cycs++

		if _, ok := seen[s]; ok {
			dist = cycs - seen[s]
			return true
		}
		seen[s] = cycs
		return false
	})

	return fmt.Sprint(dist)
}

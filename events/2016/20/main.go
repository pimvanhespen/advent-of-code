package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Ranges []Range
}

type Range struct {
	Min, Max int
}

func (r Range) Merge(other Range) Range {
	return Range{
		Min: min(r.Min, other.Min),
		Max: max(r.Max, other.Max),
	}
}

func main() {
	event := aoc.New(2016, 20, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	ranges, err := aoc.ParseLines(r, func(s string) (Range, error) {
		var rng Range
		_, err := fmt.Sscanf(s, "%d-%d", &rng.Min, &rng.Max)
		return rng, err
	})
	if err != nil {
		return Input{}, err
	}
	return Input{Ranges: ranges}, nil
}

func merge(ranges []Range) []Range {

	ms := make([]Range, len(ranges))
	copy(ms, ranges)

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Min < ms[j].Min
	})

	merged := true
	for merged {
		merged = false
		for i := len(ms) - 1; i > 0; i-- {
			curr, prev := ms[i], ms[i-1]

			// Check overlap or adjacent
			// e.g. 1-3 and 4-5 can be merged
			if !(prev.Max+1 >= curr.Min) {
				continue
			}

			ms[i-1] = ms[i-1].Merge(ms[i])
			ms = append(ms[:i], ms[i+1:]...)
			merged = true
		}
	}

	return ms
}

func part1(input Input) string {

	ranges := merge(input.Ranges)

	for prevIdx, r := range ranges[1:] {
		prev := ranges[prevIdx]
		if prev.Max+1 == r.Min {
			continue
		}
		return fmt.Sprint(prev.Max + 1)
	}
	return "n/a"
}

func part2(input Input) string {

	const limit = 4294967295

	ranges := merge(input.Ranges)

	var count int

	// lower bound is always 0
	count += ranges[0].Min
	// bounds between ranges
	for i := 1; i < len(ranges); i++ {
		count += ranges[i].Min - ranges[i-1].Max - 1
	}
	// upper bound
	count += limit - ranges[len(ranges)-1].Max

	return fmt.Sprint(count)
}

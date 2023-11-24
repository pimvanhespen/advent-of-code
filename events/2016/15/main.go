package main

import (
	"fmt"
	"io"
	"sort"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []Disc

func main() {
	event := aoc.New(2016, 15, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

type Disc struct {
	ID        int
	Positions int
	Offset    int
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) (Disc, error) {
		var d Disc
		_, err := fmt.Sscanf(s, "Disc #%d has %d positions; at time=0, it is at position %d.", &d.ID, &d.Positions, &d.Offset)
		return d, err
	})
}

func bounce(discs []Disc, t int) bool {
	for _, d := range discs {
		if (d.Offset+t+d.ID)%d.Positions != 0 {
			return false
		}
	}
	return true

}

func part1(input Input) string {
	sort.Slice(input, func(i, j int) bool {
		return input[i].ID < input[j].ID
	})

	for t := 0; ; t++ {
		if bounce(input, t) {
			return fmt.Sprint(t)
		}
	}
}

func part2(input Input) string {

	input = append(input, Disc{ID: len(input) + 1, Positions: 11, Offset: 0})
	sort.Slice(input, func(i, j int) bool {
		return input[i].ID < input[j].ID
	})

	for t := 0; ; t++ {
		if bounce(input, t) {
			return fmt.Sprint(t)
		}
	}
}

package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"sort"
)

type Input aoc.Grid

type Vec2 struct {
	X, Y int
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%d, %d)", v.X, v.Y)
}

func main() {
	event := aoc.New(2023, 11, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	g, err := aoc.ParseGrid(r)
	if err != nil {
		return Input{}, err
	}
	return Input(g), nil
}

func part1(input Input) string {
	sum := solve(input, 1)
	return fmt.Sprintf("%d", sum)
}

func part2(input Input) string {
	sum := solve(input, 1_000_000-1)
	return fmt.Sprintf("%d", sum)
}

func solve(input Input, scale int) int {
	var galaxies []Vec2

	for i, b := range input.Data {
		if b == '#' {
			galaxies = append(galaxies, Vec2{i % input.Width, i / input.Width})
		}
	}

	galaxies = rescale(galaxies, scale)

	sum := sumShortestPaths(galaxies)

	return sum
}

func rescale(galaxies []Vec2, scale int) []Vec2 {

	rescaled := make([]Vec2, len(galaxies))
	copy(rescaled, galaxies)

	sort.Slice(rescaled, func(i, j int) bool {
		return rescaled[i].Y < rescaled[j].Y
	})

	var inc int
	var prev int

	for i, g := range rescaled {
		if g.Y != prev {
			inc += g.Y - prev - 1
			prev = g.Y
		}
		rescaled[i].Y += inc * scale
	}

	sort.Slice(rescaled, func(i, j int) bool {
		return rescaled[i].X < rescaled[j].X
	})

	inc = 0
	prev = 0

	for i, g := range rescaled {
		if g.X != prev {
			inc += g.X - prev - 1
			prev = g.X
		}
		rescaled[i].X += inc * scale
	}

	return rescaled
}

func sumShortestPaths(galaxies []Vec2) int {
	var sum int
	for i, a := range galaxies {
		for _, b := range galaxies[i+1:] {
			sum += manhattanDistance(a, b)
		}
	}
	return sum
}

func manhattanDistance(a, b Vec2) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

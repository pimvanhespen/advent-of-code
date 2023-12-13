package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
)

type Grid [][]byte

func (g Grid) String() string {
	a, b := flip(g), g

	if len(a) < len(b) {
		a, b = b, a
	}

	var buf bytes.Buffer
	for i := range a {
		buf.Write(a[i])
		buf.WriteString("     ")
		if i < len(b) {
			buf.Write(b[i])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (g Grid) Copy() Grid {
	c := make(Grid, len(g))
	for i, row := range g {
		c[i] = make([]byte, len(row))
		copy(c[i], row)
	}
	return c
}

type Input []Grid

func main() {
	event := aoc.New(2023, 13, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}

	sets := bytes.Split(b, []byte("\n\n"))

	var grids []Grid
	for _, set := range sets {
		lines := bytes.Split(bytes.TrimSpace(set), []byte("\n"))
		grids = append(grids, lines)
	}

	return grids, nil
}

func part1(input Input) string {

	var sum int

	for _, grid := range input {
		sum += score(grid, 0)
	}

	return aoc.Result(sum)
}

func part2(input Input) string {
	var sum int
	for _, grid := range input {
		sum += score(grid, 1)
	}
	return aoc.Result(sum)
}

func score(grid Grid, maxSmudge int) int {
	cols := verticalSmudge(grid, maxSmudge)
	rows := horizontalSmudge(grid, maxSmudge)
	return cols + rows*100
}

func diff(a, b []byte) int {
	if len(a) != len(b) {
		panic("length mismatch")
	}
	var count int

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count
}

func verticalSmudge(grid Grid, maxSmudge int) int {
	return horizontalSmudge(flip(grid), maxSmudge)
}

func flip(grid Grid) Grid {
	flipped := make(Grid, len(grid[0]))
	for i := range flipped {
		flipped[i] = make([]byte, len(grid))
	}

	for i := range grid {
		for j := range grid[i] {
			flipped[j][i] = grid[i][j]
		}
	}

	return flipped

}

func horizontalSmudge(grid Grid, maxSmudge int) int {

	var totals []int

outer:
	for y := 0; y < len(grid)-1; y++ {
		if diff(grid[y], grid[y+1]) > maxSmudge {
			continue
		}

		// possible match
		limit := min(y, len(grid)-1-(y+1)) // least distance to top or bottom

		var dist int
		for offset := 0; offset <= limit; offset++ {
			top, bottom := y-offset, y+1+offset // top and bottom row + offset

			dist += diff(grid[top], grid[bottom])
			if dist > maxSmudge {
				continue outer
			}
		}

		if dist == maxSmudge {
			totals = append(totals, y)
		}
	}

	var sum int
	for _, t := range totals {
		sum += t + 1
	}
	return sum
}

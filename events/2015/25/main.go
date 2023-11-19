package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"os"
)

type Input struct {
	Row, Col uint
	Start    uint
}

func main() {
	aoc.New[Input](25, os.Stdout,
		aoc.NewParser(parse),
		aoc.NewSolver(solve1),
		aoc.NewSolver(solve2),
	)
}

func parse(reader io.Reader) (Input, error) {
	var row, col uint

	_, err := fmt.Fscanf(reader, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.\n", &row, &col)
	if err != nil {
		return Input{}, fmt.Errorf("parsing: %w", err)
	}

	return Input{Row: row, Col: col, Start: 20151125}, nil
}

func next(n uint) uint {
	return n * 252533 % 33554393
}

func solve1(in Input) string {

	r, c := in.Row, in.Col

	o := offset(r, c)

	current := in.Start

	for i := uint(0); i < o; i++ {
		current = next(current)
	}

	return fmt.Sprintf("%d", current)
}

func solve2(in Input) string {
	return "n/a"
}

func partial(n uint) uint {
	return (n * (n + 1)) / 2
}

// offset returns the distance from 1, 1 to the given position
func offset(row, col uint) uint {
	return partial(row+col-1) - row
}

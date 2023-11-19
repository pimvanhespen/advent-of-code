package main

import (
	"fmt"
	"io"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input []Instruction

type Instruction []byte

func main() {
	event := aoc.New(2016, 2, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func part1(input Input) string {
	code := make([]byte, len(input))
	for i, line := range input {

		var x, y int

		for _, c := range line {
			switch c {
			case 'U':
				y = max(0, y-1)
			case 'D':
				y = min(2, y+1)
			case 'L':
				x = max(0, x-1)
			case 'R':
				x = min(2, x+1)
			}
		}

		code[i] = byte('1' + y*3 + x)
	}

	return string(code)
}

type Vector2 struct {
	X, Y int
}

func part2(input Input) string {
	code := make([]byte, len(input))
	for i, line := range input {

		var pos Vector2

		for _, c := range line {
			switch c {
			case 'U':
				pos = move(pos, Vector2{0, 1})
			case 'D':
				pos = move(pos, Vector2{0, -1})
			case 'L':
				pos = move(pos, Vector2{-1, 0})
			case 'R':
				pos = move(pos, Vector2{1, 0})
			}
		}

		code[i] = digit(pos)
	}

	return string(code)
}

func move(pos, dir Vector2) Vector2 {
	newPos := Vector2{
		X: pos.X + dir.X,
		Y: pos.Y + dir.Y,
	}

	if !newPos.valid() {
		return pos
	}

	return newPos
}

func (v Vector2) valid() bool {
	return Manhattan(v, Vector2{0, 0}) <= 2
}

func Manhattan(a, b Vector2) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func digit(pos Vector2) byte {
	switch pos {
	case Vector2{0, 2}:
		return '1'
	case Vector2{-1, 1}:
		return '2'
	case Vector2{0, 1}:
		return '3'
	case Vector2{1, 1}:
		return '4'
	case Vector2{-2, 0}:
		return '5'
	case Vector2{-1, 0}:
		return '6'
	case Vector2{0, 0}:
		return '7'
	case Vector2{1, 0}:
		return '8'
	case Vector2{2, 0}:
		return '9'
	case Vector2{-1, -1}:
		return 'A'
	case Vector2{0, -1}:
		return 'B'
	case Vector2{1, -1}:
		return 'C'
	case Vector2{0, -2}:
		return 'D'
	}

	panic("invalid position")
}

func parse(r io.Reader) (Input, error) {
	parseLine := func(s string) (Instruction, error) { return []byte(s), nil }

	lines, err := aoc.ParseLines(r, parseLine)
	if err != nil {
		return Input{}, err
	}

	return lines, nil
}

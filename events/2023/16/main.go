package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
)

type Input [][]byte

func main() {
	event := aoc.New(2023, 16, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	b = bytes.TrimSpace(b)
	grid := bytes.Split(b, []byte("\n"))
	return grid, nil
}

var (
	Up    = Vec2{0, -1}
	Down  = Vec2{0, 1}
	Left  = Vec2{-1, 0}
	Right = Vec2{1, 0}
)

type Vec2 struct {
	X, Y int
}

func (v Vec2) Add(o Vec2) Vec2 {
	return Vec2{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

type Beam struct {
	Pos Vec2
	Dir Vec2
}

func Energize(grid [][]byte, initial Beam) int {
	seen := make(map[Beam]struct{})

	var beams []Beam
	beams = append(beams, initial)

	for len(beams) > 0 {
		curr := beams[0]
		beams = beams[1:]

		if _, ok := seen[curr]; ok {
			continue
		}

		// prevent caching out of bounds
		if curr.Pos.X < 0 || curr.Pos.X >= len(grid[0]) || curr.Pos.Y < 0 || curr.Pos.Y >= len(grid) {
			// do nothing
		} else {
			seen[curr] = struct{}{}
		}

		// get results of next step
		next := curr.Pos.Add(curr.Dir)

		// Check bounds of grid
		if next.X < 0 || next.X >= len(grid[0]) || next.Y < 0 || next.Y >= len(grid) {
			continue
		}

		switch grid[next.Y][next.X] {
		case '|':
			if curr.Dir == Up || curr.Dir == Down {
				beams = append(beams, Beam{Pos: next, Dir: curr.Dir})
			} else {
				beams = append(beams, Beam{Pos: next, Dir: Up})
				beams = append(beams, Beam{Pos: next, Dir: Down})
			}
		case '-':
			if curr.Dir == Left || curr.Dir == Right {
				beams = append(beams, Beam{Pos: next, Dir: curr.Dir})
			} else {
				beams = append(beams, Beam{Pos: next, Dir: Left})
				beams = append(beams, Beam{Pos: next, Dir: Right})
			}
		case '/':
			switch curr.Dir {
			case Up:
				beams = append(beams, Beam{Pos: next, Dir: Right})
			case Down:
				beams = append(beams, Beam{Pos: next, Dir: Left})
			case Left:
				beams = append(beams, Beam{Pos: next, Dir: Down})
			case Right:
				beams = append(beams, Beam{Pos: next, Dir: Up})
			}
		case '\\':
			switch curr.Dir {
			case Up:
				beams = append(beams, Beam{Pos: next, Dir: Left})
			case Down:
				beams = append(beams, Beam{Pos: next, Dir: Right})
			case Left:
				beams = append(beams, Beam{Pos: next, Dir: Up})
			case Right:
				beams = append(beams, Beam{Pos: next, Dir: Down})
			}
		default:
			beams = append(beams, Beam{Pos: next, Dir: curr.Dir})
		}
	}

	unique := make(map[Vec2]struct{})
	for k := range seen {
		unique[k.Pos] = struct{}{}
	}

	return len(unique)
}

func part1(input Input) string {
	total := Energize(input, Beam{Pos: Vec2{-1, 0}, Dir: Right})
	return aoc.Result(total)
}

func part2(input Input) string {
	var total int
	for y := range input {
		left := Energize(input, Beam{Pos: Vec2{-1, y}, Dir: Right})
		right := Energize(input, Beam{Pos: Vec2{len(input[0]), y}, Dir: Left})
		total = max(total, left, right)
	}

	for x := range input[0] {
		top := Energize(input, Beam{Pos: Vec2{x, -1}, Dir: Down})
		bottom := Energize(input, Beam{Pos: Vec2{x, len(input)}, Dir: Up})
		total = max(total, top, bottom)
	}

	return aoc.Result(total)
}

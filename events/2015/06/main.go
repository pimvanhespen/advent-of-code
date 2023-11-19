package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Command struct {
	Op    Operation
	Range Range
}

type Operation int

const (
	unknown Operation = iota
	OnOperation
	OffOperation
	ToggleOperation
)

type Point struct {
	X, Y int
}

type Grid[T any] [1000][1000]T

func (g *Grid[T]) Apply(r Range, fun func(T) T) {
	for x := r.From.X; x <= r.To.X; x++ {
		for y := r.From.Y; y <= r.To.Y; y++ {
			g[x][y] = fun(g[x][y])
		}
	}
}

func (g *Grid[T]) Count(valuer func(T) int) int {
	var count int
	for _, row := range g {
		for _, cell := range row {
			count += valuer(cell)
		}
	}
	return count
}

func main() {
	input, err := aoc.NewChallenge(2015, 6).Input()
	if err != nil {
		log.Fatal(err)
	}

	commands, err := parse(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solve1(commands))
	fmt.Println("Part 2:", solve2(commands))
}

type Range struct {
	From, To Point
}

func TurnOn(_ bool) bool {
	return true
}

func TurnOff(_ bool) bool {
	return false
}

func Toggle(b bool) bool {
	return !b
}

func solve1(commands []Command) int {
	var grid Grid[bool]

	for _, cmd := range commands {
		switch cmd.Op {
		case OnOperation:
			grid.Apply(cmd.Range, TurnOn)
		case OffOperation:
			grid.Apply(cmd.Range, TurnOff)
		case ToggleOperation:
			grid.Apply(cmd.Range, Toggle)
		}
	}

	return grid.Count(func(b bool) int {
		if b {
			return 1
		}
		return 0
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve2(commands []Command) int {
	var grid Grid[int]

	for _, cmd := range commands {
		switch cmd.Op {
		case OnOperation:
			grid.Apply(cmd.Range, func(i int) int { return i + 1 })
		case OffOperation:
			grid.Apply(cmd.Range, func(i int) int { return max(0, i-1) })
		case ToggleOperation:
			grid.Apply(cmd.Range, func(i int) int { return i + 2 })
		}
	}

	return grid.Count(func(i int) int {
		return i
	})
}

var cmdre = regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)

func parse(input io.Reader) ([]Command, error) {
	b, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	commands := make([]Command, len(lines))

	for i, line := range lines {
		matches := cmdre.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		var op Operation
		switch matches[1] {
		case "turn on":
			op = OnOperation
		case "turn off":
			op = OffOperation
		case "toggle":
			op = ToggleOperation
		default:
			return nil, fmt.Errorf("unknown operation: %s", matches[1])
		}

		var _range [2]Point

		_range[0].X, err = strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}
		_range[0].Y, err = strconv.Atoi(matches[3])
		if err != nil {
			return nil, err
		}
		_range[1].X, err = strconv.Atoi(matches[4])
		if err != nil {
			return nil, err
		}
		_range[1].Y, err = strconv.Atoi(matches[5])
		if err != nil {
			return nil, err
		}

		commands[i] = Command{
			Op:    op,
			Range: Range{_range[0], _range[1]},
		}
	}

	return commands, nil
}

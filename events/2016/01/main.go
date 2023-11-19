package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input struct {
	Steps []Step
}

type Step struct {
	IsLeft   bool
	Distance int
}

type Vector2 struct {
	X, Y int
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vector2) Mul(scalar int) Vector2 {
	return Vector2{
		X: v.X * scalar,
		Y: v.Y * scalar,
	}
}

func (v Vector2) Rotate(degrees int) Vector2 {
	if degrees%90 != 0 {
		panic("rotation must be a multiple of 90 degrees")
	}

	degrees = (degrees + 360) % 360
	switch degrees {
	case 0:
		return v
	case 90:
		return Vector2{
			X: v.Y,
			Y: -v.X,
		}
	case 180:
		return Vector2{
			X: -v.X,
			Y: -v.Y,
		}
	case 270:
		return Vector2{
			X: -v.Y,
			Y: v.X,
		}
	default:
		panic(fmt.Sprintf("invalid rotation: %d", degrees))
	}
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

func main() {
	event := aoc.New(2016, 1, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func part1(i Input) string {
	position := Vector2{}
	direction := Vector2{X: 0, Y: 1} // north

	for _, step := range i.Steps {
		if step.IsLeft {
			direction = direction.Rotate(90)
		} else {
			direction = direction.Rotate(-90)
		}

		position = position.Add(direction.Mul(step.Distance))
	}

	return fmt.Sprint(Manhattan(Vector2{}, position))
}

func part2(i Input) string {
	var position Vector2
	direction := Vector2{X: 0, Y: 1} // north

	cache := make(map[Vector2]bool)
	cache[position] = true

	for _, step := range i.Steps {
		if step.IsLeft {
			direction = direction.Rotate(90)
		} else {
			direction = direction.Rotate(-90)
		}

		for i := 0; i < step.Distance; i++ {
			position = position.Add(direction)
			if cache[position] {
				return fmt.Sprint(Manhattan(Vector2{}, position))
			} else {
				cache[position] = true
			}
		}
	}

	return "not found"
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Input{}, err
	}

	b = bytes.TrimSpace(b)

	parts := bytes.Split(b, []byte(", "))

	steps := make([]Step, len(parts))
	for i, part := range parts {
		var step Step
		if part[0] == 'L' {
			step.IsLeft = true
		}

		dst, err := strconv.Atoi(string(part[1:]))
		if err != nil {
			return Input{}, fmt.Errorf("could not parse distance: %w", err)
		}
		step.Distance = dst
		steps[i] = step
	}

	return Input{Steps: steps}, nil
}

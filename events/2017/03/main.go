package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Input = int

func main() {
	event := aoc.New(2017, 3, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return strconv.Atoi(string(aoc.Must(aoc.ReadAll(r))))
}

func bounds(input Input) int {
	for i := 1; ; i += 2 {
		if i*i >= input {
			return i
		}
	}
}

func numToCoord(input Input) (int, int) {

	if input <= 1 {
		return 0, 0
	}

	bound := bounds(input)

	x := bound / 2
	y := bound / 2

	prev := bound - 2
	prevSquared := prev * prev
	rem := input - prevSquared
	dxBound := (bound*bound - prevSquared) / 4

	option := rem / dxBound

	switch option {
	case 0:
		return x, -y + rem
	case 1:
		return x - rem%dxBound, y
	case 2: // left top to left bottom
		return -x, y - rem%dxBound
	case 3:
		return rem%dxBound - bound/2, -y
	case 4:
		return x, -y // right bottom corner
	}

	panic("unreachable")
}

func coordToNum(x, y int) int {
	mx := max(abs(x), abs(y))
	size := mx*2 + 1

	total := size * size

	prev := size - 2

	prevSquared := prev * prev

	dxBound := (size*size - prevSquared) / 4

	if y == -mx {
		return total - (mx - x)
	}

	if y == mx {
		return total - 3*dxBound + (mx - x)
	}

	if x == -mx {
		return total - 2*dxBound + (mx - y)
	}

	if x == mx {
		return total - 4*dxBound + (mx + y)
	}

	panic("unreachable")
}

func part1(input Input) string {
	if input == 1 {
		return "0"
	}

	x, y := numToCoord(input)

	return aoc.Result(fmt.Sprint(abs(x) + abs(y)))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(input Input) string {

	before := make([]int, 1, 100)
	before[0] = 1

	for i := len(before); ; i++ {
		x, y := numToCoord(i + 1) // +1 because we start at 1

		if n := coordToNum(x, y) - 1; n != i {
			panic(fmt.Sprintf("expected %d, got %d", i, n))
		}

		var sum int

		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}

				pos := coordToNum(x+dx, y+dy) - 1 // -1 because we start at 1

				if pos >= len(before) {
					continue
				}

				sum += before[pos]
			}
		}

		if sum > input {
			return aoc.Result(fmt.Sprint(sum))
		}

		before = append(before, sum)
	}

	return "n/a"
}

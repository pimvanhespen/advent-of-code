package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

func main() {
	event := aoc.New(2023, 3, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ReadMap(r)
}

func part1(input Input) string {
	var result int
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if !isDigit(input[y][x]) {
				continue
			}

			n, err := parseNumber(input, Vec2{X: x, Y: y})
			if err != nil {
				panic(err)
			}

			size := n.End.X - n.Begin.X + 1

			// check if all neighbours are empty
			neighbours := input.Neighbours(Vec2{X: x, Y: y}, size)
			if bytes.Count(neighbours, []byte{'.'}) < len(neighbours) {
				result += n.Number
			}

			// skip ahead
			x += size
		}
	}
	return fmt.Sprint(result)
}

func part2(input Input) string {

	var sum int

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if !isGear(input[y][x]) {
				continue
			}

			// found a gear, need to find the accompanying numbers
			nums := findAdjacentNumbers(input, Vec2{X: x, Y: y})
			if len(nums) != 2 {
				continue
			}

			sum += nums[0] * nums[1]
		}
	}

	return fmt.Sprint(sum)
}

type Vec2 struct {
	X, Y int
}

type Input [][]byte

func (i Input) Neighbours(v Vec2, offset int) []byte {
	result := make([]byte, 0, 8+offset*2)
	for y := v.Y - 1; y <= v.Y+1; y++ {
		for x := v.X - 1; x <= v.X+offset; x++ {
			if !i.InBounds(Vec2{X: x, Y: y}) {
				continue
			}

			// Skip w/e is within the offset.. we only want the outer ring - the neighbours
			if y == v.Y && x >= v.X && x <= v.X+offset-1 {
				continue
			}
			result = append(result, i[y][x])
		}
	}
	return result
}

func (i Input) InBounds(v Vec2) bool {
	if v.Y < 0 || v.Y >= len(i) {
		return false
	}
	if v.X < 0 || v.X >= len(i[v.Y]) {
		return false
	}
	return true
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isGear(b byte) bool {
	return b == '*'
}

func findAdjacentNumbers(input Input, v Vec2) []int {
	var result []int
	for y := v.Y - 1; y <= v.Y+1; y++ {
		for x := v.X - 1; x <= v.X+1; x++ {
			if y == v.Y && x == v.X {
				continue
			}
			if !input.InBounds(Vec2{X: x, Y: y}) {
				continue
			}
			if !isDigit(input[y][x]) {
				continue
			}

			pr, err := parseNumber(input, Vec2{X: x, Y: y})
			if err != nil {
				panic(err)
			}

			result = append(result, pr.Number)
			if pr.End.X > x {
				x = pr.End.X
			}
		}
	}
	return result
}

type ParseResult struct {
	Number int
	Begin  Vec2
	End    Vec2
}

func parseNumber(input Input, v Vec2) (ParseResult, error) {
	// find left bound
	x1, x2 := v.X, v.X

	for i := v.X; i >= 0; i-- {
		if !isDigit(input[v.Y][i]) {
			break
		}
		x1 = i
	}

	// find right bound
	for i := v.X; i < len(input[v.Y]); i++ {
		if !isDigit(input[v.Y][i]) {
			break
		}
		x2 = i
	}

	n, err := strconv.Atoi(string(input[v.Y][x1 : x2+1]))
	if err != nil {
		return ParseResult{}, err
	}

	r := ParseResult{
		Number: n,
		Begin:  Vec2{X: x1, Y: v.Y},
		End:    Vec2{X: x2, Y: v.Y},
	}
	return r, nil
}

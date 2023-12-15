package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

type Grid [][]byte

func (g Grid) Equal(other Grid) bool {
	if len(g) != len(other) {
		return false
	}
	for i := range g {
		if !bytes.Equal(g[i], other[i]) {
			return false
		}
	}
	return true
}

func (g Grid) String() string {
	var buf bytes.Buffer
	for i := range g {
		buf.Write(g[i])
		buf.WriteByte('\n')
	}
	return buf.String()
}

type Input Grid

func main() {
	event := aoc.New(2023, 14, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) ([]byte, error) {
		return []byte(s), nil
	})
}

type Vec2 struct {
	X, Y int
}

func flip90(grid Grid) Grid {
	flipped := make(Grid, len(grid[0]))
	for i := range flipped {
		flipped[i] = make([]byte, len(grid))
	}

	for i := range grid {
		for j := range grid[i] {
			ii := len(grid) - 1 - i
			flipped[j][ii] = grid[i][j]
		}
	}

	return flipped
}

func flip270(grid Grid) Grid {

	return flip90(flip180(grid))

	//flipped := make(Grid, len(grid[0]))
	//for i := range flipped {
	//	flipped[i] = make([]byte, len(grid))
	//}
	//
	//for i := range grid {
	//	for j := len(grid[i]) - 1; j >= 0; j-- {
	//		flipped[j][i] = grid[i][j]
	//	}
	//}
	//
	//return flipped
}

func flip180(grid Grid) Grid {
	flipped := make(Grid, len(grid))
	for i := range flipped {
		flipped[i] = make([]byte, len(grid[i]))
	}

	for i := range grid {
		for j := range grid[i] {
			i2 := len(grid) - 1 - i
			flipped[i2][j] = grid[i][len(grid[i])-1-j]
		}
	}

	return flipped
}

var (
	North = Vec2{0, -1}
	South = Vec2{0, 1}
	West  = Vec2{-1, 0}
	East  = Vec2{1, 0}
)

func slideBoulders(grid Grid, direction Vec2) Grid {

	switch direction {
	case North:
		grid = flip90(grid)
	case East:
	case South:
		grid = flip270(grid)
	case West:
		grid = flip180(grid)
	}

	grid = slideBouldersEast(grid)

	switch direction {
	case North:
		grid = flip270(grid)
	case East:
	case South:
		grid = flip90(grid)
	case West:
		grid = flip180(grid)
	}

	return grid
}

const (
	Squared = '#'
	Round   = 'O'
	Floor   = '.'
)

// slideBouldersEast slides boulders to the east and returns the new grid
func slideBouldersEast(grid Grid) Grid {

	// bruteforce shift boulders...
	for y := range grid {

		last := len(grid[y]) - 1
		for x := last; x >= 0; x-- {
			switch grid[y][x] {
			case Squared:
				last = x
			case Round:
				if x == last {
					continue
				}
				// move boulder to front
				offset := 1 + getLast(grid[y][x+1:last+1], Floor)
				if offset == 0 {
					continue
				}
				grid[y][x], grid[y][x+offset] = grid[y][x+offset], grid[y][x] // swap floor and boulder
			case Floor:
			}
		}
	}
	return grid
}

func slideNorth(grid Grid) Grid {
	for x := range grid[0] {
		last := len(grid)
		for y := 0; y < last; y++ {
			switch grid[y][x] {
			case Squared:
				last = y
			case Round:
				if y == last {
					continue
				}

				other := y
				for ; other > 0; other-- {
					next := other - 1
					if next < 0 || grid[next][x] != Floor {
						break
					}
				}
				if other == y {
					continue
				}
				// move boulder to front
				grid[y][x], grid[other][x] = grid[other][x], grid[y][x] // swap floor and boulder
			case Floor:
			}
		}
	}
	return grid
}

func getLastY(grid Grid, x int, b byte) int {
	for i := len(grid) - 1; i >= 0; i-- {
		if grid[i][x] != b {
			return i + 1
		}
	}
	return 0
}

func getLast(grid []byte, b byte) int {
	for i := 0; i < len(grid); i++ {
		if grid[i] != b {
			return i - 1
		}
	}
	return len(grid) - 1
}

func countLoad(grid Grid) int {
	size := len(grid)
	count := 0
	for y, row := range grid {
		for _, b := range row {
			if b == Round {
				count += size - y
			}
		}
	}
	return count
}

func part1(input Input) string {

	grid := Grid(input)

	grid = slideBoulders(grid, North)

	return aoc.Result(countLoad(grid))
}

func cycle(grid Grid) Grid {
	for i := 0; i < 4; i++ {
		grid = flip90(grid)
		grid = slideBouldersEast(grid)
	}
	return grid
}

func part2(input Input) string {

	const limit = 1_000_000_000

	m := make(map[string]Grid)
	last := make(map[string]int)

	prev := Grid(input).String()
	var curr = Grid(input)

	for i := 0; i < 1_000_000_000; i++ {
		if i%1_000_000 == 0 {
			log.Printf("%d/%d (%.2f)", i, limit, float64(i)/float64(limit)*100)
		}

		// Detect repeating pattern
		if v, ok := last[prev]; ok {

			// pattern found
			// check the size of the pattern
			size := i - v
			// calculate the remaining iterations
			remaining := (limit - i) % size

			for n := 0; n < remaining; n++ {
				curr = cycle(curr)
			}

			return aoc.Result(countLoad(curr))
		}

		last[prev] = i

		if v, ok := m[prev]; ok {
			curr = v
			continue
		}

		curr = cycle(curr)
		m[prev] = curr
		prev = curr.String()
	}

	// this'll only happen if the pattern is larger than 1_000_000_000
	return aoc.Result(countLoad(curr))
}

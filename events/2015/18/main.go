package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"slices"
	"strings"
)

type Board struct {
	X, Y  int
	Cells []bool
}

func (b *Board) Copy() Board {
	cells := make([]bool, len(b.Cells))
	copy(cells, b.Cells)
	return Board{
		X:     b.X,
		Y:     b.Y,
		Cells: cells,
	}
}

func (b *Board) Count() int {
	var count int
	for _, cell := range b.Cells {
		if cell {
			count++
		}
	}
	return count
}

func NewBoard(x int, y int) Board {
	return Board{
		X:     x,
		Y:     y,
		Cells: make([]bool, x*y),
	}
}

var offsets = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (b *Board) Set(x, y int, value bool) {
	b.Cells[x+b.X*y] = value
}

func (b *Board) Get(x, y int) bool {
	if x < 0 || x >= b.X || y < 0 || y >= b.Y {
		return false
	}
	return b.Cells[x+b.X*y]
}

func (b *Board) Neighbours(x, y int) int {
	var count int

	for _, offset := range offsets {
		if b.Get(x+offset[0], y+offset[1]) {
			count++
		}
	}

	return count
}

func (b *Board) Print(writer io.Writer) error {
	sb := strings.Builder{}

	for i, cell := range b.Cells {
		if cell {
			sb.WriteByte(On)
		} else {
			sb.WriteByte(Off)
		}

		if i%b.X == b.X-1 {
			sb.WriteByte('\n')
		}
	}

	_, err := io.Copy(writer, strings.NewReader(sb.String()))
	if err != nil {
		return err
	}

	return nil
}

const (
	On  = '#'
	Off = '.'
)

func Next2(current, next Board) {
	if current.X != next.X || current.Y != next.Y {
		panic("dimensions do not match")
	}

	var value bool

	for y := 0; y < current.Y; y++ {
		for x := 0; x < current.X; x++ {
			count := current.Neighbours(x, y)

			if current.Get(x, y) {
				value = count == 2 || count == 3
			} else {
				value = count == 3
			}

			next.Set(x, y, value)
		}
	}

	next.Set(0, 0, true)
	next.Set(0, next.Y-1, true)
	next.Set(next.X-1, 0, true)
	next.Set(next.X-1, next.Y-1, true)
}

func Next(current, next Board) {
	if current.X != next.X || current.Y != next.Y {
		panic("dimensions do not match")
	}

	var value bool

	for y := 0; y < current.Y; y++ {
		for x := 0; x < current.X; x++ {
			count := current.Neighbours(x, y)

			if current.Get(x, y) {
				value = count == 2 || count == 3
			} else {
				value = count == 3
			}

			next.Set(x, y, value)
		}
	}
}

func main() {
	reader, err := aoc.NewChallenge(2015, 18).Input()
	if err != nil {
		panic(err)
	}

	data, err := parse(reader)
	if err != nil {
		panic(err)
	}

	// Part 1
	fmt.Println("Part 1:", solve1(data.Copy()))
	// Part 2
	fmt.Println("Part 2:", solve2(data.Copy()))

}

func solve1(data Board) int {
	next := NewBoard(data.X, data.Y)

	for i := 0; i < 100; i++ {
		Next(data, next)
		data, next = next, data
	}

	var count int
	for _, cell := range data.Cells {
		if cell {
			count++
		}
	}
	return count
}

func solve2(data Board) int {
	next := NewBoard(data.X, data.Y)

	data.Set(0, 0, true)
	data.Set(0, data.Y-1, true)
	data.Set(data.X-1, 0, true)
	data.Set(data.X-1, data.Y-1, true)

	for i := 0; i < 100; i++ {
		Next2(data, next)
		data, next = next, data
	}

	var count int
	for _, cell := range data.Cells {
		if cell {
			count++
		}
	}
	return count
}

func parse(reader io.Reader) (Board, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return Board{}, err
	}
	limX := bytes.IndexByte(b, '\n')
	limY := bytes.Count(b, []byte{'\n'})

	if limX != limY {
		panic(fmt.Sprintf("x != y (%d != %d)", limX, limY))
	}

	b = slices.Clip(b)
	fmt.Println(string(b))

	b = bytes.ReplaceAll(b, []byte{'\r'}, []byte{})
	b = bytes.ReplaceAll(b, []byte{'\n'}, []byte{})
	b = slices.Clip(b)

	cells := make([]bool, len(b))

	for i, cell := range b {
		cells[i] = cell == On
	}

	display := Board{
		X:     limX,
		Y:     limY,
		Cells: cells,
	}

	return display, nil
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	const input = `.#.#.#
...##.
#....#
..#...
#.#..#
####..
`
	reader := strings.NewReader(input)

	board, err := parse(reader)
	if err != nil {
		t.Error(err)
	}

	if len(board.Cells) != 36 {
		t.Errorf("Expected 36 cells, got %d", len(board.Cells))
	}

	writer := bytes.Buffer{}
	if err = board.Print(&writer); err != nil {
		t.Error(err)
	}

	res := writer.String()
	if res != input {
		t.Errorf("Expected %q, got %q", input, res)
	}
}

var nextSamples = []string{
	`..##..
..##.#
...##.
......
#.....
#.##..`,
	`..###.
......
..###.
......
.#....
.#....`,
	`...#..
......
...#..
..##..
......
......`,
	`......
......
..##..
..##..
......
......`,
}

func TestNext(t *testing.T) {
	const input = `.#.#.#
...##.
#....#
..#...
#.#..#
####..
`

	reader := strings.NewReader(input)
	board, err := parse(reader)
	if err != nil {
		t.Error(err)
	}

	next := NewBoard(board.X, board.Y)

	for _, sample := range nextSamples {
		Next(board, next)

		writer := bytes.Buffer{}
		if err = next.Print(&writer); err != nil {
			t.Error(err)
		}

		fmt.Println("NEXT")
		_ = next.Print(os.Stdout)
		fmt.Println("EXPECTED")
		fmt.Println(sample)

		res := writer.String()
		res = strings.TrimSpace(res)
		if res != sample {
			t.Errorf("Expected %q, got %q", sample, res)
		}

		board, next = next, board
	}
}

const input = `.#.#.#
...##.
#....#
..#...
#.#..#
####..
`

func TestBoard_Neighbours(t *testing.T) {
	board, err := parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}

	neighbours := make([][]int, board.Y)
	for y := 0; y < board.Y; y++ {
		line := make([]int, board.X)
		for x := 0; x < board.X; x++ {
			line[x] = board.Neighbours(x, y)
		}
		neighbours[y] = line
	}

	fmt.Println("NEIGHBOURS")
	for _, line := range neighbours {
		fmt.Println(line)
	}

	expected := [][]int{
		{1, 0, 3, 2, 4, 1},
		{2, 2, 3, 2, 4, 3},
		{0, 2, 2, 3, 3, 1},
		{2, 4, 1, 2, 2, 2},
		{2, 6, 4, 4, 2, 0},
		{2, 4, 3, 2, 2, 1},
	}

	fmt.Println("EXPECTED")
	for _, line := range expected {
		fmt.Println(line)
	}

	for y, expect := range expected {
		got := neighbours[y]

		for x, exp := range expect {
			found := got[x]

			if found != exp {
				t.Errorf("Expected (%d,%d) %d, got %d", x, y, exp, found)
			}
		}
	}
}

var next2Samples = []string{
	`##.#.#
...##.
#....#
..#...
#.#..#
####.#
`,
	`#.##.#
####.#
...##.
......
#...#.
#.####`,
	`#..#.#
#....#
.#.##.
...##.
.#..##
##.###`,
	`#...##
####.#
..##.#
......
##....
####.#`,
	`#.####
#....#
...#..
.##...
#.....
#.#..#`,
	`##.###
.##..#
.##...
.##...
#.#...
##...#`,
}

func TestNext2(t *testing.T) {

	begin, err := parse(strings.NewReader(next2Samples[0]))
	if err != nil {
		t.Fail()
	}

	next := NewBoard(begin.X, begin.Y)

	for _, sample := range next2Samples[1:] {
		Next2(begin, next)

		writer := bytes.Buffer{}
		if err = next.Print(&writer); err != nil {
			t.Error(err)
		}

		fmt.Println()
		fmt.Println("NEXT")
		_ = next.Print(os.Stdout)

		fmt.Println("EXPECTED")
		fmt.Println(sample)

		res := writer.String()
		res = strings.TrimSpace(res)
		if res != sample {
			t.Errorf("Expected %q, got %q", sample, res)
		}

		begin, next = next, begin
	}

}

func TestNext22(t *testing.T) {
	const initial = `##.#.#
...##.
#....#
..#...
#.#..#
####.#
`

	board, err := parse(strings.NewReader(initial))
	if err != nil {
		t.Error(err)
	}

	next := NewBoard(board.X, board.Y)

	for i := 0; i < 5; i++ {
		Next2(board, next)
		board, next = next, board
	}

	res := board.Count()
	if res != 17 {
		t.Errorf("Expected 17, got %d", res)
	}

	_ = board.Print(os.Stdout)
}

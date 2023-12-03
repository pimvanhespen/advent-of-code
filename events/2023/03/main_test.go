package main

import (
	"strings"
	"testing"
)

const example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func Test_part1(t *testing.T) {

	i, err := parse(strings.NewReader(example))
	if err != nil {
		t.Fatal(err)
	}

	got := part1(i)

	want := "4361"
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}

}

func Test_part2(t *testing.T) {

	i, err := parse(strings.NewReader(example))
	if err != nil {
		t.Fatal(err)
	}

	got := part2(i)

	want := "467835"
	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}

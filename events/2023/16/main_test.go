package main

import (
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"reflect"
	"strings"
	"testing"
)

// exampleInput form the puzzle
const exampleInput = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....

`

func Test_parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want Input
	}{
		{
			name: "example",
			args: args{
				r: strings.NewReader(exampleInput),
			},
			want: Input{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.r)
			if err != nil {
				t.Fatalf("parse() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

const exampleSplit = `.|.`

const exampleSplit2 = `
.\.
.-.
`

const exampleSplit3 = `
...-..|.
.....|-.
`

const exampleCorners1 = `
.\
..
..
..
\/
`

const exampleCorners2 = `
...\.\...
..../...\
...\..|..
.....\../
....\./..
`

const exampleCorners3 = `
\...
\...
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "46",
		},
		{
			name:  "example split",
			input: aoc.Must(parse(strings.NewReader(exampleSplit))),
			want:  "2",
		},
		{
			name:  "example split 2",
			input: aoc.Must(parse(strings.NewReader(exampleSplit2))),
			want:  "5",
		},
		{
			name:  "example split 3",
			input: aoc.Must(parse(strings.NewReader(exampleSplit3))),
			want:  "10",
		},
		{
			name:  "example corners 1",
			input: aoc.Must(parse(strings.NewReader(exampleCorners1))),
			want:  "10",
		},
		{
			name:  "example corners 2",
			input: aoc.Must(parse(strings.NewReader(exampleCorners2))),
			want:  "26",
		},
		{
			name:  "example corners 3",
			input: aoc.Must(parse(strings.NewReader(exampleCorners3))),
			want:  "5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part1() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "51",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("part2() = %s, want %s", got, tt.want)
			}
		})
	}
}

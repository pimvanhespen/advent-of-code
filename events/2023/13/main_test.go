package main

import (
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"reflect"
	"strings"
	"testing"
)

// exampleInput form the puzzle
const exampleInput = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`

const exampleA = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

const exampleB = `
#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
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
			want: Input{
				Grid{
					[]byte("#.##..##."),
					[]byte("..#.##.#."),
					[]byte("##......#"),
					[]byte("##......#"),
					[]byte("..#.##.#."),
					[]byte("..##..##."),
					[]byte("#.#.##.#."),
				},
				Grid{
					[]byte("#...##..#"),
					[]byte("#....#..#"),
					[]byte("..##..###"),
					[]byte("#####.##."),
					[]byte("#####.##."),
					[]byte("..##..###"),
					[]byte("#....#..#"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.r)
			if err != nil {
				t.Fatalf("parse() error = %v", err)
			}

			for i := range got {
				if !reflect.DeepEqual(got[i].String(), tt.want[i].String()) {
					t.Errorf("parse() got = \n%vwant \n%v", got[i], tt.want[i])
				}
			}

		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleA))),
			want:  "5",
		},
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleB))),
			want:  "400",
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
			want:  "400",
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

func Test_flip(t *testing.T) {
	type args struct {
		grid Grid
	}
	tests := []struct {
		name string
		args args
		want Grid
	}{
		{
			name: "example",
			args: args{
				grid: Grid{
					[]byte("####"),
					[]byte("...."),
				},
			},
			want: Grid{
				[]byte("#."),
				[]byte("#."),
				[]byte("#."),
				[]byte("#."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flip(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flip() = %v, want %v", got, tt.want)
			}
		})
	}
}

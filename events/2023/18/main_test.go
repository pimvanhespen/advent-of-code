package main

import (
	"image"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

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
				Step{Dir: image.Pt(1, 0), Len: 6, Color: "#70c710"},
				Step{Dir: image.Pt(0, -1), Len: 5, Color: "#0dc571"},
				Step{Dir: image.Pt(-1, 0), Len: 2, Color: "#5713f0"},
				Step{Dir: image.Pt(0, -1), Len: 2, Color: "#d2c081"},
				Step{Dir: image.Pt(1, 0), Len: 2, Color: "#59c680"},
				Step{Dir: image.Pt(0, -1), Len: 2, Color: "#411b91"},
				Step{Dir: image.Pt(-1, 0), Len: 5, Color: "#8ceee2"},
				Step{Dir: image.Pt(0, 1), Len: 2, Color: "#caa173"},
				Step{Dir: image.Pt(-1, 0), Len: 1, Color: "#1b58a2"},
				Step{Dir: image.Pt(0, 1), Len: 2, Color: "#caa171"},
				Step{Dir: image.Pt(1, 0), Len: 2, Color: "#7807d2"},
				Step{Dir: image.Pt(0, 1), Len: 3, Color: "#a77fa3"},
				Step{Dir: image.Pt(-1, 0), Len: 2, Color: "#015232"},
				Step{Dir: image.Pt(0, 1), Len: 2, Color: "#7a21e3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.r)
			if err != nil {
				t.Fatalf("parse() error = %v", err)
			}

			if len(got) != len(tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)

				for i := range got {
					if !reflect.DeepEqual(got[i], tt.want[i]) {
						t.Errorf("parse() got[%d] = %v, want[%d] %v", i, got[i], i, tt.want[i])
					}
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
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "62",
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
			want:  "952408144115",
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

func Test_shoelace(t *testing.T) {
	type args struct {
		points []image.Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Wiki example",
			args: args{
				points: []image.Point{
					image.Pt(1, 6),
					image.Pt(3, 1),
					image.Pt(7, 2),
					image.Pt(4, 4),
					image.Pt(8, 5),
				},
			},
			want: 33,
		},
		{
			name: "Wiki example",
			args: args{
				points: []image.Point{
					image.Pt(1-4, 6-4),
					image.Pt(3-4, 1-4),
					image.Pt(7-4, 2-4),
					image.Pt(4-4, 4-4),
					image.Pt(8-4, 5-4),
				},
			},
			want: 33,
		},
		{
			name: "Demo Square",
			args: args{
				points: []image.Point{
					image.Pt(0, 0),
					image.Pt(10, 0),
					image.Pt(10, 10),
					image.Pt(0, 10),
				},
			},
			want: 2 * 10 * 10,
		},
		{
			name: "Demo Wiki2",
			args: args{
				points: []image.Point{
					image.Pt(3, 1),
					image.Pt(7, 2),
					image.Pt(4, 4),
					image.Pt(8, 6),
					image.Pt(1, 7),
				},
			},
			want: 41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Area(tt.args.points); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseStep(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Step
	}{
		{
			name: "example",
			args: args{
				s: "#70c710",
			},
			want: Step{
				Dir: image.Pt(1, 0),
				Len: 461937,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseStep(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

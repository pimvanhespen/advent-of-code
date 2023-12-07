package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `Time:      7  15   30
Distance:  9  40  200`

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
				{Time: 7, Distance: 9},
				{Time: 15, Distance: 40},
				{Time: 30, Distance: 200},
			},
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

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "288",
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
			want:  "71503",
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

func Test_decSize(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{n: 1},
			want: 10,
		},
		{
			args: args{n: 2},
			want: 10,
		},
		{
			args: args{n: 10},
			want: 100,
		},
		{
			args: args{n: 99},
			want: 100,
		},
		{
			args: args{n: 100},
			want: 1000,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("padding(%d)", tt.args.n)
		t.Run(name, func(t *testing.T) {
			if got := padding(tt.args.n); got != tt.want {
				t.Errorf("decSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersect(t *testing.T) {
	type args struct {
		time     int
		distance int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "example",
			args: args{
				time:     5,
				distance: 4,
			},
			want: []float64{1.0, 4.0},
		},
		{
			name: "example",
			args: args{
				time:     10,
				distance: 25,
			},
			want: []float64{5.0},
		},
		{
			name: "example",
			args: args{
				time:     -10,
				distance: 25,
			},
			want: []float64{-5.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roots(-1, float64(tt.args.time), float64(-tt.args.distance)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("roots() = %v, want %v", got, tt.want)
			}
		})
	}
}

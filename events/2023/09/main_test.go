package main

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

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
				{0, 3, 6, 9, 12, 15},
				{1, 3, 6, 10, 15, 21},
				{10, 13, 16, 21, 30, 45},
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
			want:  "114",
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
			want:  "0",
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

func Test_extrapolate(t *testing.T) {
	type args struct {
		line []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				line: []int{0, 3, 6, 9, 12, 15},
			},
			want: 18,
		},
		{
			name: "example 2",
			args: args{
				line: []int{1, 3, 6, 10, 15, 21},
			},
			want: 28,
		},
		{
			name: "example 3",
			args: args{
				line: []int{10, 13, 16, 21, 30, 45},
			},
			want: 68,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extrapolateForward(tt.args.line); got != tt.want {
				t.Errorf("extrapolateForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkParts(b *testing.B) {
	p, _ := os.Getwd()
	f, err := os.Open(filepath.Join(p, "input.txt"))
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	ex, _ := parse(f)
	b.ResetTimer()

	b.Run("part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part1(ex)
		}
	})

	b.Run("part2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2(ex)
		}
	})
}

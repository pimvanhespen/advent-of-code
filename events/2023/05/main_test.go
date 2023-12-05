package main

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

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

			if got.String() != exampleInput {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
				t.Log(got.String())
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
			want:  "35",
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
			want:  "46",
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

func TestMap_NextMapping(t *testing.T) {
	type fields struct {
		From   string
		To     string
		Scales []Scale
	}
	type args struct {
		r Range
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Mapping
	}{
		{
			name: "example",
			fields: fields{
				From: "seed",
				To:   "soil",
				Scales: []Scale{
					{
						Dst: 0,
						Src: 2,
						Len: 98,
					},
					{
						Dst: 150,
						Src: 100,
						Len: 50,
					},
				},
			},
			args: args{
				r: Range{
					From: 0,
					To:   200,
				},
			},
			want: []Mapping{
				{
					Before: Range{
						From: 0,
						To:   1,
					},
					After: Range{
						From: 0,
						To:   1,
					},
				},
				{
					Before: Range{
						From: 2,
						To:   100,
					},
					After: Range{
						From: 0,
						To:   98,
					},
				},
				{
					Before: Range{
						From: 100,
						To:   150,
					},
					After: Range{
						From: 150,
						To:   200,
					},
				},
				{
					Before: Range{
						From: 150,
						To:   200,
					},
					After: Range{
						From: 150,
						To:   200,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Map{
				From:   tt.fields.From,
				To:     tt.fields.To,
				Scales: tt.fields.Scales,
			}
			if got := m.NextMapping(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

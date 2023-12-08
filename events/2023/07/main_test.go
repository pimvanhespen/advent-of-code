package main

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

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
				{Cards: "32T3K", Bid: 765},
				{Cards: "T55J5", Bid: 684},
				{Cards: "KK677", Bid: 28},
				{Cards: "KTJJT", Bid: 220},
				{Cards: "QQQJA", Bid: 483},
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
			want:  "6440",
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
			want:  "5905",
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

func TestType2(t *testing.T) {
	tests := []struct {
		arg  string
		want HandType
	}{
		{"12345", HighCard},
		{"J1234", OnePair},
		{"J1123", ThreeOfAKind},
		{"J1122", FullHouse},
		{"J1112", FourOfAKind},
		{"J1222", FourOfAKind},
		{"J1111", FiveOfAKind},
		{"JJ123", ThreeOfAKind},
		{"JJ112", FourOfAKind},
		{"JJ111", FiveOfAKind},
		{"JJJ12", FourOfAKind},
		{"JJJ11", FiveOfAKind},
		{"JJJJ1", FiveOfAKind},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := Type2(tt.arg); got != tt.want {
				t.Errorf("Type2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"strings"
	"testing"
)

const part2Input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_part2(t *testing.T) {
	lines, err := parse(strings.NewReader(part2Input))
	if err != nil {
		t.Fatal(err)
	}

	got := part2(lines)

	if got != "281" {
		t.Errorf("part2() = %v, want %v", got, "281")
	}
}

func Test_digits2(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one1", args{"one1"}, 11},
		{"two1", args{"two1"}, 21},
		{"twone", args{"twone"}, 21},
		{"one", args{"one"}, 11},
		{"two", args{"two"}, 22},
		{"three", args{"three"}, 33},
		{"four", args{"four"}, 44},
		{"five", args{"five"}, 55},
		{"six", args{"six"}, 66},
		{"seven", args{"seven"}, 77},
		{"eight", args{"eight"}, 88},
		{"nine", args{"nine"}, 99},
		{"zero", args{"zero"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digits2(tt.args.line); got != tt.want {
				t.Errorf("digits2() = %v, want %v", got, tt.want)
			}
		})
	}
}

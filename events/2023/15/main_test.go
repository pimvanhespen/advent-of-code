package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
)

// exampleInput form the puzzle
const exampleInput = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

//func Test_parse(t *testing.T) {
//	type args struct {
//		r io.Reader
//	}
//	tests := []struct {
//		name string
//		args args
//		want Input
//	}{
//		{
//			name: "example",
//			args: args{
//				r: strings.NewReader(exampleInput),
//			},
//			want: Input{},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := parse(tt.args.r)
//			if err != nil {
//				t.Fatalf("parse() error = %v", err)
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("parse() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input Input
		want  string
	}{
		{
			name:  "example",
			input: aoc.Must(parse(strings.NewReader(exampleInput))),
			want:  "1320",
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
			want:  "145",
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

func Test_hash(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "example",
			args: args{
				b: []byte("HASH"),
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hash(tt.args.b); got != tt.want {
				t.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

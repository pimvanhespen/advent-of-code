package main

import (
	"strings"
	"testing"
)

const part1Example = `cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a`

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"example", "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			in, err := parse(strings.NewReader(part1Example))
			if err != nil {
				t.Fatal(err)
			}

			if got := part1(in); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

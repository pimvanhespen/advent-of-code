package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example", args{Input{0, 2, 7, 0}}, "5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

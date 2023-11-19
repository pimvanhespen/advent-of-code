package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		i Input
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "exmaple",
			args: args{
				i: Input{
					Weights: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
				},
			},
			want: "99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.i); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		i Input
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example",
			args: args{
				i: Input{Weights: []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}},
			},
			want: "44",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.i); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		args args
		want string
	}{
		{args{[]byte("ADVENT")}, "6"},
		{args{[]byte("A(1x5)BC")}, "7"},
		{args{[]byte("(3x3)XYZ")}, "9"},
		{args{[]byte("A(2x2)BCD(2x2)EFG")}, "11"},
		{args{[]byte("(6x1)(1x3)A")}, "6"},
		{args{[]byte("X(8x2)(3x3)ABCY")}, "18"},
		{args{[]byte("1\n2\n3\t4   5")}, "5"},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.input), func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		args args
		want string
	}{
		{args{[]byte("(3x3)XYZ")}, "9"},
		{args{[]byte("X(8x2)(3x3)ABCY")}, "20"},
		{args{[]byte("(27x12)(20x12)(13x14)(7x10)(1x12)A")}, "241920"},
		{args{[]byte("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN")}, "445"},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.input), func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

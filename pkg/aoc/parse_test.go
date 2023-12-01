package aoc

import "testing"

func TestAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{"0"}, 0},
		{"1", args{"1"}, 1},
		{"12", args{"12"}, 12},
		{"123", args{"123"}, 123},
		{"1234", args{"1234"}, 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Atoi(tt.args.s); got != tt.want {
				t.Errorf("Atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

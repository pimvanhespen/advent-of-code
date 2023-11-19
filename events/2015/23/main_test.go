package main

import "testing"

func Test_part1(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: Input{
					Program: []Instruction{
						{Op: "inc", Args: []any{"a"}},
						{Op: "jio", Args: []any{"a", 2}},
						{Op: "tpl", Args: []any{"a"}},
						{Op: "inc", Args: []any{"a"}},
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

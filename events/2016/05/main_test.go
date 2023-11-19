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
		{
			name: "abc",
			args: args{
				input: []byte("abc"),
			},
			want: "18f47a30",
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

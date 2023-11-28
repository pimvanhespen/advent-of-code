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
			name: "example",
			args: args{
				input: Input{
					Seed: []byte(".^^.^.^^^^"),
					Rows: 10,
				},
			},
			want: "38",
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

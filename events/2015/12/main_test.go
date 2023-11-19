package main

import "testing"

func Test_part2(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: `{"d":"red","e":[1,2,3,4],"f":5}`,
			},
			want: 0,
		},
		{
			name: "example",
			args: args{
				input: `[1,{"c":"red","b":2},3]`,
			},
			want: 4,
		},
		{
			name: "example",
			args: args{
				input: `[1,"red",5]`,
			},
			want: 6,
		},
		{
			name: "example",
			args: args{
				input: `[1,{"c":"red","b":2},3,{"c":"red","b":2},3]`,
			},
			want: 7,
		},
		{
			name: "example",
			args: args{
				input: `[{"red":1, "blue":[1,2,3]},3]`,
			},
			want: 3,
		},
		{
			name: "example",
			args: args{
				input: `[1,"red",{"a":["red",1,1], "b":1}]`,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

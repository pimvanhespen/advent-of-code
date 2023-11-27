package main

import (
	"reflect"
	"testing"
)

func Test_doors(t *testing.T) {
	type args struct {
		passcode string
		path     string
	}
	tests := []struct {
		name string
		args args
		want [4]bool
	}{
		{
			name: "example",
			args: args{
				passcode: "hijkl",
				path:     "",
			},
			want: [4]bool{true, true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doors(tt.args.passcode, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doors() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
					Passcode: "ihgpwlah",
				},
			},
			want: "DDRRRD",
		},
		{
			name: "example",
			args: args{
				input: Input{
					Passcode: "kglvqrro",
				},
			},
			want: "DDUDRLRRUDRD",
		},
		{
			name: "example",
			args: args{
				input: Input{
					Passcode: "ulqzkmiv",
				},
			},
			want: "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
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

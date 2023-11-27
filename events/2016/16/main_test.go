package main

import (
	"reflect"
	"testing"
)

func Test_checksum8(t *testing.T) {
	tests := []struct {
		name string
		arg  byte
		want byte
	}{
		{
			name: "example 1",
			arg:  0x0,
			want: 0x0,
		},
		{
			name: "example 2",
			arg:  0x1,
			want: 0x1,
		},
		{
			name: "example 3",
			arg:  0x2,
			want: 0x1,
		},
		{
			name: "example 4",
			arg:  0x3,
			want: 0x0,
		},
		{
			name: "example 5",
			arg:  0b01010101,
			want: 0x0F,
		},
		{
			name: "example 5",
			arg:  0b11000101,
			want: 0b0011,
		},
		{
			name: "example 5",
			arg:  0b11000111,
			want: 0b0010,
		},
		{
			name: "example 5",
			arg:  0b00110101,
			want: 0b0011,
		},
		{
			name: "example 5",
			arg:  0b00101110,
			want: 0b0101,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksum8(tt.arg); got != tt.want {
				t.Errorf("checksum8() = %08b, want %08b", got, tt.want)
			}
		})
	}
}

func Test_checksum16(t *testing.T) {

	tests := []struct {
		name string
		arg  [2]byte
		want byte
	}{
		{
			name: "example 1",
			arg:  [2]byte{0b00110101, 0b00101110},
			want: 0b11001010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksum16(tt.arg); got != tt.want {
				t.Errorf("checksum16() = %08b, want %08b", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {

	type testcast struct {
		name string
		arg  Input
		want string
	}

	tests := []testcast{
		{
			name: "example 1",
			arg: Input{
				Seed: []byte("110010110100"),
				Size: 12,
			},
			want: "100",
		},
		{
			name: "example 2",
			arg: Input{
				Seed: []byte("10000"),
				Size: 20,
			},
			want: "01100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.arg); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pat1_disk(t *testing.T) {
	type args struct {
		data []byte
		size uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				data: []byte("110010110100"),
				size: 12,
			},
			want: "100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pat1_disk(tt.args.data, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pat1_disk() = %q, want %q", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestFloor_Options(t *testing.T) {
	type fields struct {
		Chip uint8
		RTG  uint8
	}
	tests := []struct {
		name   string
		fields fields
		want   [][2]uint8
	}{
		{
			name: "single chip",
			fields: fields{
				Chip: 1,
			},
			want: [][2]uint8{
				{1, 0},
			},
		},
		{
			name: "single rtg",
			fields: fields{
				RTG: 1,
			},
			want: [][2]uint8{
				{0, 1},
			},
		},
		{
			name: "single chip and rtg",
			fields: fields{
				Chip: 1,
				RTG:  1,
			},
			want: [][2]uint8{
				{1, 0},
				{0, 1},
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Floor{
				Chip: tt.fields.Chip,
				RTG:  tt.fields.RTG,
			}
			if got := f.Options(Floor{}); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor_IsSafeWith(t *testing.T) {
	type fields struct {
		Chip uint8
		RTG  uint8
	}
	type args struct {
		c Components
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "single chip",
			fields: fields{
				Chip: 1,
			},
			args: args{
				c: Components{0, 0},
			},
			want: true,
		},
		{
			name: "single RTG",
			fields: fields{
				RTG: 1,
			},
			args: args{
				c: Components{0, 0},
			},
			want: true,
		},
		{
			name: "single chip and RTG",
			fields: fields{
				Chip: 1,
				RTG:  1,
			},
			args: args{
				c: Components{0, 0},
			},
			want: true,
		},
		{
			name: "extra RTG",
			fields: fields{
				Chip: 1,
				RTG:  1,
			},
			args: args{
				c: Components{0, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Floor{
				Chip: tt.fields.Chip,
				RTG:  tt.fields.RTG,
			}
			if got := f.IsSafeWith(tt.args.c); got != tt.want {
				t.Errorf("IsSafeWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSafe(t *testing.T) {
	type args struct {
		chip uint8
		rtg  uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "single chip",
			args: args{
				chip: 1,
				rtg:  0,
			},
			want: true,
		},
		{
			name: "single RTG",
			args: args{
				chip: 0,
				rtg:  1,
			},
			want: true,
		},
		{
			name: "equal chip and RTG",
			args: args{
				chip: 1,
				rtg:  1,
			},
			want: true,
		},
		{
			name: "diff RTG",
			args: args{
				chip: 1,
				rtg:  2,
			},
			want: false,
		},
		{
			name: "diff chip",
			args: args{
				chip: 2,
				rtg:  1,
			},
			want: false,
		},
		{
			name: "patial overlap chip",
			args: args{
				chip: 3,
				rtg:  1,
			},
			want: true,
		},
		{
			name: "patial overlap RTG",
			args: args{
				chip: 1,
				rtg:  3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSafe(tt.args.chip, tt.args.rtg); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

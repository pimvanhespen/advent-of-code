package main

import (
	"reflect"
	"sort"
	"strings"
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
		target Floor
		want   []Components
	}{
		{
			name: "single chip",
			fields: fields{
				Chip: 1,
			},
			want: []Components{
				{1, 0},
			},
		},
		{
			name: "single rtg",
			fields: fields{
				RTG: 1,
			},
			want: []Components{
				{0, 1},
			},
		},
		{
			name: "single chip and rtg",
			fields: fields{
				Chip: 1,
				RTG:  1,
			},
			want: []Components{
				{1, 0},
				{0, 1},
				{1, 1},
			},
		},
		{
			name: "2 chip and 2 rtg",
			fields: fields{
				Chip: 3,
				RTG:  3,
			},
			want: []Components{
				{1, 0},
				{2, 0},
				{1, 1},
				{2, 2},
				{3, 0},
				{0, 3},
			},
		},
		{
			name: "2 chip and 3 rtg",
			fields: fields{
				Chip: 3,
				RTG:  7,
			},
			want: []Components{
				{3, 0},
				{1, 0},
				{2, 0},
				{1, 1},
				{2, 2},
				{0, 4},
			},
		},
		{
			name: "demo",
			fields: fields{
				Chip: 1,
			},
			target: Floor{
				RTG: 3,
			},
			want: []Components{
				{1, 0},
			},
		},
		{
			name: "demo",
			fields: fields{
				RTG: 15,
			},
			want: []Components{
				{0, 1},
				{0, 2},
				{0, 3},
				{0, 4},
				{0, 5},
				{0, 6},
				{0, 8},
				{0, 9},
				{0, 10},
				{0, 12},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Floor{
				Chip: tt.fields.Chip,
				RTG:  tt.fields.RTG,
			}

			m := make(map[Components]bool)
			for _, c := range tt.want {
				m[c] = true
			}

			for _, c := range f.Options(tt.target) {
				if !m[c] {
					t.Errorf("Options() = %v too much", c)
				} else {
					m[c] = false
				}
			}

			for c, ok := range m {
				if ok {
					t.Errorf("Options() = %v missing", c)
				}
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
			want: true,
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

func TestFloor_IsSafeWithout(t *testing.T) {
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
				c: Components{1, 0},
			},
			want: true,
		},
		{
			name: "single RTG",
			fields: fields{
				RTG: 1,
			},
			args: args{
				c: Components{0, 1},
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
				c: Components{1, 1},
			},
			want: true,
		},
		{
			name: "extra RTG",
			fields: fields{
				Chip: 3,
				RTG:  3,
			},
			args: args{
				c: Components{1, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Floor{
				Chip: tt.fields.Chip,
				RTG:  tt.fields.RTG,
			}
			if got := f.IsSafeWithout(tt.args.c); got != tt.want {
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
			want: false,
		},
		{
			name: "patial overlap RTG",
			args: args{
				chip: 1,
				rtg:  3,
			},
			want: true,
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

func TestExample(t *testing.T) {
	const input = `The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

	floors, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if got := part1(floors); got != "11" {
		t.Errorf("part1() = %v, want %v", got, "11")
	}
}

func Test_normalize(t *testing.T) {
	type args struct {
		s State
	}
	tests := []struct {
		name string
		args args
		want State
	}{
		{
			name: "empty",
			args: args{
				s: State{},
			},
			want: State{},
		},
		{
			name: "single chip",
			args: args{
				s: State{
					Floors: [4]Floor{
						{
							Chip: 1 << 3,
							RTG:  1 << 3,
						},
					},
				},
			},
			want: State{
				Floors: [4]Floor{
					{
						Chip: 1,
						RTG:  1,
					},
				},
			},
		},
		{
			name: "multiple",
			args: args{
				s: State{
					Floors: [4]Floor{
						{1 << 4, 1 << 5},
						{1 << 6, 1 << 7},
						{1, 1},
						{1<<5 | 1<<7, 1<<4 | 1<<6},
					},
				},
			},
			want: State{
				Floors: [4]Floor{
					{1, 2},
					{4, 8},
					{16, 16},
					{10, 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_permute(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "one",
			args: args{
				n: 1,
			},
			want: []uint8{1},
		},
		{
			name: "two",
			args: args{
				n: 2,
			},
			want: []uint8{2},
		},
		{
			name: "three",
			args: args{
				n: 3,
			},
			want: []uint8{1, 2, 3},
		},
		{
			name: "seven",
			args: args{
				n: 7,
			},
			want: []uint8{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.n)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func permute(n uint8) []uint8 {
	perms := make([]uint8, 0)

	for i := uint8(0); i < 8; i++ {
		if n&(1<<i) == 0 {
			continue
		}
		perms = append(perms, 1<<i)
	}

	size := len(perms)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			perms = append(perms, perms[i]|perms[j])
		}
	}

	return perms
}

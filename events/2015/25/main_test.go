package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_offset(t *testing.T) {
	type args struct {
		row uint
		col uint
	}
	tests := []struct {
		args args
		want uint
	}{
		{
			args: args{},
			want: 0,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("(%d,%d)==%d", tt.args.col, tt.args.row, tt.want)
		t.Run(name, func(t *testing.T) {
			if got := offset(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve1(t *testing.T) {

	expect := [][]uint{
		{20151125, 18749137, 17289845, 30943339, 10071777, 33511524},
		{31916031, 21629792, 16929656, 7726640, 15514188, 4041754},
		{16080970, 8057251, 1601130, 7981243, 11661866, 16474243},
		{24592653, 32451966, 21345942, 9380097, 10600672, 31527494},
		{77061, 17552253, 28094349, 6899651, 9250759, 31663883},
		{33071741, 6796745, 25397450, 24659492, 1534922, 27995004},
	}

	for y, row := range expect {
		for x, v := range row {
			in := Input{Row: uint(y + 1), Col: uint(x + 1), Start: 20151125}

			got := solve1(in)

			if got != strconv.Itoa(int(v)) {
				t.Errorf("(%4d,%4d): got %s, want %d", x, y, got, v)
			}
		}
	}

}

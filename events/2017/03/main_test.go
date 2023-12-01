package main

import (
	"strconv"
	"testing"
)

func Test_part1(t *testing.T) {

	tests := [][2]int{
		{1, 0},
		{12, 3},
		{23, 2},
		{25, 4},
		{1024, 31},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt[0]), func(t *testing.T) {
			if got := part1(tt[0]); strconv.Itoa(tt[1]) != got {
				t.Errorf("part1() = %v, want %v", got, tt[1])
			}
		})
	}
}

func Test_coordToNum(t *testing.T) {
	tests := [][3]int{
		{0, 0, 1},
		{1, 0, 2},
		{1, 1, 3},
		{0, 1, 4},
		{-1, 1, 5},
		{-1, 0, 6},
		{-1, -1, 7},
		{0, -1, 8},
		{1, -1, 9},
		{2, -1, 10},
		{2, 0, 11},
		{2, 1, 12},
		{2, 2, 13},
		{1, 2, 14},
		{0, 2, 15},
		{-1, 2, 16},
		{-2, 2, 17},
		{-2, 1, 18},
		{-2, 0, 19},
		{-2, -1, 20},
		{-2, -2, 21},
		{-1, -2, 22},
		{0, -2, 23},
		{1, -2, 24},
		{2, -2, 25},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt[2]), func(t *testing.T) {
			if got := coordToNum(tt[0], tt[1]); tt[2] != got {
				t.Errorf("coordToNum() = %v, want %v", got, tt[2])
			}
		})
	}
}

func Test_numToCoord(t *testing.T) {
	tests := [][3]int{
		{0, 0, 1},
		{1, 0, 2},
		{1, 1, 3},
		{0, 1, 4},
		{-1, 1, 5},
		{-1, 0, 6},
		{-1, -1, 7},
		{0, -1, 8},
		{1, -1, 9},
		{2, -1, 10},
		{2, 0, 11},
		{2, 1, 12},
		{2, 2, 13},
		{1, 2, 14},
		{0, 2, 15},
		{-1, 2, 16},
		{-2, 2, 17},
		{-2, 1, 18},
		{-2, 0, 19},
		{-2, -1, 20},
		{-2, -2, 21},
		{-1, -2, 22},
		{0, -2, 23},
		{1, -2, 24},
		{2, -2, 25},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt[2]), func(t *testing.T) {
			gx, gy := numToCoord(tt[2])
			if tt[0] != gx || tt[1] != gy {
				t.Errorf("numToCoord() = %v, %v, want %v, %v", gx, gy, tt[0], tt[1])
			}
		})
	}
}

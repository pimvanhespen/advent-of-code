package main

import "testing"

func Test_part1(t *testing.T) {
	got := part1(Input{
		{ID: 1, Positions: 5, Offset: 4},
		{ID: 2, Positions: 2, Offset: 1},
	})

	if got != "5" {
		t.Errorf("part1() = %v, want %v", got, "5")
	}
}

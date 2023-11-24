package main

import "testing"

func Test_part1(t *testing.T) {
	got := part1([]byte("abc"))
	want := "22728"
	if got != want {
		t.Errorf("part1() = %v, want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	got := part2([]byte("abc"))
	want := "22551"
	if got != want {
		t.Errorf("part2() = %v, want %v", got, want)
	}
}

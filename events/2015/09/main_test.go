package main

import (
	"strings"
	"testing"
)

const input = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
`

func TestPart1(t *testing.T) {
	r, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	got := part1(r)
	want := 605

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	r, err := parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	got := part2(r)
	want := 982

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

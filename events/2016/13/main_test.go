package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	input := Input{
		MagicNumber: 10,
		Target:      Vec2{X: 7, Y: 4},
	}

	res := part1(input)

	if res != "11" {
		t.Errorf("Expected 11, got %s", res)
	}
}

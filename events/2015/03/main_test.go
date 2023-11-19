package main

import (
	"strings"
	"testing"
)

func Test_solve2(t *testing.T) {
	directions, err := parse(strings.NewReader(`^v^v^v^v^v`))
	if err != nil {
		t.Fatal(err)
	}

	num := solve2(directions)
	if num != 11 {
		t.Fatalf("Expected 11, got %d", num)
	}
}

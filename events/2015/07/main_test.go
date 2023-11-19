package main

import (
	"strings"
	"testing"
)

var input = `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`

func TestRun(t *testing.T) {

	expect := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}

	r := strings.NewReader(input)

	parsed, err := parse(r)
	if err != nil {
		t.Fatal(err)
	}

	result, err := eval(parsed.Wires)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range expect {

		n, ok := result[k]
		if !ok {
			t.Errorf("expected %s to exist", k)
			continue
		}

		if n != v {
			t.Errorf("%s: expected %d, got %d", k, v, n)
		}
	}
}

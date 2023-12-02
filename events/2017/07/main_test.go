package main

import (
	"strings"
	"testing"
)

const example = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func Test_part2(t *testing.T) {
	nodes, err := parse(strings.NewReader(example))
	if err != nil {
		t.Fatal(err)
	}

	got := part2(nodes)
	want := "60"

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

package main

import (
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"log"
	"maps"
	"strconv"
	"strings"
)

type Operation = string

const (
	And    = Operation("AND")
	Or     = Operation("OR")
	Not    = Operation("NOT")
	LShift = Operation("LSHIFT")
	RShift = Operation("RSHIFT")
	Set    = Operation("SET")
)

type Wire struct {
	Name      string
	Operation Operation
	Sources   []string
}

func NewWire(line string) (*Wire, error) {
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid line: %s", line)
	}

	formula, name := parts[0], parts[1]

	operation, sources, err := parseFormula(formula)
	if err != nil {
		return nil, err
	}

	w := &Wire{
		Name:      name,
		Operation: operation,
		Sources:   sources,
	}

	return w, nil
}

func (w *Wire) String() string {
	switch w.Operation {
	case And, Or, LShift, RShift:
		return fmt.Sprintf("%s %s %s -> %s", w.Sources[0], w.Operation, w.Sources[1], w.Name)
	case Not:
		return fmt.Sprintf("%s %s -> %s", w.Operation, w.Sources[0], w.Name)
	case Set:
		return fmt.Sprintf("%s -> %s", w.Sources[0], w.Name)
	default:
		return fmt.Sprintf("%v", *w)
	}
}

type Input struct {
	Wires map[string]*Wire
}

func main() {
	input, err := aoc.Get(7)
	if err != nil {
		log.Fatal(err)
	}

	m, err := parse(input)
	if err != nil {
		log.Fatal(err)
	}

	one := solve1(m)
	log.Println("Part one:", one)

	two := solve2(m)
	log.Println("Part two:", two)
}

func parse(reader io.Reader) (Input, error) {

	b, err := io.ReadAll(reader)
	if err != nil {
		return Input{}, err
	}

	lines := strings.Split(string(b), "\n")

	m := make(map[string]*Wire, len(lines))

	for _, line := range lines {

		if line == "" {
			continue
		}

		w, err := NewWire(line)
		if err != nil {
			return Input{}, err
		}

		m[w.Name] = w
	}

	return Input{Wires: m}, nil
}

func parseFormula(formula string) (Operation, []string, error) {
	switch {
	case strings.Contains(formula, And):

		return split(formula, And)
	case strings.Contains(formula, Or):
		return split(formula, Or)
	case strings.Contains(formula, Not):
		return split(formula, Not)
	case strings.Contains(formula, LShift):
		return split(formula, LShift)
	case strings.Contains(formula, RShift):
		return split(formula, RShift)
	default:
		return Set, []string{strings.TrimSpace(formula)}, nil
	}
}

func split(formula string, op Operation) (Operation, []string, error) {
	parts := strings.Split(formula, op)

	for i := len(parts) - 1; i >= 0; i-- {
		parts[i] = strings.TrimSpace(parts[i])
		if parts[i] == "" {
			parts = append(parts[:i], parts[i+1:]...)
		}
	}

	return op, parts, nil
}

func solve1(m Input) uint16 {

	wires := maps.Clone(m.Wires)

	a, ok := wires["a"]
	if !ok {
		log.Fatal("wire a not found")
	}

	results, err := eval(wires)
	if err != nil {
		log.Fatal(err)
	}

	return results[a.Name]
}

func solve2(i Input) uint16 {

	wires := maps.Clone(i.Wires)

	results, err := eval(wires)
	if err != nil {
		log.Fatal(err)
	}

	b := results["a"]

	wires = maps.Clone(i.Wires)

	wires["b"].Operation = Set
	wires["b"].Sources = []string{strconv.FormatUint(uint64(b), 10)}

	results, err = eval(wires)
	if err != nil {
		log.Fatal(err)
	}

	return results["a"]
}

func isNumber(s string) bool {
	if s[0] < '0' || s[0] > '9' {
		return false
	}

	return true
}

func parseUint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func keysInSet(keys []string, set map[string]struct{}) bool {
	for _, s := range keys {
		if isNumber(s) {
			continue
		}

		if _, ok := set[s]; !ok {
			return false
		}
	}
	return true
}

func get(m map[string]uint16, key string) (uint16, error) {
	if isNumber(key) {
		return parseUint16(key)
	}

	v, ok := m[key]
	if !ok {
		return 0, fmt.Errorf("could not find %q", key)
	}

	return v, nil
}

func get2(m map[string]uint16, left, right string) (uint16, uint16, error) {
	l, err := get(m, left)
	if err != nil {
		return 0, 0, err
	}

	r, err := get(m, right)
	if err != nil {
		return 0, 0, err
	}

	return l, r, nil
}

func eval(m map[string]*Wire) (map[string]uint16, error) {

	m = maps.Clone(m)

	var layers [][]*Wire

	classified := make(map[string]struct{}, len(m))
	result := make(map[string]uint16, len(m))

	for len(m) > 0 {
		var layer []*Wire

		for _, w := range m {
			if !keysInSet(w.Sources, classified) {
				continue
			}

			layer = append(layer, w)
			delete(m, w.Name)
		}

		if len(layer) == 0 {
			return nil, fmt.Errorf("could not classify any wires")
		}

		layers = append(layers, layer)
		for _, w := range layer {
			classified[w.Name] = struct{}{}
		}
	}

	for _, layer := range layers {
		for _, w := range layer {
			switch w.Operation {
			case Set:
				v, err := get(result, w.Sources[0])
				if err != nil {
					return nil, fmt.Errorf("%s = %s: %w", w.Name, w.Sources[0], err)
				}
				result[w.Name] = v

			case And:
				left, right, err := get2(result, w.Sources[0], w.Sources[1])
				if err != nil {
					return nil, fmt.Errorf("%s: %w", w, err)
				}

				result[w.Name] = left & right

			case Or:
				left, right, err := get2(result, w.Sources[0], w.Sources[1])
				if err != nil {
					return nil, fmt.Errorf("%s: %w", w, err)
				}

				result[w.Name] = left | right

			case LShift:
				left, right, err := get2(result, w.Sources[0], w.Sources[1])
				if err != nil {
					return nil, fmt.Errorf("%s: %w", w, err)
				}

				result[w.Name] = left << right

			case RShift:
				left, right, err := get2(result, w.Sources[0], w.Sources[1])
				if err != nil {
					return nil, fmt.Errorf("%s: %w", w, err)
				}

				result[w.Name] = left >> right

			case Not:
				v, err := get(result, w.Sources[0])
				if err != nil {
					return nil, fmt.Errorf("%s: %w", w, err)
				}

				result[w.Name] = ^v
			}
		}
	}

	return result, nil
}

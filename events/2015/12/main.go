package main

import (
	"encoding/json"
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"io"
	"regexp"
	"strconv"
)

func main() {
	r, err := aoc.Get(12)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(string(b)))
	fmt.Println("Part 2:", part2(string(b)))
}

var nums = regexp.MustCompile(`-?\d+`)

func part1(input string) int {
	matches := nums.FindAll([]byte(input), -1)

	var total int
	for _, match := range matches {
		n, _ := strconv.Atoi(string(match))
		total += n
	}
	return total
}

type T interface{}

func part2(input string) int {
	var v T
	err := json.Unmarshal([]byte(input), &v)
	if err != nil {
		panic(err)
	}

	return count(v)
}

func count(v any) int {
	if isRed(v) {
		return 0
	}

	switch x := v.(type) {
	case float64:
		return int(x)
	case string:
		return 0
	case []interface{}:
		total := 0
		for _, vv := range v.([]interface{}) {
			if vv == "red" {
				continue
			}
			total += count(vv)
		}
		return total
	case map[string]interface{}:
		for k, vv := range v.(map[string]interface{}) {
			if k == "red" || vv == "red" {
				return 0
			}
		}

		total := 0
		for _, vv := range v.(map[string]interface{}) {
			total += count(vv)
		}
		return total
	default:
		fmt.Printf("Unknown type %T\n", v)
		return 0
	}

	return 0
}

func isRed(v any) bool {
	s, ok := v.(string)
	return ok && s == "red"
}

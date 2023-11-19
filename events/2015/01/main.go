package main

import (
	"fmt"
	"io"
	"log"

	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
)

func main() {
	input, err := aoc.Get(1)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solve1(b))
	fmt.Println("Part 2:", solve2(b))

}

func solve1(input []byte) int {
	var floor int
	for _, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	return floor
}

func solve2(input []byte) int {
	var floor int
	for i, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 {
			return i + 1
		}
	}
	return floor
}

package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
)

type Input = []string

func main() {
	event := aoc.New(2023, 1, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ReadLines(r)
}

func part1(input Input) string {

	var sum int

	for _, line := range input {
		var nums []int
		for _, char := range line {
			if char >= '0' && char <= '9' {
				nums = append(nums, int(char)-'0')
			}
		}
		sum += number(nums)
	}

	return fmt.Sprintf("%d", sum)
}

func digits2(line string) []int {
	var nums []int

	for i, c := range line {
		switch c {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			nums = append(nums, int(c-'0'))
		case 'o':
			if i+2 < len(line) && line[i:i+3] == "one" {
				nums = append(nums, 1)
			}
		case 't':
			if i+2 < len(line) && line[i:i+3] == "two" {
				nums = append(nums, 2)
			}
			if i+4 < len(line) && line[i:i+5] == "three" {
				nums = append(nums, 3)
			}
		case 'f':
			if i+3 < len(line) && line[i:i+4] == "four" {
				nums = append(nums, 4)
			}
			if i+3 < len(line) && line[i:i+4] == "five" {
				nums = append(nums, 5)
			}
		case 's':
			if i+2 < len(line) && line[i:i+3] == "six" {
				nums = append(nums, 6)
			}
			if i+4 < len(line) && line[i:i+5] == "seven" {
				nums = append(nums, 7)
			}
		case 'e':
			if i+4 < len(line) && line[i:i+5] == "eight" {
				nums = append(nums, 8)
			}
		case 'n':
			if i+3 < len(line) && line[i:i+4] == "nine" {
				nums = append(nums, 9)
			}
		case 'z':
			if i+3 < len(line) && line[i:i+4] == "zero" {
				nums = append(nums, 0)
			}
		}
	}

	return nums
}

func number(digits []int) int {
	return 10*digits[0] + digits[len(digits)-1]
}

func part2(input Input) string {

	var sum int

	for _, line := range input {
		sum += number(digits2(line))
	}

	return fmt.Sprintf("%d", sum)
}

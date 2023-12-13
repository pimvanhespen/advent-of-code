package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

type Input []Record

type Record struct {
	Row    []byte
	Broken []int
}

func main() {
	event := aoc.New(2023, 12, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	return aoc.ParseLines(r, func(s string) (Record, error) {
		parts := strings.Split(s, " ")
		row := []byte(parts[0])
		broken, err := aoc.Ints(parts[1])
		if err != nil {
			return Record{}, fmt.Errorf("parse broken: %w", err)
		}
		return Record{
			Row:    row,
			Broken: broken,
		}, nil
	})
}

func part1(input Input) string {
	var total int
	m := make(map[key]int)
	for _, r := range input {
		total += permuteCached(m, r.Row, r.Broken)
	}
	return fmt.Sprint(total)
}

func part2(input Input) string {

	m := make(map[key]int, 100_000)

	var tot int
	for _, r := range input {
		in := unfold(r)

		n := permuteCached(m, in.Row, in.Broken)
		tot += n
	}

	return fmt.Sprint(tot)
}

func fits(row []byte, index, broken int) bool {
	if index+broken-1 >= len(row) {
		return false
	}

	// check that the last byte is not a #
	if index+broken < len(row) {
		if row[index+broken] == '#' {
			return false
		}
	}

	for i := index; i < index+broken; i++ {
		if row[i] == '.' {
			return false
		}
	}
	return true
}

func size(broken []int) int {
	if len(broken) == 0 {
		return 0
	}
	if len(broken) == 1 {
		return broken[0]
	}
	sum := broken[0]
	for _, b := range broken[1:] {
		sum += 1 + b
	}
	return sum
}

func permuteCached(m map[key]int, row []byte, nums []int) int {

	mapKey := newKey(row, nums)
	if v, ok := m[mapKey]; ok {
		return v
	}

	if len(nums) == 0 {
		if len(row) > 0 {
			// Make sure that none of the remaining bytes are #
			if bytes.ContainsAny(row, "#") {
				m[mapKey] = 0
				return 0
			}
		}
		m[mapKey] = 1
		return 1
	}

	var permutations int
	for i := 0; i < len(row); i++ {
		if row[i] == '.' {
			continue
		}

		// Does the allocation fit at this position?
		if !fits(row[i:], 0, nums[0]) {
			if row[i] == '#' || size(nums) >= len(row[i:]) {
				break // permutation is invalid - stop
			}
			continue
		}

		// Does the remainder _POSSIBLY_ fit in the size of slice?
		if size(nums[1:]) >= len(row[i:]) {
			break
		}

		// Lock bytes in place
		newSize := i + nums[0]

		// try to add a 'whitespace' unit
		if newSize < len(row) {
			newSize++
		}

		// Recurse
		permutations += permuteCached(m, row[newSize:], nums[1:])
		if row[i] == '#' {
			break
		}
	}
	m[mapKey] = permutations
	return permutations
}

type key struct {
	row  string
	nums string
}

func newKey(row []byte, nums []int) key {
	return key{
		row:  string(row),
		nums: concat(nums),
	}
}

func concat(nums []int) string {
	var b strings.Builder
	for _, n := range nums {
		b.WriteString(strconv.Itoa(n))
	}
	return b.String()
}

func unfold(record Record) Record {

	copies := make([][]byte, 5)
	for i := range copies {
		copies[i] = make([]byte, len(record.Row))
		copy(copies[i], record.Row)
	}

	row := bytes.Join(copies, []byte{'?'})

	nums := make([]int, 0, 5*len(record.Broken))
	for i := 0; i < 5; i++ {
		nums = append(nums, record.Broken...)
	}

	return Record{
		Row:    row,
		Broken: nums,
	}
}

package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
)

type Input = []byte

func main() {
	event := aoc.New(2016, 14, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return Input{}, fmt.Errorf("read input: %w", err)
	}

	return Input(bytes.TrimSpace(b)), nil
}

func hasTriplet(s string) (byte, bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i], true
		}
	}

	return 0, false
}

func hasQuintuplet(s string, char byte) bool {
	for i := 0; i < len(s)-4; i++ {
		if s[i] == char && s[i+1] == char && s[i+2] == char && s[i+3] == char && s[i+4] == char {
			return true
		}
	}
	return false
}

func hashSalt(salt []byte, index uint) string {
	data := []byte(fmt.Sprintf("%s%d", string(salt), index))

	return fmt.Sprintf("%x", md5.Sum(data))
}

func part1(salt []byte) string {

	keys := solve(salt, hashSalt)

	return fmt.Sprint(keys[len(keys)-1].index)
}

type Hasher func(salt []byte, index uint) string

type key struct {
	index uint
	hash  string
}

func solve(salt []byte, fn Hasher) []key {

	var keys []key

	seen := make([]string, 1, 1001)
	for index := uint(0); ; index++ {
		seen = seen[1:]

		if len(seen) == 0 {
			seen = append(seen, fn(salt, index))
		}

		hs := seen[0]

		char, ok := hasTriplet(hs)
		if !ok {
			continue
		}

		for i := uint(1); i <= 1000; i++ {
			if i >= uint(len(seen)) {
				seen = append(seen, fn(salt, index+i))
			}

			s := seen[i]
			if !hasQuintuplet(s, char) {
				continue
			}

			keys = append(keys, key{index, hs})
			if len(keys) == 64 {
				return keys
			}

			break
		}
	}
}

func hashSalt2016(salt []byte, index uint) string {
	hash := hashSalt(salt, index)
	for i := 0; i < 2016; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	return hash
}

func part2(input []byte) string {
	keys := solve(input, hashSalt2016)

	return fmt.Sprint(keys[len(keys)-1].index)
}

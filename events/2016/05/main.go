package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"log"
	"strconv"
)

type Input = []byte

func main() {
	event := aoc.New(2016, 5, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

func parse(r io.Reader) (Input, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	b = bytes.TrimSpace(b)
	return b, nil
}

func part1(input Input) string {
	var password [8]byte
	var index int

	hash := md5.New()
	data := make([]byte, 64)

	var offset int
	for ; ; index++ {
		n := number(index)
		hash.Reset()
		hash.Write(input)
		hash.Write(n)
		hex.Encode(data[:], hash.Sum(nil))

		if data[0] == 48 && data[1] == 48 && data[2] == 48 && data[3] == 48 && data[4] == 48 {
			password[offset] = data[5]
			log.Printf("Character %c found at %d (%s)", data[5], index, string(password[:]))

			offset++
			if offset == len(password) {
				break
			}
		}
	}

	return string(password[:])
}

func number(n int) []byte {
	return []byte(strconv.Itoa(n))
}

func part2(input Input) string {
	var password [8]byte
	var mask [8]bool
	var index int

	hash := md5.New()
	data := make([]byte, 64)

	var offset int
	for ; ; index++ {
		n := number(index)
		hash.Reset()
		hash.Write(input)
		hash.Write(n)
		hex.Encode(data[:], hash.Sum(nil))

		if data[0] == 48 && data[1] == 48 && data[2] == 48 && data[3] == 48 && data[4] == 48 {

			pos := data[5] - '0'

			if int(pos) >= len(mask) {
				log.Printf("Skipping %d", pos)
				continue
			}

			if mask[pos] {
				log.Printf("Skipping %d; already set", pos)
				continue
			}

			mask[pos] = true
			ch := data[6]
			password[pos] = ch

			log.Printf("Character %c found at %d (%s)", ch, index, string(password[:]))

			offset++
			if offset == len(password) {
				break
			}
		}
	}

	return string(password[:])
}

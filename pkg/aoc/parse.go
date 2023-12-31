package aoc

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var IgnoreLine = errors.New("ignore line")

func ParseLines[T any](reader io.Reader, fn func(string) (T, error)) ([]T, error) {
	scanner := bufio.NewScanner(reader)
	var result []T
	for scanner.Scan() {
		line := scanner.Text()
		value, err := fn(line)
		if err != nil {
			if errors.Is(err, IgnoreLine) {
				continue
			}
			return nil, err
		}
		result = append(result, value)
	}
	return result, scanner.Err()
}

func ReadLines(reader io.Reader) ([]string, error) {
	return ParseLines(reader, func(line string) (string, error) {
		return line, nil
	})
}

func ReadAll(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return bytes.TrimSpace(b), nil
}

func ParseInput[T any](reader io.Reader, fn func(string) (T, error)) (T, error) {
	b, err := ReadAll(reader)
	if err != nil {
		var zero T
		return zero, err
	}
	return fn(string(b))
}

func ReadMap(reader io.Reader) ([][]byte, error) {
	return ParseLines(reader, func(line string) ([]byte, error) {
		return []byte(line), nil
	})
}

var numberReg = regexp.MustCompile(`(-?\d+)`)

func Ints(s string) ([]int, error) {
	matches := numberReg.FindAllStringSubmatch(s, -1)
	nums := make([]int, len(matches))
	for i, match := range matches {
		var err error
		nums[i], err = strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %w", match[1], err)
		}
	}
	return nums, nil
}

type Grid struct {
	Width, Height int
	Data          []byte
}

func (g Grid) Get(x, y int) (byte, bool) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return 0, false
	}
	return g.Data[y*g.Width+x], true
}

func (g Grid) Set(x, y int, value byte) {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return
	}
	g.Data[y*g.Width+x] = value
}

func (g Grid) String() string {
	var buf bytes.Buffer
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			buf.WriteByte(g.Data[y*g.Width+x])
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func ParseGrid(reader io.Reader) (Grid, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return Grid{}, err
	}

	b = bytes.TrimSpace(b)
	width := bytes.IndexByte(b, '\n')
	height := bytes.Count(b, []byte{'\n'}) + 1

	return Grid{
		Width:  width,
		Height: height,
		Data:   bytes.ReplaceAll(b, []byte{'\n'}, nil),
	}, nil
}

func ParseGrid2D(reader io.Reader) ([][]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	b = bytes.TrimSpace(b)
	lines := bytes.Split(b, []byte{'\n'})
	return lines, nil
}

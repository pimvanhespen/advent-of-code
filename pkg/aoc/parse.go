package aoc

import (
	"bufio"
	"bytes"
	"errors"
	"io"
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

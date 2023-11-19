package main

import "strings"

const input = "cqjxjnds"

func main() {
	println("Part 1:", next(input))
	println("Part 2:", next(next(input)))
}

func next(input string) string {
	input = increment(input)
	for !valid(input) {
		input = increment(input)
	}
	return input
}

func increment(input string) string {
	chars := []byte(input)
	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] == 'z' {
			chars[i] = 'a'
		} else {
			chars[i]++
			break
		}
	}
	return string(chars)
}

func valid(input string) bool {
	return !checkInvalidChars(input) && checkTriplet(input) && checkPairs(input)
}

func checkTriplet(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i+1] == input[i]+1 && input[i+2] == input[i]+2 {
			return true
		}
	}
	return false
}

func checkInvalidChars(input string) bool {
	return strings.ContainsAny(input, "iol")
}

func checkPairs(input string) bool {
	pairs := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i+1] == input[i] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

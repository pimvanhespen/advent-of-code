package main

import (
	"fmt"
	"strings"
	"time"
)

type Sequence = string

var memo = map[string]string{
	"1":   "11",
	"2":   "12",
	"3":   "13",
	"11":  "21",
	"111": "31",
	"22":  "22",
	"222": "32",
	"33":  "23",
	"333": "33",
}

func lookAndSay(s string) string {
	sb := strings.Builder{}

	for i := 0; i < len(s); i++ {
		j := i + 1
		for j < len(s) && s[i] == s[j] {
			j++
		}
		sb.WriteString(memo[s[i:j]])
		i = j - 1
	}

	return sb.String()
}

const input = `1113222113`

func main() {
	t1 := time.Now()
	sequence := Sequence(input)
	for i := 0; i < 40; i++ {
		sequence = lookAndSay(sequence)
	}
	fmt.Println(len(sequence))

	for i := 0; i < 10; i++ {
		sequence = lookAndSay(sequence)
	}
	fmt.Println(len(sequence))
	d1 := time.Now().Sub(t1)
	fmt.Println(d1)
}

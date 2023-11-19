package main

import (
	"fmt"
	"github.com/pimvanhespen/aoc/2015/pkg/aoc"
	"strconv"
)

func main() {
	reader, err := aoc.Get(17)
	if err != nil {
		panic(err)
	}

	containers, err := aoc.ParseLines(reader, strconv.Atoi)
	if err != nil {
		panic(err)
	}

	perms := permutations(containers, 150)

	// Part 1
	fmt.Println("Part 1:", len(perms))

	// Part 2
	fmt.Println("Part 2:", least(perms))
}

func least(perms [][]int) int {
	shortest := len(perms[0])

	for _, perm := range perms {
		if len(perm) < shortest {
			shortest = len(perm)
		}
	}

	var count int
	for _, perm := range perms {
		if len(perm) == shortest {
			count++
		}
	}

	return count
}

func permutations(containers []int, target int) [][]int {
	var result [][]int

	for i, container := range containers {
		if container == target {
			result = append(result, []int{container})
		} else if container < target {
			for _, permutation := range permutations(containers[i+1:], target-container) {
				result = append(result, append([]int{container}, permutation...))
			}
		}
	}
	return result
}

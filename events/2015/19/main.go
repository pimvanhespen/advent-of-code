package main

import (
	"bufio"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"github.com/pimvanhespen/advent-of-code/pkg/datastructures/heap"
	"io"
	"log"
	"math"
	"strings"
	"sync"
	"time"
)

type Data struct {
	replacements map[string][]string
	molecule     string
}

func main() {
	reader, err := aoc.NewChallenge(2015, 19).Input()
	if err != nil {
		panic(err)
	}

	data, err := parse(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", solve1(data))
	fmt.Println("Part 2:", solve2(data))
}

func parse(reader io.Reader) (Data, error) {
	data := Data{
		replacements: make(map[string][]string),
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return Data{}, err
		}

		text := scanner.Text()
		if text == "" {
			break
		}
		parts := strings.Split(scanner.Text(), " => ")

		data.replacements[parts[0]] = append(data.replacements[parts[0]], parts[1])
	}

	scanner.Scan()
	data.molecule = scanner.Text()

	return data, scanner.Err()
}

// todo: learn how to solve this
func solve1(data Data) int {

	source := data.molecule

	seen := make(map[string]bool)

	for i := 0; i < len(source); i++ {
		for a, replacements := range data.replacements {
			if !strings.HasPrefix(source[i:], a) {
				continue
			}

			for _, replacement := range replacements {
				m := source[:i] + replacement + source[i+len(a):]
				seen[m] = true
			}

			i += len(a) - 1
		}
	}

	return len(seen)
}

func solve2(data Data) int {

	type state struct {
		molecule string
		steps    int
	}

	least := state{data.molecule, math.MaxInt}

	mh := heap.NewMin[int, state]()

	mh.Push(state{data.molecule, 0}, 0)

	seen := make(map[string]bool)

	tasks := make(chan state)
	results := make(chan state, 100)

	reduced := make(chan state, 100)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		lastPop := time.Now()

	outer:
		for {
			select {
			case s := <-results:
				if seen[s.molecule] {
					continue
				}

				seen[s.molecule] = true
				mh.Push(s, len(s.molecule))
			default:
				if mh.Empty() {
					if time.Since(lastPop) > time.Second {
						break outer
					}
					continue outer
				}
				s := mh.Pop()
				tasks <- s
			}
		}
		close(tasks)
	}()

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				log.Println("worker done")
			}()
			for task := range tasks {
				if task.steps > least.steps {
					continue
				}

				if task.molecule == "e" {
					reduced <- task
				}

				for a, replacements := range data.replacements {
					for _, replacement := range replacements {

						if !strings.Contains(task.molecule, replacement) {
							continue
						}

						m := strings.Replace(task.molecule, replacement, a, 1)
						s := state{m, task.steps + 1}

						results <- s
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
		close(reduced)
	}()

	for s := range reduced {
		log.Println(s)
		if s.steps < least.steps {
			least = s
		}
	}

	return least.steps
}

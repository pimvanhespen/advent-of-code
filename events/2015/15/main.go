package main

import (
	"bytes"
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
	"strconv"
)

type Ingredient struct {
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

type Fractions []int

type Recipe struct {
	Ingredients []Ingredient
}

func Score(ingredients []Ingredient, fractions Fractions) int {
	var a, b, c, d int

	for i, ingredient := range ingredients {
		amount := fractions[i]
		a += amount * ingredient.Capacity
		b += amount * ingredient.Durability
		c += amount * ingredient.Flavor
		d += amount * ingredient.Texture
	}

	if a < 0 || b < 0 || c < 0 || d < 0 {
		return 0
	}

	return a * b * c * d
}

func Calories(ingredients []Ingredient, fracts Fractions) int {
	total := 0
	for i, ingredient := range ingredients {
		total += fracts[i] * ingredient.Calories
	}
	return total
}

func main() {
	reader, err := aoc.NewChallenge(2015, 15).Input()
	if err != nil {
		panic(err)
	}

	ingredients, err := parseIngredients(reader)
	if err != nil {
		panic(err)
	}

	fracts := getOptimalRecipe(ingredients)

	fmt.Println("Part 1:", Score(ingredients, fracts))

	fracts2 := getOptimalRecipe(ingredients, func(ingredients []Ingredient, fracts Fractions) bool {
		return Calories(ingredients, fracts) == 500
	})
	fmt.Println("Part 2:", Score(ingredients, fracts2))
}

type KeepFunc func([]Ingredient, Fractions) bool

func getOptimalRecipe(ingredients []Ingredient, keepFuncs ...KeepFunc) Fractions {

	// This is a naive brute0force approach. If laptop == slow, speed it up.
	// Calculate all possble fractions for the ingredients
	fracts := permute(len(ingredients), 100)

	high := 0
	var set Fractions

outer:
	for i, fract := range fracts {

		for _, keep := range keepFuncs {
			if !keep(ingredients, fract) {
				continue outer
			}
		}

		score := Score(ingredients, fract)
		if score > high {
			high = score
			set = fracts[i]
		}
	}

	return set
}

// permute bruteforce all possible fractions for the ingredients
func permute(n, total int) []Fractions {
	if n == 1 {
		return []Fractions{{total}}
	}

	var result []Fractions
	for i := 0; i <= total; i++ {
		for _, p := range permute(n-1, total-i) {
			result = append(result, append(Fractions{i}, p...))
		}
	}

	return result
}

func parseIngredients(reader io.Reader) ([]Ingredient, error) {

	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	b = bytes.TrimSuffix(b, []byte("\n"))
	lines := bytes.Split(b, []byte("\n"))

	ingredients := make([]Ingredient, len(lines))
	for i, l := range lines {
		fields := bytes.Fields(l)

		//name := bytes.TrimSuffix(fields[0], []byte(":"))

		cs := string(bytes.TrimSuffix(fields[2], []byte(",")))
		capacity, err := strconv.Atoi(cs)
		if err != nil {
			return nil, fmt.Errorf("could not parse capacity: %q: %w", cs, err)
		}

		durability, err := strconv.Atoi(string(bytes.TrimSuffix(fields[4], []byte(","))))
		if err != nil {
			return nil, fmt.Errorf("could not parse durability: %w", err)
		}

		flavor, err := strconv.Atoi(string(bytes.TrimSuffix(fields[6], []byte(","))))
		if err != nil {
			return nil, fmt.Errorf("could not parse flavor: %w", err)
		}

		texture, err := strconv.Atoi(string(bytes.TrimSuffix(fields[8], []byte(","))))
		if err != nil {
			return nil, fmt.Errorf("could not parse texture: %w", err)
		}

		calories, err := strconv.Atoi(string(bytes.TrimSuffix(fields[10], []byte(","))))
		if err != nil {
			return nil, fmt.Errorf("could not parse calories: %w", err)
		}

		ingredients[i] = Ingredient{
			Capacity:   capacity,
			Durability: durability,
			Flavor:     flavor,
			Texture:    texture,
			Calories:   calories,
		}

	}

	return ingredients, nil
}

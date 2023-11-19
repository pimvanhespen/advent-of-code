package main

import (
	"strings"
	"testing"
)

func Test_parseIngredients(t *testing.T) {
	const input = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
`

	expect := []Ingredient{
		{Name: "Butterscotch", Capacity: -1, Durability: -2, Flavor: 6, Texture: 3, Calories: 8},
		{Name: "Cinnamon", Capacity: 2, Durability: 3, Flavor: -2, Texture: -1, Calories: 3},
	}

	got, err := parseIngredients(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(expect) {
		t.Fatalf("expected %d ingredients, got %d", len(expect), len(got))
	}

	for i := range expect {
		if got[i] != expect[i] {
			t.Errorf("expected ingredient %d to be %v, got %v", i, expect[i], got[i])
		}
	}
}

func TestRecipe_Score(t *testing.T) {
	recipe := Recipe{
		Ingredients: []Ingredient{
			{Name: "Butterscotch", Capacity: -1, Durability: -2, Flavor: 6, Texture: 3, Calories: 8},
			{Name: "Cinnamon", Capacity: 2, Durability: 3, Flavor: -2, Texture: -1, Calories: 3},
		},
		Fractions: []uint8{44, 56},
	}

	const expect = 62842880

	got := recipe.Score()

	if got != expect {
		t.Errorf("expected score to be %d, got %d", expect, got)
	}
}

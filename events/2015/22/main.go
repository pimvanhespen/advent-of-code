package main

import (
	"fmt"
	"github.com/pimvanhespen/aoc/2015/days/22/wizsim"
	"math"
)

func main() {

	// enable for debug output
	wizsim.Logging = false

	player := wizsim.Player{
		Health: 50,
		Mana:   500,
	}

	boss := wizsim.Boss{
		Health: 55,
		Damage: 8,
	}

	g := wizsim.NewGame(player, boss)

	one := solve1(g, wizsim.Shield, wizsim.Poison, wizsim.MagicMissile, wizsim.Recharge)
	fmt.Printf("1): %d\n", one)

	g = wizsim.NewGame(player, boss)

	two := solve2(g, wizsim.MagicMissile, wizsim.Drain, wizsim.Shield, wizsim.Poison, wizsim.Recharge)
	fmt.Printf("2): %d\n", two)
}

func solve1(initial *wizsim.Game, spells ...*wizsim.Spell) int {
	return solve(initial, spells...)
}

func solve2(initial *wizsim.Game, spells ...*wizsim.Spell) int {
	initial.HardMode = true
	return solve(initial, spells...)
}

func solve(initial *wizsim.Game, spells ...*wizsim.Spell) int {
	least := math.MaxInt
	leastGame := initial

	queue := make([]*wizsim.Game, 0, 1<<10) // 1024
	queue = append(queue, initial)

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]

		if head.Player.Consumed > least {
			continue
		}

		for _, spell := range spells {

			if spell.Cost+head.Player.Consumed >= least {
				continue
			}

			next := wizsim.Play(head, spell)
			switch next.State {
			case wizsim.Crash, wizsim.BossWon:
				continue
			case wizsim.PlayerWon:
				if next.Player.Consumed < least {
					least = next.Player.Consumed
					leastGame = next
				}
			case wizsim.Running:
				queue = append(queue, next)
			}
		}
	}

	replay(initial, leastGame.Invocations)

	return least
}

func replay(initial *wizsim.Game, invocations []wizsim.Invocation) {
	for _, i := range invocations {
		initial = wizsim.Play(initial, i.Spell)
	}
	if initial.State != wizsim.PlayerWon {
		initial.PlayerStart()
	}

	if initial.State != wizsim.PlayerWon {
		panic("did not win")
	}
}

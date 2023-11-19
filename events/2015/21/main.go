package main

import "sort"

// Day 21: RPG Simulator 20XX

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

var weapons = []Item{
	{8, 4, 0},  // Dagger
	{10, 5, 0}, // Shortsword
	{25, 6, 0}, // Warhammer
	{40, 7, 0}, // Longsword
	{74, 8, 0}, // Greataxe
}

var armors = []Item{
	{13, 0, 1},  // Leather
	{31, 0, 2},  // Chainmail
	{53, 0, 3},  // Splintmail
	{75, 0, 4},  // Bandedmail
	{102, 0, 5}, // Platemail
}

var rings = []Item{
	{0, 0, 0},   // No ring
	{0, 0, 0},   // No ring
	{25, 1, 0},  // Damage +1
	{50, 2, 0},  // Damage +2
	{100, 3, 0}, // Damage +3
	{20, 0, 1},  // Defense +1
	{40, 0, 2},  // Defense +2
	{80, 0, 3},  // Defense +3
}

type Character struct {
	HP    int
	DMG   int
	Armor int
}

func main() {
	boss := Character{103, 9, 2}
	player := Character{100, 0, 0}

	least := 1000

	for _, weapon := range weapons {
		if weapon.Cost >= least {
			continue
		}
		for _, armor := range armors {
			if weapon.Cost+armor.Cost >= least {
				continue
			}
			for i, ring := range rings[:len(rings)-1] {
				if weapon.Cost+armor.Cost+ring.Cost >= least {
					continue
				}
				for _, ring2 := range rings[i+1:] {
					cost := weapon.Cost + armor.Cost + ring.Cost + ring2.Cost
					if cost >= least {
						continue
					}

					player.DMG = weapon.Damage + ring.Damage + ring2.Damage
					player.Armor = armor.Armor + ring.Armor + ring2.Armor
					if PlayerWins(player, boss) {
						if cost < least {
							least = cost
						}
					}
				}
			}
		}
	}

	println(least)

	most := 0

	sortMost(rings)
	sortMost(weapons)
	sortMost(armors)

	for _, weapon := range weapons {
		if weapon.Cost+180+102 <= most {
			continue
		}
		for _, armor := range armors {
			if weapon.Cost+armor.Cost+180 <= most {
				continue
			}
			for i, ring := range rings[:len(rings)-1] {
				if weapon.Cost+armor.Cost+ring.Cost+80 <= most {
					continue
				}
				for _, ring2 := range rings[i+1:] {
					cost := weapon.Cost + armor.Cost + ring.Cost + ring2.Cost
					if cost <= most {
						continue
					}

					player.DMG = weapon.Damage + ring.Damage + ring2.Damage
					player.Armor = armor.Armor + ring.Armor + ring2.Armor
					if !PlayerWins(player, boss) {
						if cost > most {
							most = cost
						}
					}
				}
			}
		}
	}

	println(most)
}

func sortLeast(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Cost < items[j].Cost
	})
}

func sortMost(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Cost > items[j].Cost
	})
}

func maxCost(items []Item) int {
	most := 0
	for _, item := range items {
		if item.Cost > most {
			most = item.Cost
		}
	}
	return most
}

func PlayerWins(Player, Boss Character) bool {
	playerDMG := Player.DMG - Boss.Armor
	bossDMG := Boss.DMG - Player.Armor
	if playerDMG < 1 {
		return false
	}

	if bossDMG < 1 {
		return true
	}

	playerTurns := Boss.HP / playerDMG
	if Boss.HP%playerDMG != 0 {
		playerTurns++
	}

	bossTurns := Player.HP / bossDMG
	if Player.HP%bossDMG != 0 {
		bossTurns++
	}

	return playerTurns <= bossTurns
}

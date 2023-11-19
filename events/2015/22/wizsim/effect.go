package wizsim

import "fmt"

type Effects struct {
	Init Effect
	Tick Effect
	Stop Effect
}

type EffectFunc func(player *Player, boss *Boss)

func (f EffectFunc) Apply(player *Player, boss *Boss) {
	f(player, boss)
}

type Effect interface {
	Apply(player *Player, boss *Boss)
}

func CombineEffects(effects ...Effect) Effect {
	return EffectFunc(func(player *Player, boss *Boss) {
		for _, effect := range effects {
			effect.Apply(player, boss)
		}
	})
}

func Damage(amount int) Effect {
	return EffectFunc(func(player *Player, boss *Boss) {
		if Logging {
			fmt.Printf("deals %d damage; ", amount)
		}
		boss.Health -= amount
	})
}

func Heal(amount int) Effect {
	return EffectFunc(func(player *Player, boss *Boss) {
		if Logging {
			fmt.Printf("heals %d hit points; ", amount)
		}
		player.Health += amount
	})
}

func Armor(amount int) Effect {
	return EffectFunc(func(player *Player, boss *Boss) {
		if Logging {
			if amount < 0 {
				fmt.Printf("decreses armor by %d; ", -amount)
			} else {
				fmt.Printf("increases armor by %d; ", amount)
			}
		}
		player.Armor += amount
	})
}

func Mana(amount int) Effect {
	return EffectFunc(func(player *Player, boss *Boss) {
		if Logging {
			fmt.Printf("restores %d mana; ", amount)
		}
		player.Mana += amount
	})
}

type none struct{}

func (*none) Apply(player *Player, boss *Boss) {}

func None() Effect {
	return &none{}
}

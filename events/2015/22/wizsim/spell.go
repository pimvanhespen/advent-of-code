package wizsim

import "fmt"

var (
	MagicMissile = &Spell{
		Name:  "MagicMissile",
		Cost:  53,
		Turns: 0,
		effects: Effects{
			Init: Damage(4),
			Tick: None(),
			Stop: None(),
		},
	}
	Drain = &Spell{
		Name:  "Drain",
		Cost:  73,
		Turns: 0,
		effects: Effects{
			Init: CombineEffects(Damage(2), Heal(2)),
			Tick: None(),
			Stop: None(),
		},
	}
	Shield = &Spell{
		Name:  "Shield",
		Cost:  113,
		Turns: 6,
		effects: Effects{
			Init: Armor(7),
			Tick: None(),
			Stop: Armor(-7),
		},
	}
	Poison = &Spell{
		Name:  "Poison",
		Cost:  173,
		Turns: 6,
		effects: Effects{
			Init: None(),
			Tick: Damage(3),
			Stop: None(),
		},
	}
	Recharge = &Spell{
		Name:  "Recharge",
		Cost:  229,
		Turns: 5,
		effects: Effects{
			Init: None(),
			Tick: Mana(101),
			Stop: None(),
		},
	}
)

type Spell struct {
	Name    string
	Cost    int
	Turns   int
	effects Effects
}

func (s *Spell) Init(player *Player, boss *Boss) {
	if Logging {
		fmt.Printf("Player casts %s. ", s.Name)
	}
	s.effects.Init.Apply(player, boss)
	if Logging {
		fmt.Println()
	}
}

func (s *Spell) Tick(player *Player, boss *Boss) {
	if _, ok := s.effects.Tick.(*none); ok {
		return
	}

	if Logging {
		fmt.Printf("%s ", s.Name)
	}
	s.effects.Tick.Apply(player, boss)
	if Logging {
		fmt.Println()
	}
}

func (s *Spell) Stop(player *Player, boss *Boss) {
	if _, ok := s.effects.Stop.(*none); ok {
		return
	}

	if Logging {
		fmt.Printf("Effect %s wears off. ", s.Name)
	}
	s.effects.Stop.Apply(player, boss)
	if Logging {
		fmt.Println()
	}
}

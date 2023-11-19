package wizsim

import "fmt"

type State uint8

var Logging = false

const (
	unknown State = iota
	Crash
	Running
	PlayerWon
	BossWon
)

type Player struct {
	Health   int
	Mana     int
	Armor    int
	Consumed int
}

func (p Player) String() string {
	return fmt.Sprintf("Player has %d hit points, %d armor, %d mana", p.Health, p.Armor, p.Mana)
}

type Boss struct {
	Health int
	Damage int
}

func (b Boss) String() string {
	return fmt.Sprintf("Boss has %d hits points", b.Health)
}

type Invocation struct {
	Spell *Spell
	Turns int
}

type Game struct {
	Turn        int
	Player      Player
	Boss        Boss
	Invocations []Invocation
	State       State
	HardMode    bool
}

func NewGame(player Player, boss Boss) *Game {
	return &Game{
		Player: player,
		Boss:   boss,
		State:  Running,
	}
}

func Play(game *Game, spell *Spell) *Game {

	game = game.copy()

	game.PlayerStart()
	if game.over() {
		return game
	}

	game.PlayerCast(spell)
	if game.over() {
		return game
	}

	game.BossTurn()
	if game.over() {
		return game
	}

	return game
}

func (g *Game) PlayerStart() {
	g.startTurn("Player")
}

func (g *Game) PlayerCast(spell *Spell) {

	if !g.isSpellAllowed(spell) {
		g.State = Crash
		return
	}

	g.Invocations = append(g.Invocations, Invocation{Spell: spell, Turns: spell.Turns})
	g.Player.Mana -= spell.Cost
	g.Player.Consumed += spell.Cost

	spell.Init(&g.Player, &g.Boss)
	if g.over() {
		return
	}
}

func (g *Game) over() bool {
	if g.State != Running {
		return true
	}

	if g.Boss.Health <= 0 {
		if Logging {
			fmt.Println("Boss dies")
		}

		g.State = PlayerWon
		return true
	}

	if g.Player.Health <= 0 {
		if Logging {
			fmt.Println("Player dies")
		}
		g.State = BossWon
		return true
	}

	return false
}

func (g *Game) BossTurn() {
	g.startTurn("Boss")
	if g.over() {
		return
	}

	damage := max(1, g.Boss.Damage-g.Player.Armor)
	g.Player.Health -= damage

	if Logging {
		fmt.Printf("Boss attacks for %d damage\n", damage)
	}

	if g.over() {
		return
	}
}

func (g *Game) startTurn(name string) {
	g.Turn++
	if Logging {
		fmt.Printf("-- %s - turn <%d> --\n", name, g.Turn)
		fmt.Println("- ", g.Player)
		fmt.Println("- ", g.Boss)
	}

	if g.HardMode && g.Turn%2 == 1 { // odd turns is play turn
		if Logging {
			fmt.Println("Hard mode: player loses 1 health")
		}
		g.Player.Health--
		if g.Player.Health <= 0 {
			g.State = BossWon
			return
		}
	}

	g.handleActiveEffects()
	if g.Boss.Health <= 0 {
		g.State = PlayerWon
		return
	}
}

func (g *Game) copy() *Game {
	newGame := *g
	newGame.Invocations = make([]Invocation, len(g.Invocations))
	copy(newGame.Invocations, g.Invocations)
	return &newGame
}

func (g *Game) handleActiveEffects() {
	for i := range g.Invocations {
		if g.Invocations[i].Turns == 0 {
			continue
		}
		g.Invocations[i].Turns--
		g.Invocations[i].Spell.Tick(&g.Player, &g.Boss)
		if g.over() {
			return
		}

		if g.Invocations[i].Turns == 0 {
			g.Invocations[i].Spell.Stop(&g.Player, &g.Boss)
			if g.over() {
				return
			}
		}
	}
}

// todo: prefilter spells so this step isn't required..
func (g *Game) isSpellAllowed(spell *Spell) bool {
	if g.Player.Mana < spell.Cost {
		return false
	}

	for _, s := range g.Invocations {
		if s.Spell == spell && s.Turns > 0 {
			return false
		}
	}

	return true
}

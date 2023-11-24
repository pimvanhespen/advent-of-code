package main

import (
	"fmt"
	"github.com/pimvanhespen/advent-of-code/pkg/aoc"
	"io"
)

type Input struct {
	bots    map[ID]*Bot
	outputs map[ID]*Output
	values  []Value
}

type Value struct {
	Value int
	Bot   ID
}

type Receiver interface {
	Receive(int)
}

type ID int

type Bot struct {
	ID   ID
	data []int
	Low  Receiver
	High Receiver

	compareFn func(*Bot, int, int) (int, int)
}

func (b *Bot) Receive(value int) {
	//log.Printf("bot %d receives %d", b.ID, value)
	if len(b.data) == 0 {
		b.data = append(b.data, value)
		return
	}

	low, high := b.compareFn(b, b.data[0], value)
	b.data = b.data[:0]
	b.Low.Receive(low)
	b.High.Receive(high)
}

type Output struct {
	ID   ID
	data []int
}

func NewOutput(id ID) *Output {
	return &Output{
		ID:   id,
		data: make([]int, 0, 1),
	}
}

func (o *Output) Receive(value int) {
	//log.Printf("output %d receives %d", o.ID, value)
	o.data = append(o.data, value)
}

func main() {
	event := aoc.New(2016, 10, parse)
	fmt.Println("1:", aoc.Must(event.Run(part1)))
	fmt.Println("2:", aoc.Must(event.Run(part2)))
}

type FauxReceiver struct {
	Type string
	ID   ID
}

func (f *FauxReceiver) Receive(value int) {
	panic("must not be called")
}

func parse(r io.Reader) (Input, error) {
	lines, err := aoc.ReadLines(r)
	if err != nil {
		return Input{}, err
	}

	var vals []Value

	bots := make(map[ID]*Bot)
	outs := make(map[ID]*Output)

	for _, line := range lines {
		switch line[0] {
		case 'v': // value 5 goes to bot 2
			var num, botID int
			_, err = fmt.Sscanf(line, "value %d goes to bot %d", &num, &botID)
			if err != nil {
				return Input{}, err
			}
			vals = append(vals, Value{num, ID(botID)})

		case 'b': // bot 2 gives low to bot 1 and high to bot 0
			var botID, lowID, highID int
			var lowType, highType string

			_, err = fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &botID, &lowType, &lowID, &highType, &highID)
			if err != nil {
				return Input{}, err
			}

			if lowType == "output" {
				outs[ID(lowID)] = NewOutput(ID(lowID))
			}

			if highType == "output" {
				outs[ID(highID)] = NewOutput(ID(highID))
			}

			bot := &Bot{
				ID:   ID(botID),
				Low:  &FauxReceiver{Type: lowType, ID: ID(lowID)},
				High: &FauxReceiver{Type: highType, ID: ID(highID)},
				data: make([]int, 0, 1),
			}

			bots[bot.ID] = bot

		default:
			return Input{}, fmt.Errorf("unknown line: %s", line)
		}
	}

	for _, bot := range bots {
		low := bot.Low.(*FauxReceiver)
		high := bot.High.(*FauxReceiver)

		switch low.Type {
		case "bot":
			bot.Low = bots[low.ID]
		case "output":
			bot.Low = outs[low.ID]
		}

		switch high.Type {
		case "bot":
			bot.High = bots[high.ID]
		case "output":
			bot.High = outs[high.ID]
		}
	}

	return Input{
		bots:    bots,
		outputs: outs,
		values:  vals,
	}, nil
}

func part1(input Input) string {
	var bot *Bot

	callback := func(caller *Bot, low, high int) (int, int) {
		if low > high {
			low, high = high, low
		}
		if low == 17 && high == 61 {
			bot = caller
		}
		return low, high
	}

	for _, b := range input.bots {
		b.compareFn = callback
	}

	for _, v := range input.values {
		input.bots[v.Bot].Receive(v.Value)
	}

	if bot == nil {
		return "not found"
	}

	return fmt.Sprintf("%d", bot.ID)
}

func part2(input Input) string {
	for _, b := range input.bots {
		b.compareFn = func(_ *Bot, i int, i2 int) (int, int) {
			if i > i2 {
				return i2, i
			}
			return i, i2
		}
	}

	for _, v := range input.values {
		input.bots[v.Bot].Receive(v.Value)
	}

	zero, ok := input.outputs[ID(0)]
	if !ok {
		return "zero not found"
	}

	one, ok := input.outputs[ID(1)]
	if !ok {
		return "one not found"
	}

	two, ok := input.outputs[ID(2)]
	if !ok {
		return "two not found"
	}

	sum := zero.data[0] * one.data[0] * two.data[0]
	return aoc.Result(sum)
}

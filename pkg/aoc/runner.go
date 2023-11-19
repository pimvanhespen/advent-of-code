package aoc

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

var _ Parser[int] = ParserFunc[int](nil)

type ParserFunc[T any] func(io.Reader) (T, error)

func NewParser[T any](fn ParserFunc[T]) Parser[T] {
	return fn
}

func (f ParserFunc[T]) Parse(r io.Reader) (T, error) {
	return f(r)
}

type Parser[Input any] interface {
	Parse(io.Reader) (Input, error)
}

type SolverFunc[Input any] func(Input) string

type Runner[Input any] struct {
	challenge Challenge
	parser    Parser[Input]
	input     Input
}

func (d *Runner[Input]) load() (Input, error) {
	reader, err := d.challenge.Input()
	if err != nil {
		var zero Input
		return zero, fmt.Errorf("get: %w", err)
	}

	input, err := d.parser.Parse(reader)
	if err != nil {
		var zero Input
		return zero, fmt.Errorf("parse: %w", err)
	}

	return input, nil
}

func (d *Runner[Input]) Run(solve SolverFunc[Input]) (string, error) {

	in, err := d.load()
	if err != nil {
		return "", err
	}

	result := solve(in)

	return result, nil
}

func (d *Runner[Input]) String() string {
	return fmt.Sprintf("Runner %s", d.challenge.String())
}

func New[T any](year, day int, parser ParserFunc[T], opts ...Option) *Runner[T] {
	o := defaults()
	for _, opt := range opts {
		opt(o)
	}

	th := slog.NewTextHandler(o.out, &slog.HandlerOptions{
		AddSource: false,
		Level:     o.logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "time":
				return slog.Attr{
					Key:   "time",
					Value: slog.StringValue(a.Value.Time().Format("15:04:05.000")),
				}

			}
			return a
		},
	})
	slog.SetDefault(slog.New(th))

	runner := &Runner[T]{
		challenge: Challenge{
			Year: year,
			Day:  day,
		},
		parser: parser,
	}

	return runner
}

type Option func(*options)

func WithOutput(out io.Writer) Option {
	return func(o *options) {
		o.out = out
	}
}

func WithLogLevel(level slog.Level) Option {
	return func(o *options) {
		o.logLevel = level
	}
}

type options struct {
	out      io.Writer
	logLevel slog.Level
}

func defaults() *options {
	return &options{
		out:      os.Stdout,
		logLevel: slog.LevelInfo,
	}
}

func Must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}

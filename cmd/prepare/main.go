package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	_ "embed"
)

//go:embed templates
var templates embed.FS

type Config struct {
	Year   uint
	Day    uint
	DryRun bool
	Force  bool
}

func (c Config) IsValid() bool {
	return c.Year >= 2015 && c.Day >= 1 && c.Day <= 25
}

func main() {
	var c Config

	year := time.Now().Year()

	flag.UintVar(&c.Year, "year", 0, fmt.Sprintf("year of the event (2015-%d)", year))
	flag.UintVar(&c.Day, "day", 0, "day of the event (1-25)")
	flag.BoolVar(&c.DryRun, "dry-run", false, "do not write to disk")
	flag.BoolVar(&c.Force, "force", false, "overwrite existing file")
	flag.Parse()

	if !c.IsValid() {
		flag.Usage()
		os.Exit(1)
	}

	p := NewPreparer(c)
	if err := p.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

type Preparer struct {
	cfg Config
}

func NewPreparer(cfg Config) *Preparer {
	return &Preparer{cfg: cfg}
}

type Target struct {
	Filename string
}

func (p *Preparer) Run() error {

	targets := []Target{
		{Filename: "main.go.tmpl"},
		{Filename: "main_test.go.tmpl"},
	}

	for _, t := range targets {
		if err := p.prepare(t); err != nil {
			return fmt.Errorf("creating %q: %w", t.Filename, err)
		}
	}

	return nil
}

func (p *Preparer) prepare(t Target) error {

	tpl, err := template.New(t.Filename).ParseFS(templates, filepath.Join("templates", t.Filename))
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	writer, err := p.getWriteCloser(strings.TrimSuffix(t.Filename, ".tmpl"))
	if err != nil {
		return fmt.Errorf("opening file: %w", err)
	}
	defer writer.Close()

	err = tpl.Execute(writer, p.cfg)
	if err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}

func (p *Preparer) getWriteCloser(filename string) (io.WriteCloser, error) {

	if p.cfg.DryRun {
		return os.Stdout, nil
	}

	root := "."

	// check if file exists
	fp := filepath.Join(root, "events", strconv.Itoa(int(p.cfg.Year)), fmt.Sprintf("%02d", p.cfg.Day), filename)

	_, err := os.Stat(fp)
	switch {
	case err == nil:
		if !p.cfg.Force {
			return nil, fmt.Errorf("file already exists: %s", fp)
		}
	case os.IsNotExist(err):
		// file doesn't  exists, ensure directory exists
		err = os.MkdirAll(filepath.Dir(fp), 0755)
		if err != nil {
			return nil, fmt.Errorf("creating directory: %w", err)
		}
	default:
		// any other error is unexpected
		return nil, fmt.Errorf("checking file: %w", err)
	}

	// create file
	f, err := os.Create(fp)
	if err != nil {
		return nil, fmt.Errorf("creating file: %w", err)
	}

	log.Println("writing ", fp)

	return f, nil
}

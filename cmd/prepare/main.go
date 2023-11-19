package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "embed"
)

//go:embed main.go.tmpl
var templateData string

func main() {
	var c Config

	year := time.Now().Year()

	flag.UintVar(&c.Year, "year", 0, fmt.Sprintf("year of the event (2015-%d)", year))
	flag.UintVar(&c.Day, "day", 0, "day of the event (1-25)")
	flag.BoolVar(&c.DryRun, "dry-run", false, "do not write to disk")
	flag.Parse()

	if !c.IsValid() {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(c); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

type Config struct {
	Year   uint
	Day    uint
	DryRun bool
}

func (c Config) IsValid() bool {
	return c.Year >= 2015 && c.Day >= 1 && c.Day <= 25
}

func openFile(year, day uint) (_ io.WriteCloser, err error) {

	y := strconv.Itoa(int(year))
	d := fmt.Sprintf("%02d", day)

	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("getting current working directory: %w", err)
	}

	fp := filepath.Join(wd, "events", y, d, "main.go")
	dir := filepath.Dir(fp)

	defer func() {
		if err != nil {
			return
		}
		_, _ = fmt.Fprintf(os.Stderr, "output: %s\n", fp)
	}()

	_, err = os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("checking directory: %w", err)
		}

		// Directory does not exist, create it
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, fmt.Errorf("creating directory: %w", err)
		}
	}

	_, err = os.Stat(fp)
	switch {
	case err == nil:
		return nil, fmt.Errorf("file already exists: %s", fp)
	case !os.IsNotExist(err):
		return nil, fmt.Errorf("checking file: %w", err)
	default:
		// create file ...
	}

	wc, err := os.Create(fp)
	if err != nil {
		return nil, fmt.Errorf("creating file: %w", err)
	}
	return wc, nil
}

func run(c Config) error {

	var writer io.WriteCloser

	if c.DryRun {
		writer = os.Stdout
	} else {
		var err error
		writer, err = openFile(c.Year, c.Day)
		if err != nil {
			return err
		}
	}

	defer writer.Close()

	tpl := template.Must(template.New("main").Parse(templateData))

	err := tpl.Execute(writer, c)
	if err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
}

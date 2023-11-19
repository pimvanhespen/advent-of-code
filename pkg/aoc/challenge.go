package aoc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const (
	baseURL = "https://adventofcode.com/"
)

func Result(v any) string {
	return fmt.Sprintf("%v", v)
}

type Challenge struct {
	Year int
	Day  int
}

func NewChallenge(year, day int) Challenge {
	return Challenge{
		Year: year,
		Day:  day,
	}
}

func (p Challenge) Input() (io.Reader, error) {
	dir, err := problemDir(p.Year, p.Day)
	if err != nil {
		return nil, err
	}

	input := filepath.Join(dir, "input.txt")

	if existsFile(input) {
		// already downloaded
		reader := getReader(input)
		return reader, nil
	}

	// download and save
	err = download(p.Year, p.Day)
	if err != nil {
		return nil, err
	}

	// return reader
	reader := getReader(input)
	return reader, nil
}

func (p Challenge) String() string {
	return fmt.Sprintf("%04d-%02d", p.Year, p.Day)
}

func problemDir(year, day int) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	p := filepath.Join(wd, "event", fmt.Sprintf("%04d/%02d", year, day))

	return p, nil
}

func problemInput(year, day int) (string, error) {
	dir, err := problemDir(year, day)
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "input.txt"), nil
}

func getReader(path string) io.Reader {
	reader, writer := io.Pipe()
	go func() {
		file, openErr := os.Open(path)
		if openErr != nil {
			_ = writer.CloseWithError(openErr)
			return
		}
		defer file.Close()

		_, cErr := io.Copy(writer, file)
		_ = writer.CloseWithError(cErr)
	}()

	return reader
}

func existsFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var ErrCookie = errors.New("missing cookie.txt in root directory")

func getCookie() (_ string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fp := filepath.Join(wd, "cookie.txt")

	if !existsFile(fp) {
		return "", ErrCookie
	}

	f, err := os.Open(fp)
	if err != nil {
		return "", err
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func download(year, day int) error {
	cookie, err := getCookie()
	if err != nil {
		return err
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	c, err := NewClient(http.DefaultClient, base, cookie)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	reader, err := c.DownloadInput(ctx, year, day)
	if err != nil {
		return fmt.Errorf("download: %w", err)
	}

	return storeInput(year, day, reader)
}

func storeInput(year, day int, reader io.Reader) (err error) {
	input, err := problemInput(year, day)
	if err != nil {
		return err
	}
	f, err := os.Create(input)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(f, reader)
	if err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	return nil
}
